package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var ws sync.WaitGroup
	ws.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched() //Gosched yields the processor allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			v++
			counter = v
			ws.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}
	ws.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
