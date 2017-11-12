package main

import "testing"

func BenchmarkMultiplesOf3And5(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		MultiplesOf3And5()
	}
}

func BenchmarkMultiplesOf3And5Reference(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		MultiplesOf3And5Reference()
	}
}
