package main

import "fmt"

func main() {
	var value int
	fmt.Println("Enter a number")
	fmt.Scanln(&value)

	if value%2 == 0 {
		fmt.Println(value, " is even")
	} else {
		fmt.Println(value, " is odd")
	}
}
