package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Here is the error %v", c.info)
}

func foo(e error) {
	fmt.Println("foo ran ", e)
}

func main() {
	c1 := customErr{
		info: "I need a pizza",
	}
	foo(c1) //first
}
