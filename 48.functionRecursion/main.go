package main

import "fmt"

func main() {
	x := 4

	v := factorial(x)
	fmt.Println(v)

	lf := loopFactorial(x)
	fmt.Println(lf)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	total := 1
	for i := n; i >= 1; i-- {
		total *= i
	}
	return total
}
