package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 37, 83, 27,
		19, 97, 9, 17,
	}
	small := x[0]
	for i := 1; i < len(x)-1; i++ {
		if x[i] < small {
			small = x[i]
		}
	}
	fmt.Println(small)

}
