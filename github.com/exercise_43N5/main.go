package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Anything from cat1")
	}()

	go func() {
		fmt.Println("Anything from cat2")
	}()

	fmt.Println("Exit")
}
