package kuuParser

import (
	"fmt"
	"strconv"
	"strings"
)

// Room represents a parsed room
type Room struct {
	Name string
	X    int
	Y    int
}

// parseRoom parses a room line (name x y)
func (p *Parser) parseRoom(line string) error {
	room, err := p.validateRoomFormat(line)
	if err != nil {
		return err
	}

	if err := p.validateRoomUniqueness(room); err != nil {
		return err
	}

	// Add room to farm
	p.farm.Rooms = append(p.farm.Rooms, room.Name)
	p.rooms[room.Name] = true
	p.coordinates[fmt.Sprintf("%d,%d", room.X, room.Y)] = true

	// Handle start/end room markers
	return p.handleRoomMarkers(room.Name)
}

// validateRoomFormat checks if the room line has valid format
func (p *Parser) validateRoomFormat(line string) (*Room, error) {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		return nil, fmt.Errorf("line %d: invalid room format: %s", p.lineNum, line)
	}

	name := parts[0]
	x, err1 := strconv.Atoi(parts[1])
	y, err2 := strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("line %d: invalid coordinates for room %s", p.lineNum, name)
	}

	return &Room{Name: name, X: x, Y: y}, nil
}

// validateRoomUniqueness checks for duplicate names and coordinates
func (p *Parser) validateRoomUniqueness(room *Room) error {
	if p.rooms[room.Name] {
		return fmt.Errorf("line %d: duplicate room name: %s", p.lineNum, room.Name)
	}

	coordKey := fmt.Sprintf("%d,%d", room.X, room.Y)
	if p.coordinates[coordKey] {
		return fmt.Errorf("line %d: duplicate coordinates (%d,%d) for room %s",
			p.lineNum, room.X, room.Y, room.Name)
	}

	return nil
}

// handleRoomMarkers sets start and end rooms if marked
func (p *Parser) handleRoomMarkers(roomName string) error {
	if p.nextIsStart {
		if p.farm.StartRoom != "" {
			return fmt.Errorf("line %d: multiple start rooms defined", p.lineNum)
		}
		p.farm.StartRoom = roomName
		p.startSet = true
		p.nextIsStart = false
	}

	if p.nextIsEnd {
		if p.farm.EndRoom != "" {
			return fmt.Errorf("line %d: multiple end rooms defined", p.lineNum)
		}
		p.farm.EndRoom = roomName
		p.endSet = true
		p.nextIsEnd = false
	}

	return nil
}
