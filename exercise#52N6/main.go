package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() { // without the goroutine unable to complete the below for loop.
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // this is very important to close the channel, otherwise will get the deadlock issue
	}()
	return c

}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
