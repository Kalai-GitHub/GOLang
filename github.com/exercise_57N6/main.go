package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)
	for k := 0; k < x*y; k++ {
		fmt.Println(k, <-c)
	}
}

func gen(x, y int) <-chan int {
	c := make(chan int)
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
			// close(c)
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	return c
}
