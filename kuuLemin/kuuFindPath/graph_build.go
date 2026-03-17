package kuufindpath

import kuutype "lemin/kuuLemin/kuuType"

func nodeIn(room string) string  { return room + "|in" }
func nodeOut(room string) string { return room + "|out" }

func splitCapacityRoom(room string, ants int, start, end string) int {
	if room == start || room == end {
		return ants
	}
	return 1
}

func CreateResidualGraph(farm *kuutype.FarmAnts) *kuutype.FlowNetwork {
	network := kuutype.NewFlowNetwork()

	for _, room := range farm.Rooms {
		capacity := splitCapacityRoom(room, farm.NumAnts, farm.StartRoom, farm.EndRoom)

		in := nodeIn(room)
		out := nodeOut(room)

		network.AddEdge(in, out, capacity)
	}

	for room, neighbors := range farm.Links {
		for _, neighbor := range neighbors {
			network.AddEdge(nodeOut(room), nodeIn(neighbor), 1)
		}
	}

	return network
}
