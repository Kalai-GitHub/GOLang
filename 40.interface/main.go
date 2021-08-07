package main

import "fmt"

type person struct { //[type=keyword, person=identifier, struct=other type]
	first string
	last  string
}

type secretAgent struct {
	person
	bo bool
}

type human interface { //[keyword, identifier, other type]
	speak() // any type that has method speak(), that type is also a type human. So this
}

func (sa secretAgent) speak() { // method receiver will attach the speak() to its type secretAgent
	fmt.Println(sa.first, "\n", sa.last, "\n", sa.bo)
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}

func bar(h human) {
	fmt.Println("hello ", h)
	switch h.(type) {
	case person:
		fmt.Println("I am in bar with ", h.(person).first)
	case secretAgent:
		fmt.Println("Iam in bar with ", h.(secretAgent).first)
	}
}

func main() {
	//a value can be more than one types. sa1 is a value, which is type of secretAgent and human
	sa1 := secretAgent{
		person: person{
			first: "kalai",
			last:  "murugu",
		},
		bo: true,
		/*values of secretAgent-
		[person: person{
			first: "kalai",
			last:  "murugu",
		},
		bo: true,] ===>these values can be of type secretAgent and human
		*/
	}

	p := person{
		first: "Kavin",
	}

	sa1.speak()
	p.speak()

	//polymorphism - bar function takes many different
	bar(sa1)
	bar(p)

}
