package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 32
		close(c)
	}()

	i, ok := <-c
	fmt.Println(i, ok)

	i, ok = <-c
	fmt.Println(i, ok) //i will become 0 and ok will become false

}
