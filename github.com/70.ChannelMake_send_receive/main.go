package main

import "fmt"

func main() {
	c := make(chan int)    //bidirectional,both send and recieve
	cs := make(chan<- int) //send onto the channel, read it from left to right
	cr := make(<-chan int) //recieve only channel, read it from left to right

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)

	//general to specific conversion
	cs = c
	fmt.Printf("send only to bidirectional: %T\n", cs)

	fmt.Printf("c\t%T\n", (chan<- int)(c))

	cr = c
	fmt.Printf("send only to bidirectional: %T\n", cr)
	fmt.Printf("c\t%T\n", (<-chan int)(c))

	//specific to general(bidirectional c) assignment won't work
	//c=cs //won't work
	//c=cr //won't work
	//	fmt.Printf("c\t%T\n", (chan int)(cs)) //won't work
	//	fmt.Printf("c\t%T\n", (chan int)(cr)) //won't work

}
