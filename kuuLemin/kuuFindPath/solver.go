package kuufindpath

import (
	"fmt"
	kuutype "lemin/kuuLemin/kuuType"
)

func calculateTurns(paths [][]string, ants int) int {
	lengths := make([]int, len(paths))

	for i, p := range paths {
		lengths[i] = len(p) - 1
	}

	load := make([]int, len(paths))
	maxTurns := 0

	for ant := 0; ant < ants; ant++ {
		best := 0

		for i := 1; i < len(paths); i++ {
			if load[i]+lengths[i] < load[best]+lengths[best] {
				best = i
			}
		}

		if finish := load[best] + lengths[best]; finish > maxTurns {
			maxTurns = finish
		}

		load[best]++
	}

	return maxTurns
}

func FindPaths(farm *kuutype.FarmAnts) ([][]string, error) {
	network := CreateResidualGraph(farm)

	originalCapacity := make(map[[2]string]int)
	for edge, cap := range network.Capacity {
		originalCapacity[edge] = cap
	}

	start := nodeIn(farm.StartRoom)
	end := nodeOut(farm.EndRoom)

	runEdmondsKarp(network.Adjacency, network.Capacity, start, end)

	allPaths := extractAllPaths(network.Adjacency, network.Capacity, originalCapacity, start, end)

	if len(allPaths) == 0 {
		return nil, fmt.Errorf("invalid data format: no path between start and end")
	}

	bestTurns := -1
	var bestPaths [][]string

	for i := 1; i <= len(allPaths); i++ {
		turns := calculateTurns(allPaths[:i], farm.NumAnts)

		if bestTurns == -1 || turns < bestTurns {
			bestTurns = turns
			bestPaths = make([][]string, i)
			copy(bestPaths, allPaths[:i])
		} else {
			break
		}
	}

	return bestPaths, nil
}
