package main

import "fmt"

func main() {
	x := map[string][]string{
		"Kalai": {"Kanne kalai mane", "Mannil indha"},
		"Kavin": {"Bayblade bayblade", "poo pookum osai"},
		"Kuhan": {"Thunder see the thunder", "Munbe vaa"},
	}

	//adding one record in map
	x["Thiyagu"] = []string{"Thiyagu", "thiyagu"}

	for i, val := range x {
		fmt.Printf("This is the record for '%v':\n", i)
		for j, val2 := range val {
			fmt.Println("\t", j, val2)
		}
	}
	fmt.Println("******************************")
	//delete a record
	delete(x, "Kalai")
	for i, val := range x {
		fmt.Printf("This is the record for '%v':\n", i)
		for j, val2 := range val {
			fmt.Println("\t", j, val2)
		}
	}
}
