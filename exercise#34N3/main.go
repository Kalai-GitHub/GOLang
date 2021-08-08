package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println(x)
	i := foo()
	fmt.Println(i())
}

func foo() func() int {
	var x = 10
	return func() int {
		x++
		return x
	}
}
