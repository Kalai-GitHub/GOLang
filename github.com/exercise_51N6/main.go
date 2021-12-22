package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)
	//specific only channel will not receive any value from the channel
	//cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
