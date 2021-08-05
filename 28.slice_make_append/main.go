package main

import "fmt"

func main() {
	x := make([]int, 8, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 100)
	fmt.Println(x)

	x = append(x, 101, 102) // dynamically increasing the size of the slice
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
