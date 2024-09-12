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

func (g *Graph) DFS(path *[][]string, stack *[]string, start *Vertex,end string) {
    if len(*stack) == 0 || (*stack)[len(*stack)-1] != start.key {
        *stack = append(*stack, start.key)
    }

    if start.key == end {

        currentPath := []string{}
		currentPath = append(currentPath, *stack...)
        *path = append(*path, currentPath) 

    } else {
        for _, vert := range start.adjacent {
            if !Visited(*stack, vert) {
                g.DFS(path, stack, vert, end)
            }
        }
    }

    // backtrack
    *stack = (*stack)[:len(*stack)-1]
}



func Visited(stack []string, visit *Vertex) bool {
	for _, val := range stack {
		if visit.key == val {
			return true
		}
	}
	return false
}

func main() {
	test := &Graph{}
	fileExample := "example.txt"
	_, err := ReadFile(fileExample)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= 3; i++ {
		test.AddVertex(strconv.Itoa(i))
	}

	test.AddEdge("0", "2")
	test.AddEdge("0", "3")
	test.AddEdge("2", "1")
	test.AddEdge("3", "1")
	test.AddEdge("2", "3")

	test.Print()
	path := [][]string{}
	stack := []string{}
	test.DFS(&path, &stack, test.getVertex("0"), "1")
	fmt.Println(path)
}
