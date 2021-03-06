package main

import "fmt"

var i = 43
var j = "Hello type"

func main() {
	fmt.Println("i :", i)

	// format printing
	fmt.Printf("%T\n", i)  //type of i int
	fmt.Printf("%b\n", i)  //base 2(binary) of i
	fmt.Printf("%x\n", i)  //base 16(hexadecimal), with lower letters for a-f
	fmt.Printf("%#x\n", i) //base 16(hexadecimal with 0x in the prefix)
	fmt.Printf("%T\t%b\t%#x\n", i, i, i)

	//to get normal value of i
	fmt.Printf("%v\n", i)

	fmt.Println("j :", j)
	fmt.Printf("%T\n", j)

	// String print
	s := fmt.Sprintf("%T\t%b\t%#x\n", i, i, i)
	fmt.Println(s)

}
