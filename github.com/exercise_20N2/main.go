package main

import "fmt"

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	sl := [][]string{x, y}
	fmt.Println(sl)

	for i, outArray := range sl {
		fmt.Println("index:", i)
		for j, innerArray := range outArray {
			fmt.Printf("index position: %v || value: %v\n", j, innerArray)
		}
	}

}
