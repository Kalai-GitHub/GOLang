package main

import "fmt"

//Graph represents an Adjacency list graph
type Graph struct {
	vertices []*Vertex
}

//Vertex represents a graph vertex
type Vertex struct {
	Key      int
	adjacent []*Vertex
}

//AddVertex method adds a vertex to the Graph
func (g *Graph) AddVertex(k int) {
	if !
	g.vertices = append(g.vertices, &Vertex{Key: k})
}

//contains
func contains(s []*Vertex, k int) bool{
	for _,v := range s{
		if k == v.Key{
			
		}
	}
}

//Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.Key)
		for _, v := range v.adjacent {
			fmt.Printf("%v ", v.Key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	fmt.Println("test", test)
	test.Print()
}
