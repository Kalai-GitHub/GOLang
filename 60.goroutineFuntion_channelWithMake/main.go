package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func main() {
	ch := make(chan int) //making a channel
	go func() {
		ch <- doSomething(5) // calling doSomething() and returning/put the value to the channel ch
	}()
	fmt.Println(<-ch) // printout the value of the channel

	foo()
}

func foo() {
	fmt.Println("Kalai")
}
