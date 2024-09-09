package main

import (
	"fmt"
	"strconv"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      string
	adjacent []*Vertex
}

// adds an vertex to the graph [ O ]
func (g *Graph) AddVertex(k string) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added it is an existing key ", k)
		fmt.Println(err.Error())
	} else {
		// add vertex
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// adds an edge to the graph [ -- ]
func (g *Graph) AddEdge(from, to string) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v--->%v) ", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v--->%v) ", from, to)
		fmt.Println(err.Error())
	} else if fromVertex.key == toVertex.key {
		err := fmt.Errorf("same vertex (%v--->%v) ", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge bitween us
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		toVertex.adjacent = append(toVertex.adjacent, fromVertex)
	}
}

// returns a pointer to the Vertex whith a key int
func (g *Graph) getVertex(k string) *Vertex {
	for i, val := range g.vertices {
		if val.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// check vertex key if alredy exist
func contains(s []*Vertex, k string) bool {
	for _, val := range s {
		if k == val.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, val := range g.vertices {
		fmt.Printf("\nVertex %v :", val.key)
		for _, v := range val.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	print("\n")
}

func (g *Graph) DFS(start, k string) []*Vertex {
	stack := []*Vertex{}
	for _, vert := range g.vertices {
		fmt.Println(vert.key)
		for _, aj := range vert.adjacent {
			if aj.key == k {
				stack = append(stack, vert)
				return stack
			} else if !Visited(stack, aj) {
				stack = append(stack, aj)
			}
		}
	}
	return nil
}

func Visited(stack []*Vertex, visit *Vertex) bool {
	for _, val := range stack {
		if visit.key == val.key {
			return true
		}
	}
	return false
}

func main() {
	test := &Graph{}

	for i := 0; i <= 7; i++ {
		test.AddVertex(strconv.Itoa(i))
	}
	test.AddVertex("h")

	test.AddEdge("0", "6")
	test.AddEdge("1", "3")
	test.AddEdge("4", "3")
	test.AddEdge("5", "2")
	test.AddEdge("3", "5")
	test.AddEdge("4", "2")
	test.AddEdge("2", "1")
	test.AddEdge("7", "6")
	test.AddEdge("7", "2")
	test.AddEdge("7", "4")
	test.AddEdge("6", "4")
	test.AddEdge("h", "4")

	test.Print()

	serch := test.DFS("1", "h")

	for _, val := range serch {
		fmt.Println("|", val.key, "|", val.adjacent, "|")
	}
}
