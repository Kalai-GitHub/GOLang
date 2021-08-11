package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)

	recieve(even, odd, quit)

	fmt.Println("Exit")
}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func recieve(even, odd <-chan int, quit <-chan bool) {

	for {
		select {
		case v := <-even:
			fmt.Println("from even channel:", v)
		case v := <-odd:
			fmt.Println("from odd channel:", v)
		case v, ok := <-quit: //comma Ok idiom will return i and ok value as boolean result[false,false] bcoz, quit is bool type. If quit is int, this code will return [0, false].
			if !ok {
				fmt.Println("from comma ok -quit", v, ok)
				return
			} else {
				fmt.Println("from comma ok - quit", v)
			}
		}
	}
}
