package main

import "testing"

func TestAverage(t *testing.T) {
	if average(1, 2) != 1.5 {
		t.Error("Expected", 1.5, "Got", average(1, 2))
	}
}
