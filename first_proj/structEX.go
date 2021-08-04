package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var me Person
	me.name = "Kalai"
	me.age = 35

	var husband Person
	husband.name = "Thiyagu"
	husband.age = 40

	var son1 Person
	son1.name = "Kavin"
	son1.age = 9

	var son2 Person
	son2.name = "Kuhan"
	son2.age = 5

	birthday(&son1)
	fmt.Println("son1 age", son1.age)
	birthday(&son2)
	fmt.Println("son2 age", son2.age)
}

func birthday(person *Person) {
	person.age += 1
	fmt.Println(person.age)
}
