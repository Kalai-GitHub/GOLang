package main

import "fmt"

func main() {
	m := map[string]string{
		"A": "Apple",
		"B": "Ball",
		"C": "Car",
	}
	fmt.Println(m)
	fmt.Println(m["B"])
	fmt.Println(m["C"])

	i, ok := m["K"]
	fmt.Println(i, ok)

	if j, ok := m["B"]; ok {
		fmt.Println("The value is : ", j)
		//deleting an element from map m
		delete(m, "B")
	}

	//adding element in the map m
	m["K"] = "Kite"
	fmt.Println(m)
}
