package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{
		First: "Kalai",
		Age:   57,
	}

	p2 := Person{
		First: "Apple",
		Age:   7,
	}

	p3 := Person{
		First: "Zebra",
		Age:   47,
	}

	people := []Person{p1, p2, p3}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

}
