package main

import "fmt"

func main() {
	fmt.Println("Enter a number from 1 to 5")
	var val int
	fmt.Scanf("%d", &val)
	switch val {
	case 1:
		fmt.Println("you have entered - one")
	case 2:
		fmt.Println("you have entered - two")
	case 3:
		fmt.Println("you have entered - three")
	case 4:
		fmt.Println("you have entered - four")
	case 5:
		fmt.Println("you have entered - five")
	}
}
