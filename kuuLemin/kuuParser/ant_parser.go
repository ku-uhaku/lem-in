package kuuParser

import (
	"fmt"
	"strconv"
)

// parseAnts parses the first line containing number of ants
func (p *Parser) parseAnts(line string) error {
	ants, err := strconv.Atoi(line)
	if err != nil {
		return fmt.Errorf("line %d: invalid number of ants: %v", p.lineNum, err)
	}
	if ants <= 0 {
		return fmt.Errorf("line %d: number of ants must be positive, got: %d", p.lineNum, ants)
	}
	p.farm.NumAnts = ants
	return nil
}