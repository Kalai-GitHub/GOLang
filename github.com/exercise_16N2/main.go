package main

import "fmt"

func main() {
	//getting 4th element of an array
	ar := [6]int{10, 20, 30, 40, 50, 60}
	fmt.Println(ar[4])

	//getting 4th element of a slice
	sl := []int{100, 200, 300, 400, 500, 600}
	fmt.Println(sl[4])
}
