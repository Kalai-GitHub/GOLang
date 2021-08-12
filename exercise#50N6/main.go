package main

import "fmt"

func main() {
	c := make(chan int)
	//using function literal, aka, anonymous self-executing func
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	//with buffered channel
	c1 := make(chan int, 1)
	c1 <- 42
	fmt.Println(<-c1)
}
