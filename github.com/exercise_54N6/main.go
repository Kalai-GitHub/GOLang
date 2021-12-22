package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 32

	}()
	i, ok := <-c
	fmt.Println(i, ok)
	close(c)
	i, ok = <-c
	fmt.Println(i, ok)
}
