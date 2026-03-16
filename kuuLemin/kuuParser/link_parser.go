package kuuParser

import (
	"fmt"
	"strings"
)

// Link represents a connection between two rooms
type Link struct {
	Room1 string
	Room2 string
}

// parseLink parses a link line (room1-room2)
func (p *Parser) parseLink(line string) error {
	link, err := p.validateLinkFormat(line)
	if err != nil {
		return err
	}

	if err := p.validateLinkRooms(link); err != nil {
		return err
	}

	if err := p.validateLinkUniqueness(link); err != nil {
		return err
	}

	// Add bidirectional links
	p.farm.Links[link.Room1] = append(p.farm.Links[link.Room1], link.Room2)
	p.farm.Links[link.Room2] = append(p.farm.Links[link.Room2], link.Room1)

	// Mark link as used
	p.links[link.Room1+"-"+link.Room2] = true

	return nil
}

// validateLinkFormat checks if the link has valid format
func (p *Parser) validateLinkFormat(line string) (*Link, error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("line %d: invalid link format: %s", p.lineNum, line)
	}

	room1 := strings.TrimSpace(parts[0])
	room2 := strings.TrimSpace(parts[1])

	if room1 == "" || room2 == "" {
		return nil, fmt.Errorf("line %d: empty room name in link: %s", p.lineNum, line)
	}

	return &Link{Room1: room1, Room2: room2}, nil
}

// validateLinkRooms checks if both rooms exist
func (p *Parser) validateLinkRooms(link *Link) error {
	if !p.rooms[link.Room1] {
		return fmt.Errorf("line %d: room '%s' doesn't exist", p.lineNum, link.Room1)
	}
	if !p.rooms[link.Room2] {
		return fmt.Errorf("line %d: room '%s' doesn't exist", p.lineNum, link.Room2)
	}
	return nil
}

// validateLinkUniqueness checks for duplicate links
func (p *Parser) validateLinkUniqueness(link *Link) error {
	key1 := link.Room1 + "-" + link.Room2
	key2 := link.Room2 + "-" + link.Room1

	if p.links[key1] || p.links[key2] {
		return fmt.Errorf("line %d: duplicate link between %s and %s",
			p.lineNum, link.Room1, link.Room2)
	}

	return nil
}
