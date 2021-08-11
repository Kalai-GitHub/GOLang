package main

func main() {
	// c := make(<-chan int, 1) //receive only channel

	// c <- 33 //cant send value onto the channel, if the channel is receive only.

	// fmt.Println(<-c) //invalid operation: cannot receive from send-only channel c (variable of type chan<- int)compilerInvalidReceive

}
