package kuuParser

import (
	"fmt"
)

// parseComment handles ##start, ##end, and regular comments
func (p *Parser) parseComment(line string) error {
	switch line {
	case "##start":
		if p.startSet {
			return fmt.Errorf("line %d: multiple ##start commands", p.lineNum)
		}
		p.nextIsStart = true

	case "##end":
		if p.endSet {
			return fmt.Errorf("line %d: multiple ##end commands", p.lineNum)
		}
		p.nextIsEnd = true
	}
	// Regular comments are ignored
	return nil
}
