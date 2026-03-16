package kuulemin

import (
	"fmt"
	kuufindpath "lemin/kuuLemin/kuuFindPath"
	kuuparser "lemin/kuuLemin/kuuParser"
	"lemin/kuuLemin/kuuSimulate"
	"os"
)

func Lemin() {
	lines, err := kuuparser.ReadFile()
	if err != nil {
		fmt.Printf("Read Error: %v \n", err)
		os.Exit(1)
	}
	farm, err := kuuparser.ParseFromLines(lines)
	if err != nil {
		fmt.Printf("Parse Error: %v \n", err)
		os.Exit(1)
	}

	paths, err := kuufindpath.FindPaths(farm)
	if err != nil {
		fmt.Println(err)
		return
	}

	moves := kuuSimulate.Simulate(paths, farm.NumAnts)

	// All good — print everything
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
	for _, move := range moves {
		fmt.Println(move)
	}

}
