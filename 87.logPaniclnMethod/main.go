package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("names.txt")
	if err != nil {
		log.Panicln(err) //2021/08/12 13:47:13 open names.txt: no such file or directory, panic: open names.txt: no such file or directory //goroutine 1 [running]: //exit status 2
	}

	/* Panicln is equivalent to Println() followed by a call to panic(). this will be prevent by the recover()*/
}
