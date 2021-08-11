package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Atomic is to prevent race condition as Mutex.
func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	var counter int64 //int64 is from Atomic package

	const goroutiness = 100
	var ws sync.WaitGroup
	ws.Add(goroutiness)

	for i := 0; i < goroutiness; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//time.Sleep(time.Second)
			runtime.Gosched() //Gosched yields the processor allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.
			ws.Done()

		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}

	ws.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("count :", counter)

}
