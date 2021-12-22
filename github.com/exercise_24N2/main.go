package main

import "fmt"

func main() {
	a := struct {
		firstName string
		friends   map[string]int
		favMovie  []string
	}{
		firstName: "Kavin",
		friends: map[string]int{
			"JJ": 10,
			"KK": 11,
		},
		favMovie: []string{
			"star wars",
			"avengers",
		},
	}

	fmt.Println(a.firstName)
	fmt.Println(a.friends)
	fmt.Println(a.favMovie)

	for i, v := range a.friends {
		fmt.Println(i, v)
	}

	for k, v := range a.favMovie {
		fmt.Println(k, v)
	}
}
