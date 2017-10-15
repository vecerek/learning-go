package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%v \t %T \n", d, d)
	fmt.Printf("%v \t %T \n", e, e)
	fmt.Printf("%v \t %T \n", f, f)
	fmt.Printf("%v \t %T \n", g, g)
}
