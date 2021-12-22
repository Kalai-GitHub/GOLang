package main

import "fmt"

func main() {
	xi := []int{4, 3, 2, 1, 0}
	v := sum(xi...)
	fmt.Println(v)
	fmt.Println("************************************")
	p := sum()
	fmt.Println(p)
	fmt.Println("************************************")
	s, v1 := sum1("James", 4, 5, 6)
	fmt.Println(s, v1)

}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("index is: ", i, "value is: ", v, sum)
	}
	return sum
}

func sum1(s string, x ...int) (string, int) {
	fmt.Println(s, x)
	fmt.Printf("%T\n%T\n", s, x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("index is: ", i, "value is: ", v, sum)
	}
	return s, sum
}
