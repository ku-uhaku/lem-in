package kuutype	

import (
	"fmt"
	"strings"
)

type FarmAnts struct {
	NumAnts   int
	Rooms     []string
	Links     map[string][]string
	StartRoom string
	EndRoom   string
}

// String returns a string representation of the FarmAnts
func (f *FarmAnts) String() string {
	var builder strings.Builder

	// Write number of ants
	builder.WriteString(fmt.Sprintf("Number of ants: %d\n", f.NumAnts))

	// Write rooms
	builder.WriteString("\nRooms:\n")
	for _, room := range f.Rooms {
		prefix := "  "
		if room == f.StartRoom {
			prefix = "  [START] "
		} else if room == f.EndRoom {
			prefix = "  [END]   "
		}
		builder.WriteString(fmt.Sprintf("%s%s\n", prefix, room))
	}

	// Write links
	builder.WriteString("\nLinks:\n")
	if len(f.Links) == 0 {
		builder.WriteString("  No links\n")
	} else {
		printed := make(map[string]bool)
		for room, connections := range f.Links {
			for _, conn := range connections {
				key1 := room + "-" + conn
				key2 := conn + "-" + room
				if !printed[key1] && !printed[key2] {
					builder.WriteString(fmt.Sprintf("  %s <-> %s\n", room, conn))
					printed[key1] = true
				}
			}
		}
	}

	return builder.String()
}
