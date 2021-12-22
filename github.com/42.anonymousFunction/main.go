package main

import "fmt"

func main() {

	foo()

	//anonymous function without argument
	func() {
		fmt.Println("anonymous function")
	}() //this will call this anonymous function

	//anonymous function with an argument
	func(x int) {
		fmt.Println("X value is : ", x)
	}(6) //self invoking/calling

}

func foo() {
	fmt.Println("foo")
}
