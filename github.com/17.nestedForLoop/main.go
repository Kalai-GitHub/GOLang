package main

import "fmt"

func main() {
	for i := 1; i <= 6; i++ {
		fmt.Printf("Outer for loop: %d\n", i)
		for j := 1; j < 3; j++ {
			fmt.Printf("\tInner for loop: %d\n", j)
		}
	}
	x := 1
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
