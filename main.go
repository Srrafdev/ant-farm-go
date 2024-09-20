package main

import (
	"fmt"
	"strconv"
	"strings"

	box "box/parseFile"
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

// return special paths
func greedy(paths [][]string) [][]string {
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
		fmt.Println("ERROR: invalid data format: ", err)
		return
	}

	for _, val := range farms.Rooms {
		graph.AddVertex(val)
	}
	for _, val := range farms.Links {
		valsp := strings.Split(val, "-")
		graph.AddEdge(valsp[0], valsp[1])
	}

	// fmt.Sscanf(line,"%s %d%v%d\n", &room , &x ,&y)

	graph.Print()
	println("****************************************")

	start := graph.getVertex(farms.Start)

	paths := [][]string{}
	stack := []string{}
	graph.DFS(&paths, &stack, start, farms.End)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format: not found end")
		return
	}

	paths = choosePaths(paths)
	for _, v := range paths {
		fmt.Println(v)
	}
	println("****************************************")
	move(paths, farms.NumberAnts)
	// AntsWalk(sortPaths(paths), farms.NumberAnts)
	// move(paths, farms.NumberAnts)
	// d := distributeDivision(paths, farms.NumberAnts)
	// all := 1
	// for i, path := range paths {
	// 	walk(path, d[i], all)
	// 	all += d[i]
	// }
}

func choosePaths(paths [][]string) [][]string {
	rating := rate(paths)

	SortByRate(paths, rating)

	paths = choose(paths)
	return paths
}

func choose(paths [][]string) [][]string {
	filter := [][]string{}
	m := make(map[string]bool)
	for i, path := range paths {

		for j := 1; j < len(path)-1; j++ {
			if m[paths[i][j]] {
				goto next // go to next index
			}
		}

		for j := 1; j < len(paths[i])-1; j++ {
			m[paths[i][j]] = true
		}
		filter = append(filter, paths[i])

	next:
	}
	return filter
}

func SortByRate(paths [][]string, rating []int) {
	for i := 0; i < len(paths); i++ {
		for j := i + 1; j < len(paths); j++ {
			if rating[i] > rating[j] {
				rating[i], rating[j] = rating[j], rating[i]
				paths[i], paths[j] = paths[j], paths[i]
			}
		}
	}
}

func rate(paths [][]string) []int {
	repetRom := make(map[string]int)
	for _, path := range paths {
		for _, rom := range path {
			repetRom[rom]++
		}
	}

	rating := make([]int, len(paths))
	for i, path := range paths {
		for _, rom := range path {
			rating[i] += repetRom[rom]
		}
	}

	return rating
}

func AntsWalk(paths [][]string, ants int) {
	var varince int
	// numb := make([]int, len(paths)-1)
	for _, path := range paths {
		if len(path)+ants < varince {
		}
	}
}

func distributeDivision(paths [][]string, total int) []int {
	n := len(paths)
	maxLen := 0

	// Find the maximum path length
	for _, path := range paths {
		if n > maxLen {
			maxLen = len(path)
		}
	}

	result := make([]int, n)
	remainingTotal := total

	// Calculate the initial division for each path
	for i, path := range paths {
		difference := maxLen - len(path)
		result[i] = difference
		remainingTotal -= difference
	}

	// Distribute any remaining total
	for i := 0; i < remainingTotal; i++ {
		result[i%n]++
	}

	return result
}

func walk(path []string, n, all int) {
	var walkAnt []string
	var res [][]string
	number := all
	for number <= n+all-1 {
		for _, room := range path[1:] {
			walkAnt = append(walkAnt, "L"+strconv.Itoa(number)+"-"+room+" ")
		}
		res = append(res, walkAnt)
		walkAnt = []string{}
		number++
	}

	fmt.Println(res)
}

func move(paths [][]string, numAnts int){ 
	
	positions := make([]int, numAnts)
	started := 0
	finished := 0
	for i := range paths{
		paths[i] = paths[i][1:] 
	}

	for finished < numAnts {
		line := make([]string, numAnts)
		
		// Move ants that have already started
		for i := 0; i < started; i++ {
			pathIndex := i % len(paths)
			if positions[i] < len(paths[pathIndex]) {
				currentPos := paths[pathIndex][positions[i]]
					line[i] = fmt.Sprintf("L%d-%s", i+1, currentPos)
				
				positions[i]++
				if positions[i] == len(paths[pathIndex]) {
					finished++
				}
			}
		}

		// Start a new ant if possible
		if started < numAnts && (started == 0 || positions[started-1] > 1) {
			pathIndex := started % len(paths)
			started++
			positions[started-1] = 1
			line[started-1] = fmt.Sprintf("L%d-%s", started, paths[pathIndex][0])
		}

		output := strings.Join(filterEmptyStrings(line), " ")
		if output != "" {
			fmt.Println(output)
		}
	}
}

func filterEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}