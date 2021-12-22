package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //pulled off the channel until channel closed. It indicates that no more value will be sent on it
	}()

	//receive
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Exit")
}
