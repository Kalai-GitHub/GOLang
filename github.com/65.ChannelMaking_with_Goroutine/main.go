package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Println(c)

	go func() {
		c <- 42
	}()
	fmt.Println(c)
}
