package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   int
	}
	type secretAgent struct {
		person
		spy   bool
		first string
	}

	sa := secretAgent{
		person: person{
			first: "kalai",
			last:  "murugu",
			age:   22,
		},
		spy:   true,
		first: "TTTTTT",
	}

	fmt.Println(sa)
	fmt.Println(sa.person.first)
	fmt.Println(sa.last)
	fmt.Println(sa.age)
	fmt.Println(sa.first)

}
