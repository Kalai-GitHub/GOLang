package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Here is the error %v", c.info)
}

func foo(e error) {
	//fmt.Println("foo ran ", e)
	fmt.Println("foo ran ", e, "\n", e.(customErr).info) //e.(customErr) - this is called assertion
}

func main() {
	c1 := customErr{
		info: "I need a pizza",
	}
	foo(c1)
}
