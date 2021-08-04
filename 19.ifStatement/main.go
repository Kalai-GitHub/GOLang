package main

import "fmt"

func main() {
	if !false {
		fmt.Println("001")
	}
	if 2 == 2 {
		fmt.Println("002")
	}
	if 2 != 2 {
		fmt.Println("003")
	}
	if !(2 == 2) {
		fmt.Println("004")
	}
	if !(2 != 2) {
		fmt.Println("005")
	}

	if x := 42; x == 2 {
		fmt.Println("condition is true")
	}

	fmt.Println("\nmultiple of 3 and 5= FizzBuzz, multiple of 3 = Fizz, multiple of 5 = Buzz")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d\tFizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d\tFizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d\tBuzz\n", i)
		}
	}
}
