package main

import "fmt"

func main() {
	oneFeetInMeter := 0.3048
	var feet float64
	fmt.Print("Enter feet value to convert into meter: ")
	fmt.Scanf("%f", &feet)

	meter := feet * oneFeetInMeter
	fmt.Printf("%f feet is %f meter \n", feet, meter)

}
