package main

import "fmt"

type customType int

var x customType
var y customType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y := int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
