package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"` //upper case is for export the property as a public
	Last  string `json:"Last"`
	//Age   int    `json:"Age"` // this wont display any value due to int value is not passed by the input string.
	StrAge string `json:"StrAge"`
}

func main() {
	//p := `[{"First":"Kalai","Last":"murugu","Age":37},{"First":"Kavin","Last":"Thiyagu","Age":10}]`
	p := `[{"First":"Kalai","Last":"murugu","StrAge":"37"},{"First":"Kavin","Last":"Thiyagu","StrAge":"10"}]`

	bs := []byte(p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", bs)

	var people []Person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Println("Person Number ", i)
		//fmt.Println("\t", v.First, v.Last, v.Age)
		fmt.Println("\t", v.First, v.Last, v.StrAge)
	}
}
