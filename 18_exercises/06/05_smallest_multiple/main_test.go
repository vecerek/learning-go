package main

import "testing"

func BenchmarkSmallestMultipleOfFirst(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SmallestMultipleOfFirst(20)
	}
}

func BenchmarkSmallestMultipleOfFirstReference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SmallestMultipleOfFirstReference(20)
	}
}

func BenchmarkSmallestMultipleOfFirstOptimizedReference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SmallestMultipleOfFirstOptimizedReference(20)
	}
}
