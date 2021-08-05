package main

import "fmt"

//composite used to hold different type of values
//aggregate data type is used to store the same type values in array or slice
func main() {
	//slice type 1
	var s1 []float64
	fmt.Println(s1)

	//slice type 2
	s2 := []float64{0, 1, 2, 3, 4, 5}
	fmt.Println("s2", s2)
	slice1 := s2[:]
	fmt.Println("slice1", slice1)
	slice2 := s2[1:3]
	fmt.Println("slice2", slice2)
	slice3 := s2[:3]
	fmt.Println("slice3", slice3)

	// //slice type 3 => not working
	// s2 := []float64{
	// 	4,
	// 	5,
	// 	6,
	// }
}
