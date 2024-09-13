package main

import (
	"fmt"
	"strings"

	box "box/parseFile"
	// box "box/parseFile"
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

func (g *Graph) DFS(path *[][]string, stack *[]string, start *Vertex, end string) {
	*stack = append(*stack, start.key)

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

// sort parhs
func sortPaths(path [][]string) [][]string {
	n := len(path)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if len(path[j+1]) < len(path[j]) {
				path[j+1], path[j] = path[j], path[j+1]
			}
		}
	}
	return path
}

//return special paths
func (g *Graph) greedy(paths [][]string) [][]string {
	var way string
	var res [][]string
	for _, a := range paths {
		az := a[1 : len(a)-1]
		a1 := strings.Join(az, " ")
			if !is(a1, way) {
				way += a1
				res = append(res, a)
		}
	}

	return res
}

func is(a, b string) bool {
	for _, va := range a {
		for _, vb := range b {
			if (vb != ' ' || va != ' ') && vb == va {
				return true
			}
		}
	}
	return false
}

func main() {
	graph := &Graph{}

	farms, err := box.ParseFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range farms.Rooms {
		graph.AddVertex(val)
	}
	for _, val := range farms.Links {
		valsp := strings.Split(val, "-")
		graph.AddEdge(valsp[0], valsp[1])
	}

	graph.Print()
	println("****************************************")
	start := graph.getVertex(farms.Start)
	if start == nil {
		fmt.Println("cant find vertex to (start)")
		return
	}
	end := graph.getVertex(farms.End)
	if end == nil {
		fmt.Println("cant find vertex to (end)")
		return
	}
	paths := [][]string{}
	stack := []string{}

	graph.DFS(&paths, &stack, start, farms.End)
	if len(paths) == 0 {
		fmt.Println("not found end")
		return
	}
	pathsort := sortPaths(paths)

	for _, val := range pathsort {
		fmt.Println(val)
	}
	way := graph.greedy(pathsort)
	println("*************************************************")
	for _, val := range way {
		fmt.Println(val)
	}
}
