package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := append(slice1, 5, 6)
	fmt.Println(slice1)
	fmt.Printf("%T\n", slice1)
	fmt.Println(slice2)
	fmt.Printf("%T\n", slice2)

	s1 := []int{1, 2, 3, 4}
	s1 = append(s1, 9, 9)
	fmt.Println("s1", s1)
	s2 := []int{10, 20, 30}
	s1 = append(s1, s2...) //adding all the values of s2 in s1
	fmt.Println("s1", s1)

}
