package main

import "fmt"

type customType int

var x customType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
