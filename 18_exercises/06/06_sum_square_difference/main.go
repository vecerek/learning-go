/*
	PROJECT EULER #6 Sum square difference
	https://projecteuler.net/problem=6

	The sum of the squares of the first ten natural numbers is,
				12 + 22 + ... + 102 = 385
	The square of the sum of the first ten natural numbers is,
				(1 + 2 + ... + 10)2 = 552 = 3025
	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import "fmt"

func main() {
	fmt.Println(SumSquareDifferenceOfFirst(100))
}

func SumSquareDifferenceOfFirst(number int) int {
	var sumOfSquares, sum = 1, 1
	var i int
	for i = 2; i <= number; i++ {
		sumOfSquares += i * i
		sum += i
	}

	return sum*sum - sumOfSquares
}

// Reference implementation

func SumSquareDifferenceOfFirstReference(number int) int {
	return sqSum(number) - sumSq(number)
}

func sumSq(num int) int {
	if num < 1 {
		return 0
	}
	return num*num + sumSq(num-1)
}

func sqSum(num int) int {
	var sum int
	for n := 1; n <= num; n++ {
		sum += n
	}
	return sum * sum
}
