package Testing_Example_99

// not working - check with Selva
import "fmt"

// Example is similar to test[in terminal - go test],
//Example show up in the documentation[It checks the comments and test the code. In this code, ExampleSum checks the output of the sum function and and compares the comments in this code and produce the result]
func ExampleSum() {
	fmt.Println(sum(3, 2))
	//Output
	//5 //[if I give any other value it should give error]
}
