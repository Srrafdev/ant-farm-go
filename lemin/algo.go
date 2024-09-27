package lemin

import (
	"fmt"
)

func Chouse(paths [][]string, ants int) ([][]string, []int) {
	pat := beastPaths(paths)
	beastRound, _, NAP := calcRoundsAndMoves(pat, ants)
	beastMove := 0

	beastPath := pat
	numAntpat := NAP
	for i := 0; i < len(paths); i++ {
		pat := beastPaths(paths[i:])
		round, moves, nap := calcRoundsAndMoves(pat, ants)
		if round < beastRound {
			beastRound = round
			// get best paths
			numAntpat = nap
			beastPath = pat
		} else if round == beastRound {
			if moves < beastMove || beastMove == 0 {
				beastMove = moves
				// get best paths
				beastPath = pat
			}
		}
		// if len(pat)<= flow{
		// 	break
		// }

	}
	fmt.Printf("\n\n(best path):%v  (best round):%v  (moves):%v\n\n", beastPath, beastRound, beastMove)
	return beastPath, numAntpat
}

func calcRoundsAndMoves(paths [][]string, ants int) (int, int, []int) {

	numbAntForPath := distributeDivision(paths, ants)
	fmt.Println("         ", numbAntForPath)

	rounds := (len(paths[0]) - 2) + numbAntForPath[0]
	moves := calcMoves(paths, numbAntForPath)

	fmt.Printf("\n(path):%v  (round):%v  (moves):%v", paths, rounds, moves)
	return rounds, moves, numbAntForPath
}

func calcMoves(path [][]string, divAnts []int) int {
	res := 0
	for i, v := range path {
		res += (len(v) - 1) * divAnts[i]
	}
	return res
}

// is not corect
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

// sort parhs
func SortPaths(path [][]string) [][]string {
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
