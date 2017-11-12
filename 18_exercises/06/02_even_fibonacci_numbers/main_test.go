package main

import "testing"

func BenchmarkSumOfEvenFibonacciVer1(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SumOfEvenFibonacciVer1()
	}
}

func BenchmarkSumOfEvenFibonacciVer2(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SumOfEvenFibonacciVer2()
	}
}

func BenchmarkSumOfEvenFibonacciVer3(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SumOfEvenFibonacciVer3()
	}
}

func BenchmarkReference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Reference()
	}
}
