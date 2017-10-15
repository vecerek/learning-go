package main

import (
	"fmt"
	"github.com/vecerek/learning-go/04_packages/stringutils"
)

func main() {
	fmt.Println(stringutils.MyName)
	fmt.Println(stringutils.Reverse("Hello, Go!"))
}
