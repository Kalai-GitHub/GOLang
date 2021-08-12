package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error Happened ", err) //Error Happened  open names.txt: no such file or directory
		//log.Println("Error Happened ", err) //2021/08/12 13:37:33 Error Happened  open names.txt: no such file or directory
		//log.Fatalln("Error Happened ", err) //2021/08/12 13:37:49 Error Happened  open names.txt: no such file or directory //exit status 1
		//log.Panicln(err) //2021/08/12 13:47:13 open names.txt: no such file or directory, panic: open names.txt: no such file or directory //goroutine 1 [running]: //exit status 2
		//panic(err) //panic: open names.txt: no such file or directory //goroutine 1 [running]: //exit status 2 ....//panic function will stop the normal execution of the current goroutine.
	}

	/* fatal function call os.Exit(1) after writing the log message ...*/
	//Fatalln is equivalent to Println() followed by a call to os.Exit(1)
}
