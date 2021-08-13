package main

import (
	"errors"
	"fmt"
	"log"
)

var errorMsg = errors.New("Negative value can not be passed. Square root of negative number")

func main() {
	fmt.Printf("%T\n", errorMsg) //*errors.errorString
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err) // Fatalln will be used to exit from the program
	}

}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errorMsg
	}
	return 32, nil
}
