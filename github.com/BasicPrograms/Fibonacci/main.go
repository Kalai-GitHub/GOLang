package main

import "fmt"

func main() {
	var n int
	t1 := 0
	t2 := 1
	nextTerm := 0

	fmt.Println("Enter a number to find the fibonacci series")
	fmt.Scanln(&n)

	for i := 1; i < n; i++ {
		if i == 1 {
			fmt.Print(t1, " ")
		} else if i == 2 {
			fmt.Print(t2, " ")
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		fmt.Print(nextTerm, " ")
	}
}
