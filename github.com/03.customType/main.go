package main

import "fmt"

var a int = 42

type hotdog int //custom type(hotdog)- underlying type of hotdog is int

var b hotdog = 45

func main() {
	fmt.Println("a, b = ", a, b)
	fmt.Printf("a = %T\nb = %T\n", a, b)
	a = int(b) //conversion(this is not 'type casting' like other languages)
	fmt.Println("a = ", a)
	fmt.Printf("a = %T\nb = %T\n", a, b)
}
