package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64
	const gs = 100
	var ws sync.WaitGroup
	ws.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			ws.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())

	}
	ws.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
