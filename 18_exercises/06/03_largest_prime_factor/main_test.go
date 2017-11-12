package main

import "testing"

func BenchmarkLargestPrimeFactorOf(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LargestPrimeFactorOf(600851475143)
	}
}

func BenchmarkLargestPrimeFactorOfReference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		LargestPrimeFactorOfReference(600851475143)
	}
}
