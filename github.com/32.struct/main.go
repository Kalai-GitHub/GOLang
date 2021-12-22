package main

import "fmt"

func main() {
	type person struct {
		firstName string
		lastName  string
		age       int
	}

	p1 := person{
		firstName: "Kalai",
		lastName:  "Murugu",
		age:       30,
	}
	p2 := person{
		firstName: "Kavin",
		lastName:  "Thiyagu",
		age:       10,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.firstName, p1.lastName)
	fmt.Println(p2.firstName, p2.lastName)
}
