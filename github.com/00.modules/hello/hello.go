package main

import (
	"fmt"

	"github.com/00.modules/greetings"
)

func main() {
	msg := greetings.Hello("Kalai")
	fmt.Println(msg)
}
