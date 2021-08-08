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

}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
