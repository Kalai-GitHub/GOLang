package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin CPUs", runtime.NumCPU())
	fmt.Println("begin Goroutines", runtime.NumGoroutine())

	var ws sync.WaitGroup

	ws.Add(2)
	go func() {
		fmt.Println("Anything from cat1")
		ws.Done()
	}()

	go func() {
		fmt.Println("Anything from cat2")
		ws.Done()
	}()
	fmt.Println("mid CPUs", runtime.NumCPU())
	fmt.Println("mid Goroutines", runtime.NumGoroutine())

	ws.Wait()
	fmt.Println("end CPUs", runtime.NumCPU())
	fmt.Println("end Goroutines", runtime.NumGoroutine())

	fmt.Println("Exit")

}
