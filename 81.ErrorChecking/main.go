package main

import "fmt"

func main() {
	var name, favFood, favSport string
	fmt.Print("Enter name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter favFood: ")
	_, err = fmt.Scan(&favFood)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter favSport: ")
	_, err = fmt.Scan(&favSport)
	if err != nil {
		panic(err)
	}
	fmt.Println(name, favFood, favSport)
}
