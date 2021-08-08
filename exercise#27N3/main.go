package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	defer func() {
		fmt.Println("bar inside anonymous function with statement 'defer'")
	}()
	fmt.Println("bar")
}
