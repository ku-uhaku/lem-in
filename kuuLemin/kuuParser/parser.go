package kuuParser

import (
	"fmt"
	"strings"

	kuutype "lemin/kuuLemin/kuuType"
)

// Parser holds the state during parsing
type Parser struct {
	farm    *kuutype.FarmAnts
	lines   []string
	lineNum int

	// Tracking maps for validation
	rooms       map[string]bool
	coordinates map[string]bool
	links       map[string]bool

	// State flags
	nextIsStart bool
	nextIsEnd   bool
	startSet    bool
	endSet      bool
}

// NewParser creates a new parser instance
func NewParser(lines []string) *Parser {
	return &Parser{
		farm: &kuutype.FarmAnts{
			Rooms: []string{},
			Links: make(map[string][]string),
		},
		lines:       lines,
		rooms:       make(map[string]bool),
		coordinates: make(map[string]bool),
		links:       make(map[string]bool),
	}
}

// Parse performs the complete parsing
func (p *Parser) Parse() (*kuutype.FarmAnts, error) {
	if len(p.lines) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	for i, line := range p.lines {
		p.lineNum = i + 1
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if err := p.parseLine(line); err != nil {
			return nil, err
		}
	}

	// Final validation
	if err := p.validateFinal(); err != nil {
		return nil, err
	}

	return p.farm, nil
}

// ParseFromLines parses directly from a string slice (useful for testing)
func ParseFromLines(lines []string) (*kuutype.FarmAnts, error) {
	parser := NewParser(lines)
	return parser.Parse()
}
