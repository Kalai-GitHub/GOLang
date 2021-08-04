package main

import "fmt"

func main() {
	bd := 1985
	count := 0
	for i := bd; i <= 2021; i++ {
		fmt.Println(i)
		count++
	}
	fmt.Println(count)
}
