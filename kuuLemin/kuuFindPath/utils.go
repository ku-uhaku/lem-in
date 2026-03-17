package kuufindpath

import "strings"

func isRoomNode(node string) bool {
	return strings.HasSuffix(node, "|in")
}

func extractRoom(node string) string {
	return strings.TrimSuffix(node, "|in")
}

func convertNodesToRooms(path []string) []string {
	var rooms []string

	for _, node := range path {
		if isRoomNode(node) {
			rooms = append(rooms, extractRoom(node))
		}
	}

	return rooms
}