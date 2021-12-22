package main

import "fmt"

func main() {
	c := make(chan int, 3)

	c <- 42 //send onto the channel
	c <- 52
	c <- 62

	fmt.Println(<-c) //receive from the channel
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("_________________________________")

	fmt.Printf("%T\n", c)
}
