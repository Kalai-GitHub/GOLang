package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("names.txt")
	if err != nil {

		log.Fatalln("Error Happened ", err) //2021/08/12 13:37:49 Error Happened  open names.txt: no such file or directory //exit status 1
	}

	/* fatal function call os.Exit(1) after writing the log message ...*/
	//Fatalln is equivalent to Println() followed by a call to os.Exit(1)
}

func foo() {
	fmt.Println("Foo: when os.Exit() is called, deferred foo function does not run")
}
