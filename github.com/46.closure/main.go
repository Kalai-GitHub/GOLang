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
		var y int = 30
		y = 30
		fmt.Println(y)
		fmt.Println(x)
	}
}
func foo() {
	x++
	fmt.Println(x)
}
