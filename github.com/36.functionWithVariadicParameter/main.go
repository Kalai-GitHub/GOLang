package main

import "fmt"

func main() {
	i := foo(2, 3, 4, 5, 6)
	fmt.Println("Sum:", i)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("index is: ", i, " value is: ", v, " sum is: ", sum)
	}

	return sum
}
