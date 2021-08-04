package main

import "fmt"

func main() {
	fmt.Println("Enter a number from 1 to 7")
	var val int
	fmt.Scanf("%d", &val)
	switch val {
	case 1:
		fmt.Println("you have entered - one")
		fallthrough // it will execute the next case too
	case 2:
		fmt.Println("you have entered - two")
	case 3:
		fmt.Println("you have entered - three")
	case 4:
		fmt.Println("you have entered - four")
		fallthrough // it will execute the next case too
	case 5, 6, 7:
		fmt.Println("you have entered - five/six/seven")
	default:
		fmt.Println("default case")
	}
}
