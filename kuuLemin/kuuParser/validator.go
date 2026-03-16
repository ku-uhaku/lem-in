package kuuParser

import (
	"fmt"
)

// validateFinal performs final validation after parsing all lines
func (p *Parser) validateFinal() error {
	// Check if start room exists
	if p.farm.StartRoom == "" {
		return fmt.Errorf("no start room defined (##start)")
	}

	// Check if end room exists
	if p.farm.EndRoom == "" {
		return fmt.Errorf("no end room defined (##end)")
	}

	// Validate start room exists in rooms
	if !p.rooms[p.farm.StartRoom] {
		return fmt.Errorf("start room '%s' not found in rooms", p.farm.StartRoom)
	}

	// Validate end room exists in rooms
	if !p.rooms[p.farm.EndRoom] {
		return fmt.Errorf("end room '%s' not found in rooms", p.farm.EndRoom)
	}

	return nil
}
