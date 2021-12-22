package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Kalai",
		last:  "Murugu",
		age:   37,
	}

	p.speak()
	saySomething(&p) //can pass only &p
	//saySomething(p) - this is wrong, can not use p as type human in argument to saySomething()

}
