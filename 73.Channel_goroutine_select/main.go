package main

import "fmt"

func main() {
	even := make(chan int)

	odd := make(chan int)

	quit := make(chan int)

	//send
	go foo(even, odd, quit)

	//receive
	receive(even, odd, quit)

	fmt.Println("Exit")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel", v)
		case v := <-odd:
			fmt.Println("from the Odd channel", v)
		case v := <-quit:
			fmt.Println("from the quit channel", v)
			return
		}
	}
}

func foo(even, odd, quit chan<- int) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0

}
