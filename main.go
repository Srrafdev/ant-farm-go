package main

import (
	"fmt"
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

	paths = sortPaths(paths)
	for _, v := range paths {
		fmt.Println(v)
	}
	println("****************************************")
	// pp := beastPaths(paths[2:])

	// roud, moves := calcRoundsAndMoves(pp, farms.NumberAnts)
	// fmt.Println(roud,"\n",moves)
	pp := chouse(paths, farms.NumberAnts)
	fmt.Println(pp)
	println("****************************************")
	//walkAnts(pp, farms.Start, farms.NumberAnts)
	move(paths,farms.NumberAnts)
}

func chouse(paths [][]string, ants int) [][]string {
	pat := beastPaths(paths[0:])
	beastRound, beastMove := calcRoundsAndMoves(pat, ants)

	beastPath := pat
	for i := 0; i < len(paths)-1; i++ {
		pat := beastPaths(paths[i:])
		round, moves := calcRoundsAndMoves(pat, ants)
		if round < beastRound {
			beastRound = round
			// get beast paths
			beastPath = pat
		} else if round == beastRound {
			if moves < beastMove {
				beastMove = moves
				// get beast paths
				beastPath = pat
			}
		}
	}
	fmt.Println(beastRound, beastMove)
	return beastPath
}

func calcRoundsAndMoves(paths [][]string, ants int) (int, int) {
	if len(paths) == 1 {
		return len(paths[0]) + ants, len(paths[0]) + ants
	}
	pathsLen := make([]int, len(paths))
	for i, path := range paths {
		pathsLen[i] = len(path) - 1
	}

	numbAntForPath := make([]int, len(pathsLen))

	remainingAnts := ants
	for remainingAnts > 0 {
		for i := range pathsLen {
			if remainingAnts == 0 {
				break
			}
			numbAntForPath[i]++
			remainingAnts--
		}
	}

	rounds := 0
	moves := 0

	for i, pathLen := range pathsLen {
		if numbAntForPath[i] == 0 {
			continue
		}
		roundsForPath := pathLen + (numbAntForPath[i] - 1)
		// rounds = int(math.Max(float64(rounds), float64(roundsForPath)))
		if roundsForPath > rounds {
			rounds = roundsForPath
		}
		moves += pathLen * numbAntForPath[i]
	}

	return rounds, moves
}

func beastPaths(paths [][]string) [][]string {
	var filter [][]string
	for _, path := range paths {
		if isDeferentRom(path, filter) {
			filter = append(filter, path)
		}
	}
	return filter
}

func isDeferentRom(paths []string, filter [][]string) bool {
	for _, p := range filter {
		for i := 1; i < len(p)-1; i++ {
			for _, path := range paths {
				if p[i] == path {
					return false
				}
			}
		}
	}
	return true
}

type Ant struct {
	id      int
	room    string
	canMove bool
}

func newAnt(id int, room string, canmove bool) *Ant {
	return &Ant{
		id:      id,
		room:    room,
		canMove: canmove,
	}
}

func walkAnts(paths [][]string, start string, numbAnt int) {
	var ants []Ant
	for i := 1; i <= numbAnt; i++ {
		ants = append(ants, *newAnt(i, start, true))
	}

	for _, path := range paths {
		for j := range ants {
			for i := 1; i < len(path); i++ {
				if ants[j].canMove {
					ants[j].room = path[i]
					ants[j].canMove = false
				}
				fmt.Printf("L%v-%v ", ants[j].id, ants[j].room)
			}
		}
		println()
	}
}

func move(paths [][]string, numAnts int) {
	positions := make([]int, numAnts)
	started := 0
	finished := 0
	for i := range paths {
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
// func antherAlgo(paths [][]string, numbAnts int) {
// 	n := 1
// 	way := make(map[int][]string)
// 	for n <= numbAnts {
// 		for i, path := range paths {
// 			if i+1 < len(paths) && nambTips(path, n) > nambTips(paths[i+1], 1) {
// 				delete(way, n)
// 			} else {
// 				way[n] = path
// 			}
// 		}
// 		n++
// 	}
// 	fmt.Println(way)
// }

// func nambTips(path []string, ants int) int {
// 	return (len(path) - 1) + ants
// }

// func choosePaths(paths [][]string) [][]string {
// 	rating := rate(paths)

// 	SortByRate(paths, rating)

// 	paths = choose(paths)
// 	return paths
// }

// func choose(paths [][]string) [][]string {
// 	filter := [][]string{}
// 	m := make(map[string]bool)
// 	for i, path := range paths {

// 		for j := 1; j < len(path)-1; j++ {
// 			if m[paths[i][j]] {
// 				goto next // go to next index
// 			}
// 		}

// 		for j := 1; j < len(paths[i])-1; j++ {
// 			m[paths[i][j]] = true
// 		}
// 		filter = append(filter, paths[i])

// 	next:
// 	}
// 	return filter
// }

// func SortByRate(paths [][]string, rating []int) {
// 	for i := 0; i < len(paths); i++ {
// 		for j := i + 1; j < len(paths); j++ {
// 			if rating[i] > rating[j] {
// 				rating[i], rating[j] = rating[j], rating[i]
// 				paths[i], paths[j] = paths[j], paths[i]
// 			}
// 		}
// 	}
// }

// func rate(paths [][]string) []int {
// 	repetRom := make(map[string]int)
// 	for _, path := range paths {
// 		for _, rom := range path {
// 			repetRom[rom]++
// 		}
// 	}

// 	rating := make([]int, len(paths))
// 	for i, path := range paths {
// 		for _, rom := range path {
// 			rating[i] += repetRom[rom]
// 		}
// 	}

// 	return rating
// }

// func distributeDivision(paths [][]string, total int) []int {
// 	n := len(paths)
// 	maxLen := 0

// 	// Find the maximum path length
// 	for _, path := range paths {
// 		if n > maxLen {
// 			maxLen = len(path)
// 		}
// 	}

// 	result := make([]int, n)
// 	remainingTotal := total

// 	// Calculate the initial division for each path
// 	for i, path := range paths {
// 		difference := maxLen - len(path)
// 		result[i] = difference
// 		remainingTotal -= difference
// 	}

// 	// Distribute any remaining total
// 	for i := 0; i < remainingTotal; i++ {
// 		result[i%n]++
// 	}

// 	return result
// }

// func walk(path []string, n, all int) {
// 	var walkAnt []string
// 	var res [][]string
// 	number := all
// 	for number <= n+all-1 {
// 		for _, room := range path[1:] {
// 			walkAnt = append(walkAnt, "L"+strconv.Itoa(number)+"-"+room+" ")
// 		}
// 		res = append(res, walkAnt)
// 		walkAnt = []string{}
// 		number++
// 	}

// 	fmt.Println(res)
// }

// func move(paths [][]string, numAnts int) {
// 	positions := make([]int, numAnts)
// 	started := 0
// 	finished := 0
// 	for i := range paths {
// 		paths[i] = paths[i][1:]
// 	}

// 	for finished < numAnts {
// 		line := make([]string, numAnts)

// 		// Move ants that have already started
// 		for i := 0; i < started; i++ {
// 			pathIndex := i % len(paths)
// 			if positions[i] < len(paths[pathIndex]) {
// 				currentPos := paths[pathIndex][positions[i]]
// 				line[i] = fmt.Sprintf("L%d-%s", i+1, currentPos)

// 				positions[i]++
// 				if positions[i] == len(paths[pathIndex]) {
// 					finished++
// 				}
// 			}
// 		}

// 		// Start a new ant if possible
// 		if started < numAnts && (started == 0 || positions[started-1] > 1) {
// 			pathIndex := started % len(paths)
// 			started++
// 			positions[started-1] = 1
// 			line[started-1] = fmt.Sprintf("L%d-%s", started, paths[pathIndex][0])
// 		}

// 		output := strings.Join(filterEmptyStrings(line), " ")
// 		if output != "" {
// 			fmt.Println(output)
// 		}
// 	}
// }

func filterEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
