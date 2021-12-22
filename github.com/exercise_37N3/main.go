package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	fmt.Println("in change me method:")
	//to derefernce the value - can use either (*p).first or p.first="kavin"
	(*p).first = "Kavin"
	p.last = "thiyagu"
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "kalai",
		last:  "murugu",
		age:   37,
	}
	fmt.Println("Before dereferncing in main", p)
	changeMe(&p)
	fmt.Println("after dereferencing in main", p)

}
