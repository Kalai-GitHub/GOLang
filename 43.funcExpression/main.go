package main

import "fmt"

func main() {
	//function expression - is assigning an anonymous function to a variable
	//functions are first class citizen, meaning function can do what all the other types can do

	//func expression without argument
	f := func() {
		fmt.Println("anonymous function assigned to a variable is called function expression")
	}
	f()

	//func expression with argument
	g := func(x int) {
		fmt.Println("value of x is ", x)
	}
	g(10)

}
