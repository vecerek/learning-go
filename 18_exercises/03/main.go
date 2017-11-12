package main

import "fmt"

func main() {
	x := findGreatest([]int{1, 2, 3, 4, 5, 4, 3, 0}...)
	fmt.Println(x)
}

func findGreatest(nums ...int) int {
	greatest := nums[0]
	for _, num := range nums {
		if num > greatest {
			greatest = num
		}
	}

	return greatest
}
