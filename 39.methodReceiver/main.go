package main

import "fmt"

type person struct {
	fn string
	ln string
}
type secretAgent struct {
	person
	fl bool
}

func (p person) foo() { // receiver will attach the foo function to the type person
	fmt.Println(p.fn)

}
func main() {
	sa1 := secretAgent{
		person: person{
			fn: "Kalai",
			ln: "murugu",
		},
		fl: false,
	}
	sa2 := secretAgent{
		person: person{
			fn: "Kuhan",
			ln: "Thiyagu",
		},
		fl: true,
	}

	sa1.foo()
	sa2.foo()
}
