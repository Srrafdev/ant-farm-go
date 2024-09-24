package lemin

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      string
	adjacent []*Vertex
}

// adds an vertex to the graph [ O ]
func (g *Graph) AddVertex(k string)error {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added it is an existing key ", k)
		return err
	} else {
		// add vertex
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
	return nil
}

// adds an edge to the graph [ -- ]
func (g *Graph) AddEdge(from, to string)error {
	// get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v--->%v) ", from, to)
		return err
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v--->%v) ", from, to)
		return err
	} else if fromVertex.key == toVertex.key {
		err := fmt.Errorf("same vertex (%v--->%v) ", from, to)
		return err
	} else {
		// add edge bitween us
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		toVertex.adjacent = append(toVertex.adjacent, fromVertex)
	}
	return nil
}

// returns a pointer to the Vertex whith a key int
func (g *Graph) GetVertex(k string) *Vertex {
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
