package main

import "fmt"

func main() {
	fmt.Print("Enter the fahrenheit to convert Celsius : ")
	var Fvalue float64
	fmt.Scanf("%f", &Fvalue)
	celsius := (Fvalue - 32) * 5 / 9
	fmt.Println("celsius value :", celsius)
}
