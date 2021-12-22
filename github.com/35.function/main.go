package main

import "fmt"

func main() {
	//func (r receiver) identifier(parameters) (return s){...}
	//everything in Go is 'PASS BY VALUE'

	foo()
	hello("Kalai") // pass by value
	s1 := woo("Kavin")
	fmt.Println(s1)

	x, y := mouse("kavin", "thiyagu")
	fmt.Println(x)
	fmt.Println(y)
}

//no parameter
func foo() {
	fmt.Println("Hello from foo")
}

//with a parameter
func hello(s string) {
	fmt.Println("Hello, ", s)
}

//with a parameter and a return type
func woo(s string) string {
	return fmt.Sprint("Hello, ", s)
}

//with two parameters and two return types
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}
