package main

import "testing"

func BenchmarkSumSquareDifferenceOfFirst(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SumSquareDifferenceOfFirst(100)
	}
}

func BenchmarkSumSquareDifferenceOfFirstReference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SumSquareDifferenceOfFirstReference(100)
	}
}
