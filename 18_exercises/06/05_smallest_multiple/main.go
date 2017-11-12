/*
	PROJECT EULER #5 Smallest multiple
	https://projecteuler.net/problem=5

	2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SmallestMultipleOfFirst(20))
}

func SmallestMultipleOfFirst(x int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	primesLen := len(primes)
	allPrimeFactors := make(map[int]int)
	for _, prime := range primes {
		allPrimeFactors[prime] = 1
	}

	primeFactors := make(map[int]int)
	var N, prime int
	for i := 4; i <= x; i++ {
		if _, ok := allPrimeFactors[i]; ok {
			continue
		}

		N = i
		for j := 0; j < primesLen; j++ {
			prime = primes[j]
			if N%prime == 0 {
				N = N / prime
				j--
				if _, ok := primeFactors[prime]; ok {
					primeFactors[prime] += 1
				} else {
					primeFactors[prime] = 1
				}
				if N == 1 {
					break
				}
			}
		}

		for k, v := range primeFactors {
			if v > allPrimeFactors[k] {
				allPrimeFactors[k] = v
			}
		}
		primeFactors = make(map[int]int)
	}

	lcm := 1
	for _, prime := range primes {
		if prime > x {
			break
		}
		for i := 0; i < allPrimeFactors[prime]; i++ {
			lcm *= prime
		}
	}

	return lcm
}

func SmallestMultipleOfFirstOptimizedReference(number int) int {
	primeMap := GeneratePrimeMapOfPrimesLessOrEqualTo(number)

	var val int
	for key := range primeMap {
		for {
			val = primeMap[key] * key
			if val > number {
				break
			}
			primeMap[key] = val
		}
	}

	lcm := 1
	for _, val := range primeMap {
		lcm *= val
	}
	return lcm
}

/*
	Implements Sieve of Atkins
	source: https://stackoverflow.com/a/21854246

	TODO: Create a concurrent version of this algorithm
*/
func GeneratePrimeMapOfPrimesLessOrEqualTo(N int) map[int]int {
	var x, y, n int
	nsqrt := math.Sqrt(float64(N))

	is_prime := make([]bool, N)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	primeMap := make(map[int]int)
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primeMap[x] = x
		}
	}

	return primeMap
}

// Reference implementation below

func isPrime(number int) bool {

	primeNum := 2
	maxPrime := number

	for primeNum < maxPrime {
		if number%primeNum == 0 {
			return false
		}
		maxPrime = number / primeNum
		primeNum++
	}
	return true
}

func SmallestMultipleOfFirstReference(number int) int {
	primeMap := map[int]int{2: 2, 3: 3}

	num := 4
	for num <= number {
		if isPrime(num) {
			primeMap[num] = num
		}
		num++
	}

	for key := range primeMap {
		for {
			val := primeMap[key] * key
			if val > number {
				break
			}
			primeMap[key] = val
		}
	}

	smallSum := 1
	for _, val := range primeMap {
		smallSum = smallSum * val
	}
	return smallSum
}
