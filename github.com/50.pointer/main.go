package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before :", x)  //0
	fmt.Println("x before :", &x) //address of x
	foo(&x)
	fmt.Println("x after :", x)  //42
	fmt.Println("x after :", &x) //address of x

}

func foo(y *int) {
	fmt.Println("y before :", y)  //address
	fmt.Println("y before :", *y) //0
	*y = 42
	fmt.Println("y after :", y)  //address
	fmt.Println("y after :", *y) //42
}
