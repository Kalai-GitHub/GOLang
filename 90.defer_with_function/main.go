package main

import "fmt"

func main() {
	y := foo()
	fmt.Println(y)
}

func foo() (i int) {
	defer func() { i++ }() // at the end, i will be incremented to 2
	return 1               // first, this 1 will return to i.
}
