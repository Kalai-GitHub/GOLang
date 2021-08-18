package main

import "fmt"

func main() {
	var arr [25]byte
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = byte(i)
	}

	for i, v := range arr {
		fmt.Printf("value: %d\tType :%T\tBinary :%b\n", v, v, v)
		if i > 15 {
			break
		}
	}
}
