package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgateMath error has occured: %v%v%v", n.lat, n.long, n.err)
}
func main() {
	_, err := sqrt(-10.32)
	if err != nil {
		log.Println(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number is invalid")
		return 0, norgateMathError{"11.110 ", "32.12 ", nme}
	}
	return 32, nil
}
