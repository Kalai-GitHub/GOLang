package main

import (
	"fmt"
	"runtime"
	"sync"
)

//mutex is for locking and unlocking the chunk of codes to prevent race condition
func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0

	const goroutiness = 100
	var ws sync.WaitGroup
	ws.Add(goroutiness)

	var mu sync.Mutex
	for i := 0; i < goroutiness; i++ {
		go func() {
			mu.Lock() //it is for reading and writing, if someone uses this code, others cannot use this same code at the same time until the execution is complete.
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() //Gosched yields the processor allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			v++
			counter = v
			mu.Unlock()
			ws.Done()

		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}

	ws.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count :", counter)

}
