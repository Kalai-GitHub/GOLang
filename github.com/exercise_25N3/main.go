package main

import "fmt"

func main() {
	a := foo()
	i, s := bar()
	fmt.Println(a, i, s)
}

func foo() int {
	var x int = 10
	return x
}

func bar() (string, int) {
	var s string = "Hello "
	var i int = 45
	return s, i
}
