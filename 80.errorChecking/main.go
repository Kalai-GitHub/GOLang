package main

import "fmt"

func main() {
	i, err := fmt.Println("Hello") //hello=5, println=1[print=0], 5+1=6
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
