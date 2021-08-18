package main

import "fmt"

func main() {
	var arr [35]int
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i, v := range arr {
		fmt.Printf("Value: %d\t Type: %T\tBinary: %b\n", v, v, v)
		if i > 25 {
			break
		}
	}
}
