package kuuParser

import (
	
	"strings"
)

// parseLine dispatches to the appropriate parser based on line type
func (p *Parser) parseLine(line string) error {
	// Handle comments and commands
	if strings.HasPrefix(line, "#") {
		return p.parseComment(line)
	}

	// Parse ants (first non-comment line)
	if p.farm.NumAnts == 0 {
		return p.parseAnts(line)
	}

	// Parse link or room
	if strings.Contains(line, "-") {
		return p.parseLink(line)
	}
	return p.parseRoom(line)
}