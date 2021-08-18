package main

import "fmt"

func main() {
	var arr [58]string
	fmt.Println(arr)

	for i := 65; i <= 122; i++ {
		arr[i-65] = string(rune(i)) //converting int to string
	}
	fmt.Println(arr)
	fmt.Println(arr[25])

	for i, v := range arr {
		fmt.Printf("Value :%v\tType : %T\t%v\n", v, v, []byte(v))
		if i > 10 {
			break
		}
	}

}
