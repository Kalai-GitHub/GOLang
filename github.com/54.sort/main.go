package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{3, 1, 9, 10, 0, 5}
	str := []string{"kalai", "zebra", "apple", "yolk", "ball"}

	fmt.Println("Before sorted:", x)
	sort.Ints(x)
	fmt.Println("After sorted:", x)

	fmt.Println("Before sorted:", str)
	sort.Strings(str)
	fmt.Println("After sorted:", str)

}
