package main

import "testing"

func BenchmarkLargestPalindromeProductOf3DigitFactors(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LargestPalindromeProductOf3DigitFactors()
	}
}

func BenchmarkLargestPalindromeProductOf3DigitFactorsReference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LargestPalindromeProductOf3DigitFactorsReference()
	}
}
