package main

import "fmt"

//we can return function, assign a function to a variable, and can pass function to other function as a parameter.
//callback meaning, passing function as an argument
func main() {
	x := []int{10, 15, 3, 2, 20}

	s := sum(x...)
	fmt.Println("sum", s)

	s1 := even(sum, x...)
	fmt.Println("sum of the even number in slice x", x, ":", s1)

	s2 := odd(sum, x...)
	fmt.Println("sum of the odd number in slice x ", x, ":", s2)

}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	//fmt.Println(total)
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			fmt.Println("even numbers:", v)
			y = append(y, v)
		}
	}
	fmt.Println("y[]int", y)
	return f(y...) //passing slice y to the callback function in even method.
}

func odd(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%3 == 0 {
			fmt.Println("Odd numbers", v)
			y = append(y, v)
		}
	}
	return f(y...) //passing slice y to the callback function in even method.
}
