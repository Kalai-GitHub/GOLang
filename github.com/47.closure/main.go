package main

import "fmt"

func main() {
	fmt.Println(increment()()) //x value in increment() will be 0 in the starting, returns 1
	a := increment()
	fmt.Println(a()) //x value in increment() will be 0 in the starting, returns 1
	fmt.Println(a()) //x value in increment() will be 2
	fmt.Println(a()) //x value in increment() will be 3

	b := increment() //x value in increment() will be 0 in the starting, returns 1
	fmt.Println(b())

	c := scope()
	fmt.Println("Scope Method: ", c())

	d, f := outer()
	fmt.Println(d(), f)

}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var }
	return foo
}

//below code wont compile - outer_var is not defined in this scope
// func another_scope() func() int {
// 	outer_var := 444
// 	return foo
// }

func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99
		return outer_var
	}
	return inner, outer_var
}
