package main

import "fmt"

func main() {
	//array
	arr := [5]int{10, 20, 30, 40, 50}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
	fmt.Println("***********************************")

	//slice
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", slice)

	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])

	slice = append(slice, 52)
	fmt.Println(slice)
	slice = append(slice, 53, 54, 55)
	fmt.Println(slice)

	x := []int{56, 57, 58, 59, 60}
	slice = append(slice, x...)
	fmt.Println(slice)

	fmt.Println("***********************************")

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y = append(y[:3], y[6:]...)
	fmt.Println(y)

}
