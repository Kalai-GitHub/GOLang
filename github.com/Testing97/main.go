package main

import "fmt"

func main() {
	fmt.Println("2+3 =", sum(2, 3))
	fmt.Println("4+4+4 =", sum(4, 4, 4))
	fmt.Println("6+1 =", sum(6, 1))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
