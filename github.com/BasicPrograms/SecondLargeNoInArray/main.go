package main

import "fmt"

func main() {
	arr := []int{10, 5, 7, 2, 11}
	large1 := arr[0]
	large2 := 0
	for i := 1; i < len(arr); i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
			fmt.Println(large1, large2)
		} else if large2 < arr[i] {
			large2 = arr[i]
			fmt.Println(large2)
		}
	}
	fmt.Println(large2)

}
