package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) //& gives address. int and *int both are different types
	fmt.Printf("%T\n", *&a)

	var b *int = &a //or b:= &a
	fmt.Println(*b) //dereferncing the address and getting the value. * gives the value stored at an address
	fmt.Println(b)

	*b = 43
	fmt.Println("value a:", a)
}
