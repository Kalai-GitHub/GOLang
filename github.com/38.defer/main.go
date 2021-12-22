package main

import "fmt"

func main() {
	defer foo() // always executes at the end of the function
	mouse()
}

func foo() {
	fmt.Println("foo")
}

func mouse() {
	fmt.Println("mouse")
}
