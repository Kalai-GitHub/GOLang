package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup //wg is a type from WaitGroup and from package sync

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	wg.Add(1) //meaning have to wait for one thing[foo function execution to complete]
	go foo()
	bar()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine()) //we can see the increase in the number of goroutines because of [go foo()]
	wg.Wait()                                         //it will wait, to complete the foo function execution to exit[meaning all the things are added in[wg.Add(1)] to be done]
}

func foo() {
	fmt.Println("foo:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	wg.Done() //it allows to execute the foo function to complete
}

func bar() {
	fmt.Println("bar:")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
