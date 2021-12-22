package main

import "fmt"

var a int = 42

func main() {
	// a int = 33     => not working
	fmt.Println(a)
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
