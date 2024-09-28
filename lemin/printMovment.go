package lemin

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Ant struct {
	ID        int
	PathIndex int
	Position  int
}

func PrintAntMovements(paths [][]string, numAnts int, antsPerPath []int, start, end string) {
	ants := make([]*Ant, numAnts)
	antID := 1
	for pathIndex, numAntsInPath := range antsPerPath {
		for i := 0; i < numAntsInPath; i++ {
			ants[antID-1] = &Ant{ID: antID, PathIndex: pathIndex, Position: 0}
			antID++
		}
	}

	for !allAntsFinished(paths, ants, end) {
		moves := []string{}
		occupied := make(map[string]bool)

		for _, ant := range ants {
			if ant.Position == len(paths[ant.PathIndex])-1 {
				// Ant has reached to end
				continue 
			}

			nextRoom := paths[ant.PathIndex][ant.Position+1]
			if nextRoom != end && occupied[nextRoom] {
				// Skip if next room is already occupied
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
		if paths[ant.PathIndex][ant.Position] != end {
			return false
		}
	}
	return true
}

// Result prints the final movement result of the ants.
// It distributes the ants across the available paths and tracks their movements round by round.
func Result(paths [][]string, antCounts []int, ants int, end string) {
	// pathMoves stores the paths each ant takes
	pathMoves := make([][]string, len(paths)) // [[L1 L3 L5][L2 L4]]
	antCountsCopy := make([]int, len(antCounts))
	copy(antCountsCopy, antCounts)
	index := 0

	// If all paths have the same length, distribute ants evenly
	if isSameLen(paths) {
		for i := 1; i <= ants; i++ {
			if index == len(pathMoves) {
				index = 0
			}
			pathMoves[index] = append(pathMoves[index], fmt.Sprintf("L%d", i))
			index++
		}
	} else {
		// Distribute ants based on path lengths
		for i := 1; i <= ants; i++ {
			pathMoves[index] = append(pathMoves[index], fmt.Sprintf("L%d", i))
			if index == len(pathMoves)-1 {
				index = 0
			} else if antCountsCopy[index] <= antCountsCopy[index+1] {
				index++
				if antCountsCopy[index] == 0 {
					index = 0
				}
			}
			antCountsCopy[index]--
		}
	}

	// moveMap stores the movements of the ants
	moveMap := make(map[string][]string) // L1:[L1-start L1-A L1-end] L2:[L2-start L2-B L2-end] ...
	for _, moves := range pathMoves {
		for _, ant := range moves {
			moveMap[ant] = []string{}
		}
	}
	for i, path := range paths {
		for j := 0; j < len(pathMoves[i]); j++ { // from L1 ... Lx
			for _, move := range path[1:] { // [1:] to skip start
				moveMap[pathMoves[i][j]] = append(moveMap[pathMoves[i][j]], pathMoves[i][j]+"-"+move)
			}
		}
	}

	resultToPrint := make([][]string, len(paths[0])-2+antCounts[0]) //create array with rounds len
	for _, moves := range pathMoves {
		id := 0
		for _, ant := range moves {
			resultToPrint[id] = append(resultToPrint[id], moveMap[ant][0])
			for i := 0; i < len(moveMap[ant])-1; i++ {
				resultToPrint[id+i+1] = append(resultToPrint[id+i+1], moveMap[ant][i+1])
			}
			id++ // to add in next line until the end room
		}
	}
	for i := range resultToPrint {
		// Sort moves based on the numeric part of the ant name
		sort.Slice(resultToPrint[i], func(a, b int) bool {
			numA, _ := strconv.Atoi(strings.Split(resultToPrint[i][a], "-")[0][1:])
			numB, _ := strconv.Atoi(strings.Split(resultToPrint[i][b], "-")[0][1:])
			return numA < numB
		})
	}
	fmt.Println()
	for _, round := range resultToPrint {
		fmt.Println(strings.Join(round, " "))
	}
	fmt.Println()
}

// func isSameLen(paths [][]string) bool {
// 	firstLen := len(paths[0])
// 	for _, v := range paths {
// 		if firstLen != len(v) {
// 			return false
// 		}
// 	}
// 	return true
// }
////////////////////////////
// Result prints the final movement result of the ants.
// It distributes the ants across the available paths and tracks their movements round by round.
func Result2(paths [][]string, antCounts []int, ants int, end string) {
	// pathMoves stores which ant goes on which path
	pathMoves := make([][]string, len(paths))
	antCountsCopy := make([]int, len(antCounts))
	copy(antCountsCopy, antCounts)
	index := 0

	// Distribute ants across paths
	for i := 1; i <= ants; i++ {
		pathMoves[index] = append(pathMoves[index], fmt.Sprintf("L%d", i))
		
		// Update index for next ant
		if isSameLen(paths) || index == len(pathMoves)-1 {
			index = (index + 1) % len(pathMoves)
		} else if antCountsCopy[index] <= antCountsCopy[index+1] {
			index++
		}
		antCountsCopy[index]--
	}

	// Build movement map for each ant
	moveMap := make(map[string][]string)
	for i, moves := range pathMoves {
		for _, ant := range moves {
			for _, move := range paths[i][1:] { // [1:] to skip the start
				moveMap[ant] = append(moveMap[ant], fmt.Sprintf("%s-%s", ant, move))
			}
		}
	}

	// Prepare resultToPrint array for the number of rounds
	rounds := len(paths[0]) - 2 + antCounts[0]
	resultToPrint := make([][]string, rounds)

	// Append movements into resultToPrint
	for _, moves := range pathMoves {
		for id, ant := range moves {
			for i, move := range moveMap[ant] {
				resultToPrint[id+i] = append(resultToPrint[id+i], move)
			}
		}
	}

	// Sort moves by the numeric part of ant name in each round
	for i := range resultToPrint {
		sort.Slice(resultToPrint[i], func(a, b int) bool {
			numA, _ := strconv.Atoi(strings.Split(resultToPrint[i][a], "-")[0][1:])
			numB, _ := strconv.Atoi(strings.Split(resultToPrint[i][b], "-")[0][1:])
			return numA < numB
		})
	}

	// Print the result
	fmt.Println()
	for _, round := range resultToPrint {
		fmt.Println(strings.Join(round, " "))
	}
	fmt.Println()
}

// Check if all paths have the same length
func isSameLen(paths [][]string) bool {
	firstLen := len(paths[0])
	for _, v := range paths {
		if len(v) != firstLen {
			return false
		}
	}
	return true
}

