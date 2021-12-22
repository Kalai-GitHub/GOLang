package main

import (
	"fmt"

	"github.com/example94/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	c1 := canine{
		name: "Kalai",
		age:  dog.Years(10),
	}
	fmt.Println(c1.age)
}
