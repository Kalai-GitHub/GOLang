package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		//fmt.Println(err)
		//return

		//log.Println(err)
		//return

		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return []byte{}, fmt.Errorf("There was an error : %v", err) // this works
		return []byte{}, errors.New(fmt.Sprintf("There was an error : %v", err)) //fmt.Sprintf returns a string and errors.New will accept only a string. this works

		//return []byte{}, fmt.Println("There was an error :", err) // this wont work, bcoz, the second return type needs a error

	}
	return bs, nil
}
