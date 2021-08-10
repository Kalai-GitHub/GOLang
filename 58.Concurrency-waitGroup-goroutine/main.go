package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	go foo()
	bar()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

}

func foo() {
	fmt.Println("foo:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func bar() {
	fmt.Println("bar:")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
