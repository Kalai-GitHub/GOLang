package main

import "fmt"

func main() {
	type person struct {
		firstName               string
		lastName                string
		favoriteIcecreamFlavors []string
	}

	p1 := person{
		firstName:               "Kuhan",
		lastName:                "Thiyagu",
		favoriteIcecreamFlavors: []string{"Vannila", "strawberry"},
	}

	p2 := person{
		firstName:               "Kavin",
		lastName:                "Thiyagarajan",
		favoriteIcecreamFlavors: []string{"chocolate", "Pistacho"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.favoriteIcecreamFlavors {
		fmt.Println(i, v)
	}

	x := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println("Index:", i)
		fmt.Println("First Name:", v.firstName)
		for j, v1 := range v.favoriteIcecreamFlavors {
			fmt.Println(j, v1)
		}
	}

}
