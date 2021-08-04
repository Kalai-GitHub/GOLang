package main

import "fmt"

func main() {
	i := 1
	if i > 10 {
		fmt.Println("Big")
		//i++ - not working
	} else {
		fmt.Println("Small")
	}

	fmt.Println("\nPrinting divisible by 3 numbers :")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
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
