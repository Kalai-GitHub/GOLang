package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //pulled of the channel until channel closed
	}()

	//receive
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Exit")
}
