package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//identifier needs to start with upper case for json marshal
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Kalai",
		Last:  "murugu",
		Age:   37,
	}

	p2 := Person{
		First: "Kavin",
		Last:  "Thiyagu",
		Age:   10,
	}
	people := []Person{p1, p2}
	fmt.Println(people)

	//to convert struct to json format
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
