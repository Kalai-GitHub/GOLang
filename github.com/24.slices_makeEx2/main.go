package main

import "fmt"

func main() {
	var a []int
	fmt.Println(len(a))

	//type1: make function
	a = make([]int, 5) // make function helps to set the length of the slice(5)
	fmt.Println(len(a), a)

	//type2: make function
	a = make([]int, 6, 10) //to set the length(6) and the capacity(10) of the slice
	fmt.Println(a, "len:", len(a), "capacity:", cap(a))

	for i := 0; i <= 5; i++ {
		a[i] = i

	}

	for _, v := range a {
		fmt.Println(v)
	}

}
