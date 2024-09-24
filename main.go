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

	pp := chouse(paths, farms.NumberAnts)
	fmt.Println(pp)
	println("****************************************")
	printAntMovements(pp, farms.NumberAnts, farms.Start, farms.End)
}

func chouse(paths [][]string, ants int) [][]string {
	pat := beastPaths(paths[0:])
	beastRound, _ := calcRoundsAndMoves(pat, ants)
	beastMove := 0

	beastPath := pat
	for i := 0; i < len(paths)-1; i++ {
		pat := beastPaths(paths[i:])
		round, moves := calcRoundsAndMoves(pat, ants)
		if round < beastRound {
			beastRound = round
			// get beast paths
			beastPath = pat
		} else if round == beastRound {
			if moves < beastMove || beastMove == 0 {
				beastMove = moves
				// get beast paths
				beastPath = pat
			}
		}
		// fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$",pat,round , moves)
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
			// if i+1 < len(pathsLen) && numbAntForPath[i] + i < pathsLen[i+1]{
			numbAntForPath[i]++
			remainingAnts--
			//}

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
	fmt.Println(paths)
	fmt.Println("***", rounds, moves)
	fmt.Println("***", numbAntForPath)

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
	ID        int
	PathIndex int
	Position  int
}

func printAntMovements(paths [][]string, numAnts int, start, end string) {
	ants := make([]*Ant, numAnts)
	for i := 0; i < numAnts; i++ {
		ants[i] = &Ant{ID: i + 1, PathIndex: i % len(paths), Position: -1}
	}

	for !allAntsFinished(paths, ants, end) {
		moves := []string{}
		occupied := make(map[string]bool)

		for _, ant := range ants {
			if ant.Position == len(paths[ant.PathIndex])-1 {
				continue
			}

			nextRoom := paths[ant.PathIndex][ant.Position+1]
			if nextRoom != end && occupied[nextRoom] {
				continue
			}

			ant.Position++
			currentRoom := paths[ant.PathIndex][ant.Position]
			if currentRoom != start && currentRoom != end {
				occupied[currentRoom] = true
			}
			if currentRoom != start {
				moves = append(moves, fmt.Sprintf("L%d-%s", ant.ID, currentRoom))
			}
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}

	}
}

func allAntsFinished(paths [][]string, ants []*Ant, end string) bool {
	for _, ant := range ants {
		if ant.Position == -1 || paths[ant.PathIndex][ant.Position] != end {
			return false
		}
	}
	return true
}
