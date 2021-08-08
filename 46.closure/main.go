package main

import "fmt"

//closure is a scope of the variable so the scope of the variable will be limited to that area
var x int

func main() {
	fmt.Println(x)
	foo()
	x++
	fmt.Println(x)
	{
		var y int
		y = 30
		fmt.Println(y)
	}
}
func foo() {
	x++
	fmt.Println(x)
}
