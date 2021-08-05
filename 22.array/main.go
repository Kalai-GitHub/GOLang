package main

import "fmt"

//note: Array cannot allow the negative number for the size
func main() {

	var x [5]int
	x[3] = 100
	fmt.Println(x)
	fmt.Println("*************************************")
	//array declaration type 1:
	var y [5]float64
	y[0] = 90
	y[1] = 91
	y[2] = 90
	y[3] = 87
	y[4] = 89
	fmt.Println(y)
	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println("total", total/float64(len(y)))
	fmt.Println("*************************************")

	//array declaration type 2:
	z := [5]float64{80, 80, 80, 80, 80}

	var sum float64 = 0
	//range of z, i is for index and val is for value
	for i, val := range z {
		fmt.Println(i, val)
		sum += z[i]
	}
	fmt.Println("sum", sum/float64(len(z)))

	// '_' to aviod the index value in range
	for _, val := range z {
		sum += val
	}
	fmt.Println("sum1", sum/float64(len(z)))
	fmt.Println("*************************************")

	//array declaration type 3:
	//note the , sholud be placed at the end of the value
	p := [5]float64{
		1,
		2,
		3,
		4,
		5,
	}
	for i, v := range p {
		fmt.Println("P", i, v)
	}
}
