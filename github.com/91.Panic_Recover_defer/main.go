package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f") //[eighth]
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r) //4[seventh]
		}
	}()
	fmt.Println("Calling g") //first
	g(0)
	fmt.Println("returned normally from g")

}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking...") //[fourth]
		panic(fmt.Sprintf("%v", i)) //[fifth] but move to defer and pass the value 5 to the recover function
	}
	defer fmt.Println("Defer in g", i) //3,2,1,0 [sixth]
	fmt.Println("Printin in g", i)     //0,1,2,3[second]
	g(i + 1)                           //1,2,3,4[third]
}
