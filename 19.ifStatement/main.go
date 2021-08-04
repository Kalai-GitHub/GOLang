package main

import "fmt"

func main() {
	if !false {
		fmt.Println("001")
	}
	if 2 == 2 {
		fmt.Println("002")
	}
	if 2 != 2 {
		fmt.Println("003")
	}
	if !(2 == 2) {
		fmt.Println("004")
	}
	if !(2 != 2) {
		fmt.Println("005")
	}

	if x := 42; x == 2 {
		fmt.Println("condition is true")
	}

}
