package lemin

import (
	"fmt"
	"strings"
)

type Ant struct {
	ID        int
	PathIndex int
	Position  int
}

func PrintAntMovements(paths [][]string, numAnts int, start, end string) {
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
