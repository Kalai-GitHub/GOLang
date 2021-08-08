package main

import "fmt"

//creating g with type func() and assgning value to the g
var g func() = func() { //var g = func(){
	fmt.Println("g from outside main")
}

func main() {
	a := func() {
		for i := 10; i < 15; i++ {
			fmt.Println(i)
		}
	}
	a()
	fmt.Printf("a is: %T\n", a)

	g()
	g = a
	g()
	fmt.Printf("g is :%T\n", g)
}
