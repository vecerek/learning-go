package _1

import "fmt"

func main() {
	var res int
	var even bool

	res, even = half(1)
	fmt.Println(res, even)
	res, even = half(2)
	fmt.Println(res, even)
}

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}
