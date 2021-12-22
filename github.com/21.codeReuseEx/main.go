package main

import (
	"fmt"

	"github.com/00.modules/greetings"
)

func main() {
	fmt.Println("Enter ur name : ")
	var str string
	fmt.Scanf("%s\n", &str)
	msg := greetings.Hello(str)
	fmt.Println(msg)

}
