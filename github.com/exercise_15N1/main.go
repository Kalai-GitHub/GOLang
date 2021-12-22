package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("number %v/4: remainder %v\n", i, i%4)
	}

	i := "James Bond"
	if i == "Bond" {
		fmt.Println("Bond")
	} else if i == "James Bond" {
		fmt.Println("James Bond", i)
	} else {
		fmt.Println("JamesBond")
	}

	switch {
	case true:
		fmt.Println("This is true block")
	case false:
		fmt.Println("This is false block")
	}

	favSport := "throwball"
	switch favSport {
	case "swim":
		fmt.Println("go to pool")
	case "throwball":
		fmt.Println("go to throwball court")
	case "Badminton":
		fmt.Println("go to badminton court")

	}
}
