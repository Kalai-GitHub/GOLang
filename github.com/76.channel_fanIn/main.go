package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("Exit")
}

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

//fan in - Taking values from many channels, and putting those values onto one channel.
func receive(even, odd <-chan int, fanIn chan<- int) {
	var ws sync.WaitGroup
	ws.Add(2)

	//fanIn receives the value from even and odd[interlocking coordinated pieces]
	go func() {
		for v := range even {
			fanIn <- v
		}
		ws.Done()
	}()

	go func() {
		for v := range odd {
			fanIn <- v
		}
		ws.Done()
	}()
	ws.Wait()
	close(fanIn)

}
