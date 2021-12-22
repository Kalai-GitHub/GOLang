package greetings

import "fmt"

func Hello(s string) string {
	message := fmt.Sprintf("Hi, %v.Welcome", s)
	return message
}
