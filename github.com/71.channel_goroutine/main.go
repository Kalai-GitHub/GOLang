package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)
	fmt.Println("Done!")
}

func foo(c chan<- int) {
	c <- 32
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
