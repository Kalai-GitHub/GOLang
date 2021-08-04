package main

import "fmt"

func main() {
	for i := 30; i <= 50; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
