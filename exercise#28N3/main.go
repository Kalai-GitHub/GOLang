package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Kalai",
		last:  "murugu",
		age:   37,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}
