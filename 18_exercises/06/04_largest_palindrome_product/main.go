/*
	PROJECT EULER #4 Largest prime factor
	https://projecteuler.net/problem=4

	A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
	Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(LargestPalindromeProductOf3DigitFactors())
}

func LargestPalindromeProductOf3DigitFactors() (int, int, int) {
	for i := 999; i >= 100; i-- {
		palindrome, _ := buildPalindrome(i)
		a, b := ThreeDigitFactorsOf(palindrome)
		if a != palindrome {
			return palindrome, a, b
		}
	}

	return 0, 0, 0
}

func ThreeDigitFactorsOf(num int) (int, int) {
	for i := 101; i < 1000; i++ {
		if num%i == 0 {
			b := num / i
			if b < 1000 {
				return i, b
			}
		}
	}

	return num, 1
}

func buildPalindrome(i int) (int, error) {
	firstHalf := strconv.Itoa(i)
	sPalindrome := firstHalf + reverseString(firstHalf)
	return strconv.Atoi(sPalindrome)
}

func reverseString(s string) string {
	r := []rune(s)
	tmp := r[0]
	if tmp != r[2] {
		r[0] = r[2]
		r[2] = tmp
	}
	return string(r)
}

// Reference solution found below

func reverse(init string) string {
	value := strings.Split(init, "")
	for i := 0; i < len(value) / 2; i++ {
		value[i], value[len(value) - i - 1] = value[len(value) - i - 1], value[i]
	}
	return strings.Join(value, "")
}

func LargestPalindromeProductOf3DigitFactorsReference() int {
	largest := 0
	for a := 999; a > 99; a-- {
		for b := 999; b > 99; b-- {
			v := a * b
			if largest > v {
				break
			}
			forward := strconv.Itoa(v)
			backward := reverse(forward)
			if forward == backward {
				largest = v
			}
		}
	}

	return largest
}
