package main

import "fmt"

func main() {
	func() {
		fmt.Println("from an anonymous function")
		for i := 10; i < 20; i++ {
			fmt.Println(i)
		}
	}()
}
