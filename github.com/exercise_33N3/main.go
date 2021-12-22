package main

import "fmt"

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		} else if len(xi) == 1 {
			return 1
		} else {
			return xi[0] + xi[len(xi)-1]
		}
	}
	a := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(a)
}

func foo(f func(v []int) int, x []int) int {
	n := f(x)
	n++
	return n
}
