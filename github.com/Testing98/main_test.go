package main

import "testing"

func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 2, 4}, 7},
		{[]int{4, 4}, 8},
		{[]int{9, 0}, 9},
	}
	for _, v := range tests {
		x := sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
