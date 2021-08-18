//Stacks worked as LIFO, and has two methods, Push and Pop
package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) pop() int {
	top := len(s.items) - 1
	toRemove := s.items[top]
	s.items = s.items[:top]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.push(10)
	myStack.push(20)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)

}
