package lemin

import (
	"fmt"
	"os"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      string
	Adjacent []*Vertex
}

var MAX_CAN_HUNDLE = 10000000

// adds an vertex to the graph [ O ]
func (g *Graph) AddVertex(k string) error {
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
func (g *Graph) AddEdge(from, to string) error {
	// get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v<--->%v) ", from, to)
		return err
	} else if contains(fromVertex.Adjacent, to) {
		err := fmt.Errorf("existing edge (%v<--->%v) ", from, to)
		return err
	} else if fromVertex.key == toVertex.key {
		err := fmt.Errorf("same vertex (%v<--->%v) ", from, to)
		return err
	} else {
		// add edge bitween us
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
		toVertex.Adjacent = append(toVertex.Adjacent, fromVertex)
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
		for _, v := range val.Adjacent {
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
		for _, vert := range start.Adjacent {
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

// ////////////////////////////// for big graph
func (g *Graph) BFS(path *[][]string, start *Vertex, end string) {
	// Queue stores the current path being explored
	queue := [][]*Vertex{{start}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		// Dequeue the first path
		currentPath := queue[0]
		queue = queue[1:]

		// Get the last vertex in the path
		currentVertex := currentPath[len(currentPath)-1]

		// If we have reached the end vertex, save the path
		if currentVertex.key == end {
			var pathStr []string
			for _, v := range currentPath {
				pathStr = append(pathStr, v.key)
			}
			*path = append(*path, pathStr)
			continue
		}

		// Mark the vertex as visited
		visited[currentVertex.key] = true

		// Explore all Adjacent vertices
		for _, AdjacentVertex := range currentVertex.Adjacent {
			if !visited[AdjacentVertex.key] {
				// Create a new path extending the current path
				newPath := append([]*Vertex{}, currentPath...)
				newPath = append(newPath, AdjacentVertex)
				queue = append(queue, newPath)
			}
		}
	}
}

// func (g *Graph) BFS(start *Vertex, end string) [][]string {
// 	var allPaths [][]string
// 	queue := [][]string{{start.key}}

// 	for len(queue) > 0 {
// 		path := queue[0]
// 		queue = queue[1:]
// 		lastVertex := path[len(path)-1]

// 		if lastVertex == end {
// 			allPaths = append(allPaths, path)
// 		} else {
// 			currentVertex := g.GetVertex(lastVertex)
// 			for _, neighbor := range currentVertex.Adjacent {
// 				if !containsVertex(path, neighbor.key) {
// 					newPath := make([]string, len(path))
// 					copy(newPath, path)
// 					newPath = append(newPath, neighbor.key)
// 					queue = append(queue, newPath)
// 				}
// 			}
// 		}
// 	}

// 	return allPaths
// }

// func containsVertex(path []string, key string) bool {
// 	for _, v := range path {
// 		if v == key {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (g *Graph) MaxFlow(start, end string) int {
// 	startVertex := g.GetVertex(start)
// 	if startVertex == nil {
// 		return 0 // Or handle error
// 	}

// 	var paths [][]string
// 	stack := []string{}
// 	g.DFS(&paths, &stack, startVertex, end)

// 	// Find edge-disjoint paths
// 	maxFlow := 0
// 	usedEdges := make(map[string]bool)

// 	for _, path := range paths {
// 		isDisjoint := true
// 		for i := 0; i < len(path)-1; i++ {
// 			edge := path[i] + "-" + path[i+1]
// 			reverseEdge := path[i+1] + "-" + path[i]
// 			if usedEdges[edge] || usedEdges[reverseEdge] {
// 				isDisjoint = false
// 				break
// 			}
// 		}

// 		if isDisjoint {
// 			maxFlow++
// 			for i := 0; i < len(path)-1; i++ {
// 				edge := path[i] + "-" + path[i+1]
// 				reverseEdge := path[i+1] + "-" + path[i]
// 				usedEdges[edge] = true
// 				usedEdges[reverseEdge] = true
// 			}
// 		}
// 	}

// 	return maxFlow
// }
// //================================================
// func (g *Graph) FindAllPathsAndMaxFlow(source, sink string) ([][]string, int) {
// 	// List to store all paths
// 	var allPaths [][]string

// 	// Track visited vertices and edges
// 	visitedEdges := make(map[string]map[string]bool)

// 	// Initialize visitedEdges to track edges used in paths
// 	for _, v := range g.vertices {
// 		visitedEdges[v.key] = make(map[string]bool)
// 	}

// 	// Run DFS to find all paths
// 	for {
// 		stack := []*Vertex{}
// 		path := []string{}
// 		visitedVertices := make(map[string]bool)

// 		// If no new path is found, break
// 		if !g.DDFS(&path, &stack, g.GetVertex(source), sink, visitedVertices, visitedEdges) {
// 			break
// 		}

// 		// Add the found path to the list of all paths
// 		allPaths = append(allPaths, path)

// 		// Mark edges in the path as used
// 		for i := 0; i < len(path)-1; i++ {
// 			from := path[i]
// 			to := path[i+1]
// 			visitedEdges[from][to] = true
// 			visitedEdges[to][from] = true // Since it's undirected
// 		}
// 	}

// 	// Maximum flow is the number of disjoint paths
// 	maxFlow := len(allPaths)
// 	return allPaths, maxFlow
// }

// // DFS function to find paths considering visited vertices and edges
// func (g *Graph) DDFS(path *[]string, stack *[]*Vertex, start *Vertex, end string, visitedVertices map[string]bool, visitedEdges map[string]map[string]bool) bool {
// 	*stack = append(*stack, start)
// 	*path = append(*path, start.key)
// 	visitedVertices[start.key] = true

// 	// If the end vertex is reached, return true
// 	if start.key == end {
// 		return true
// 	}

// 	// Explore Adjacent vertices
// 	for _, neighbor := range start.Adjacent {
// 		// Check if the edge has been used in a previous path
// 		if !visitedEdges[start.key][neighbor.key] && !visitedVertices[neighbor.key] {
// 			if g.DDFS(path, stack, neighbor, end, visitedVertices, visitedEdges) {
// 				return true
// 			}
// 		}
// 	}

// 	// Backtrack if no path is found
// 	*stack = (*stack)[:len(*stack)-1]
// 	*path = (*path)[:len(*path)-1]
// 	return false
// }

//qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq

func (g *Graph) DFSIterative(start, end *Vertex) [][]string {
	var result [][]string
	stack := [][]*Vertex{{start}}
	visited := make(map[string]bool)

	for len(stack) > 0 {
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		vertex := path[len(path)-1]
		if visited[vertex.key] {
			continue
		}

		visited[vertex.key] = true

		if vertex.key == end.key {
			var currentPath []string
			for _, v := range path {
				currentPath = append(currentPath, v.key)
			}
			result = append(result, currentPath)
		} else {
			for _, adj := range vertex.Adjacent {
				if !visited[adj.key] {
					newPath := append([]*Vertex{}, path...)
					newPath = append(newPath, adj)
					stack = append(stack, newPath)
				}
			}
		}
		visited[vertex.key] = false // Unmark the vertex to allow other paths
	}
	return result
}

// [[[[[[[[[[[[[[[[[[[[[[[[[]]]]]]]]]]]]]]]]]]]]]]]]]
var n int

func (g *Graph) DFSS(path *[][]string, stack []string, start *Vertex, end string, visited map[string]bool) {
	n++
	if n >= MAX_CAN_HUNDLE {
		fmt.Println("this graph is soo big")
		os.Exit(0)
	}

	stack = append(stack, start.key)
	visited[start.key] = true

	if start.key == end {

		currentPath := append([]string{}, stack...)
		*path = append(*path, currentPath)

	} else {
		for _, vert := range start.Adjacent {
			if !visited[vert.key] {
				g.DFSS(path, stack, vert, end, visited)
			}
		}
	}

	// backtrack
	stack = (stack)[:len(stack)-1]
	visited[start.key] = false
}
