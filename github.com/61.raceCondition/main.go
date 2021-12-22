package main

import (
	"fmt"
	"runtime"
	"sync"
)

//RaceCondition - loop is executing parallely so called race condition.
func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0

	const goroutiness = 100
	var ws sync.WaitGroup
	ws.Add(goroutiness)

	for i := 0; i < goroutiness; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() //Gosched yields the processor allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			v++
			counter = v
			ws.Done()

		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}

	ws.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count :", counter)

}
