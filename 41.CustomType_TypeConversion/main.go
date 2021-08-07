package main

import "fmt"

type hotdog int

type contact struct {
	greeting string
	name     string
}

func main() {
	var x hotdog = 40
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x) // type conversion
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//fmt.Println(x + hotdog(y)) //type conversion
	fmt.Println("*****************")
	c := contact{
		greeting: "Hello",
		name:     "House",
	}
	var c1 = contact{"Pillaiyar", "House"}
	switchOnType(7)
	switchOnType("Kalai")
	switchOnType(true)
	switchOnType(c)
	switchOnType(c1)
	switchOnType(c1.greeting)
	switchOnType(c1.name)

}

func switchOnType(x interface{}) { // empty interface
	switch x.(type) { // type conversion
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("default")

	}
}
