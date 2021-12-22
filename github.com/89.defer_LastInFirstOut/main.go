package main

import "fmt"

func main() {
	b()
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i) // last in first out, 3,2,1,0
	}
}
