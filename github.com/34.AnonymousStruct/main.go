package main

import "fmt"

func main() {
	// anonymous struct - without having any name
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "kalai",
		last:  "murugu",
		age:   30,
	}

	fmt.Println(p.first, p.last, p.age)

	q := struct {
		first string
	}{
		first: "UUUU",
	}
	fmt.Println(q.first)
}
