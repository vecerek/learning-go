/*
	PROJECT EULER #3 Largest prime factor
	https://projecteuler.net/problem=3

	The prime factors of 13195 are 5, 7, 13 and 29.
	What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(LargestPrimeFactorOf(600851475143))
}

func LargestPrimeFactorOf(N uint64) uint64 {
	var i uint64 = 2
	limit := uint64(math.Sqrt(float64(N)))
	for ; i < limit; i++ {
		if N%i == 0 {
			limit = uint64(math.Sqrt(float64(N)))
			N = N / i
		}
	}

	return N
}

func LargestPrimeFactorOfReference(N uint64) uint64 {
	var i uint64 = 2
	for ; i < N; i++ {
		if N%i == 0 {
			N = N / i
		}
	}

	return N
}
