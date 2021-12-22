package main

import "fmt"

var i = 43
var j = "Hello type"

func main() {
	fmt.Println("i :", i)

	fmt.Printf("%T\n", i)
	fmt.Printf("%b\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%T\t%b\t%#x\n", i, i, i)

	fmt.Println("j :", j)
	fmt.Printf("%T\n", j)
}
