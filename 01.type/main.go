package main

import "fmt"

var x int = 5
var y string

// int, float64, string
func main() {
	a, b, c := 100+100, 100.10+100.0, 100.0+100.0 //variables, operands and operator

	fmt.Println("a = ", a)

	fmt.Println("b = ", b)
	fmt.Printf("%T\n", b)

	fmt.Println("c = ", c)
	fmt.Printf("%T\n", c)

	fmt.Println("x = ", x)
	fmt.Printf("%T\n", x)

	y = "45"
	fmt.Println("y = ", y)
	fmt.Printf("%T\n", y)

	if (true && false) || (false && true) || !(false && false) {
		fmt.Println("Hello, " + "Wolrd")
		fmt.Println("length : ", len("Hello Wolrd"))
		fmt.Println("Hello Wolrd"[1])
	}

}
