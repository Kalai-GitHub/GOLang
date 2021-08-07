package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("%T\n", x)
	c := x()
	fmt.Println(c)
}

func bar() func() int {
	return func() int {
		return 45
	}
}
