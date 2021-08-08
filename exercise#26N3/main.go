package main

import "fmt"

func main() {
	x := []int{5, 10, 15, 20, 25, 30}
	f := foo(x...) //unfurling a slice
	b := bar(x)
	fmt.Println(f, b)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
