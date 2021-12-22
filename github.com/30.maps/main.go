package main

import "fmt"

func main() {

	//map type1
	x := make(map[string]int)
	x["H2O"] = 100
	fmt.Println(x, len(x))
	fmt.Println("*************************************")

	//map type2
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
	}

	//adding element in the map
	m["K"] = "Kite"
	fmt.Println(m)

	fmt.Println("*************************************")

	el := map[string]map[string]string{
		"flower": {
			"name": "rose",
			"age":  "20",
		},
		"gadjets": {
			"name": "TV",
			"age":  "10",
		},
	}
	fmt.Println(el)

	if ele, ok := el["gadjets"]; ok {
		fmt.Println(ele["name"], ele["age"])
	}
	if ele, ok := el["flower"]["name"]; ok {
		fmt.Println(ele)
	}
}
