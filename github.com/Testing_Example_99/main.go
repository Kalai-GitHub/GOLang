//package Testing99 asks to get a job from FANG company
package Testing_Example_99

//sun function adds unlimited number of values of type int
func sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
