package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   // shows the running program's operating system target
	fmt.Println(runtime.GOARCH) // shows the running program's architecture target
}
