package main

import "fmt"

func main() {
	s := "A"

	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs) //prints int value of A

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%#X\n", n)

}
