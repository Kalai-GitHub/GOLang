package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

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
	sort.Sort(ByName(people))
	fmt.Println(people)

}
