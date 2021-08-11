package main

import "fmt"

func main() {
	c := make(chan<- int, 1) //only to send the value onto the channel

	c = 33

	fmt.Println(<-c) //invalid operation: cannot receive from send-only channel c (variable of type chan<- int)compilerInvalidReceive

}
