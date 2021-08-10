package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Age   int
}

func main() {

	p1 := Person{
		First: "Kalai",
		Age:   37,
	}

	p2 := Person{
		First: "Zebra",
		Age:   39,
	}

	people := []Person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
