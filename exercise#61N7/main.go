package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//se := errors.New("square root of negative number is invalid") // errors.New will accept only string
		se := fmt.Errorf("square root of negative number is invalid %v ", f) //fmt.Errorf will accept string and value
		return 0, sqrtError{"11.110 ", "32.12 ", se}
	}
	return 42, nil
}
