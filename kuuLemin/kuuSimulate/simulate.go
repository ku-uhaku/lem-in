package kuuSimulate

import (
	"fmt"
	"strings"
)

func Simulate(paths [][]string, numAnts int) []string {

	pathLen := getPathLengths(paths)

	antPath, antStart, totalTurns := assignAnts(paths, pathLen, numAnts)

	return buildOutput(paths, pathLen, antPath, antStart, totalTurns)
}

func getPathLengths(paths [][]string) []int {
	lengths := make([]int, len(paths))

	for i := range paths {
		lengths[i] = len(paths[i]) - 1
	}

	return lengths
}

func assignAnts(paths [][]string, pathLen []int, numAnts int) ([]int, []int, int) {

	load := make([]int, len(paths))
	antPath := make([]int, numAnts)
	antStart := make([]int, numAnts)

	totalTurns := 0

	for ant := 0; ant < numAnts; ant++ {

		best := 0
		for i := 1; i < len(paths); i++ {
			if load[i]+pathLen[i] < load[best]+pathLen[best] {
				best = i
			}
		}

		antPath[ant] = best
		antStart[ant] = load[best]

		finish := antStart[ant] + pathLen[best]
		if finish > totalTurns {
			totalTurns = finish
		}

		load[best]++
	}

	return antPath, antStart, totalTurns
}

func buildOutput(paths [][]string, pathLen []int, antPath []int, antStart []int, totalTurns int) []string {

	var output []string

	for turn := 0; turn < totalTurns; turn++ {

		var moves []string

		for ant := range antPath {

			p := antPath[ant]

			if turn < antStart[ant] || turn >= antStart[ant]+pathLen[p] {
				continue
			}

			step := turn - antStart[ant] + 1
			room := paths[p][step]

			moves = append(moves, fmt.Sprintf("L%d-%s", ant+1, room))
		}

		if len(moves) > 0 {
			output = append(output, strings.Join(moves, " "))
		}
	}

	return output
}
