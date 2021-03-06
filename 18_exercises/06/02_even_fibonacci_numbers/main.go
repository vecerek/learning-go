/*
	PROJECT EULER #2 Even Fibonacci numbers
	https://projecteuler.net/problem=2

	Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

									1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package main

import "fmt"

func main() {
	fmt.Println(SumOfEvenFibonacciVer1())
	fmt.Println(SumOfEvenFibonacciVer2())
	fmt.Println(SumOfEvenFibonacciVer3())
	fmt.Println(Reference())
}

func SumOfEvenFibonacciVer1() int {
	var sum int
	previous, current := 1, 1

	for sum <= 4000000 {
		tmp := current
		current += previous
		previous = tmp

		if current%2 == 0 {
			sum += current
		}
	}

	return sum
}

func SumOfEvenFibonacciVer2() int {
	var sum int
	previous, current, i := 1, 1, 3

	for sum <= 4000000 {
		tmp := current
		current += previous
		previous = tmp

		if i == 3 {
			sum += current
			i = 0
		}
		i++
	}

	return sum
}

func SumOfEvenFibonacciVer3() int {
	sum, current, previous := 2, 2, 1

	for {
		for i := 0; i < 3; i++ {
			tmp := current
			current += previous
			previous = tmp
		}
		if sum <= 4000000 {
			sum += current
		} else {
			break
		}
	}

	return sum
}

func Reference() int {
	const limit = 4000000

	var (
		sum      int
		prevPrev int
		prev     = 1
	)

	for next := 0; next < limit; next = prevPrev + prev {
		prevPrev = prev
		prev = next

		if next%2 == 0 {
			sum += next
		}
	}

	return sum
}
