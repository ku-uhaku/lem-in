package kuuParser

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile() ([]string, error) {
	args := os.Args

	if len(args) != 2 {
		return nil, fmt.Errorf("go run . <filename>")
	}
	filename := args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Cannot read file %s \n %v", filename, err)
	}
	lines := strings.Split(string(data), "\n")

	return lines, nil
}
