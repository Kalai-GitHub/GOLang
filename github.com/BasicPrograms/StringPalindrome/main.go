package main

import "fmt"

func main() {
	var str string
	var reverse string = ""
	fmt.Println("Enter a string ")
	fmt.Scanln(&str)

	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	// for i := range str {
	// 	if reverse[i] != str[i] {
	// 		fmt.Printf("%v is not a palindrome", str)
	// 	} else {
	// 		fmt.Printf("%v is a palindrome", str)
	// 	}
	// }

	if reverse == str {
		fmt.Printf("%v is a palindrome\n", str)
	} else {
		fmt.Printf("%v is not a palindrome\n", str)
	}
}
