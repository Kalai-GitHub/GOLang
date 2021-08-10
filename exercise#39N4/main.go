package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Kalai","Age":37},{"First":"Zebra","Age":39}]`

	bs := []byte(s)

	var people []Person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
}
