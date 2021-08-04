package main

import "fmt"

func main() {
	s := "play, ground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // slice of byte
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) // byte is alias of unit8

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U \n", s[i]) // gives UTFA code that is 'rune' - this is alias of int32
	}

	//to get index and value of s(string)
	for i, v := range s {
		fmt.Println(i, v)
	}

	//to get s value in hexadecimal
	for i, v := range s {
		fmt.Printf("Hexadecimal value of the index %d is %#x \n", i, v)
	}

	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "first " + x
	fmt.Println(x)

}
