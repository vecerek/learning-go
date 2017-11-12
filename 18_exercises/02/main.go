package main

import "fmt"

func main() {
	var res int
	var even bool
	half := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}

	res, even = half(1)
	fmt.Println(res, even)
	res, even = half(2)
	fmt.Println(res, even)
}
