package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("names.txt")
	if err != nil {
		panic(err) //panic: open names.txt: no such file or directory //goroutine 1 [running]: //exit status 2 ....//panic function will stop the normal execution of the current goroutine.
	}

	/* panic function call Panic after writing the log message. executes the defer functions */
}

func foo() {
	fmt.Println("defer foo")
}
