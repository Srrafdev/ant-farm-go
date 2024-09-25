package lemin

import "fmt"

func Chouse(paths [][]string, ants int) [][]string {
	pat := beastPaths(paths)
	beastRound, _ := calcRoundsAndMoves(pat, ants)
	beastMove := 0

	beastPath := pat
	for i := 0; i < len(paths); i++ {
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
	}
	fmt.Printf("\n\n(beast path):%v  (beast round):%v  (moves):%v\n\n",beastPath,beastRound,beastMove,)
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
			 if i+1 < len(pathsLen) && numbAntForPath[i] + i < pathsLen[i+1]{
			numbAntForPath[i]++
			remainingAnts--
			}else if i+1 < len(pathsLen){
				numbAntForPath[i+1]++
				remainingAnts--
			}

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
		// calcul number of maves
		//moves += pathLen * numbAntForPath[i]
	}
	fmt.Printf("\n(path):%v  (round):%v  (moves):%v",paths,rounds,moves)
	moves = calcMoves(paths,numbAntForPath)
	return rounds, moves
}

func calcMoves(path [][]string, divAnts []int) int {
    res := 0
    for i, v := range path {
        res += (len(v) - 1) * divAnts[i]
    }
    return res
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

