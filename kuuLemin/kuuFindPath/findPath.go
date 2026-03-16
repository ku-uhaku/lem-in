package kuufindpath

import (
	"fmt"
	kuutype "lemin/kuuLemin/kuuType"
	"strings"
)

// FlowNetwork represents a residual flow graph
type FlowNetwork struct {
	Adjacency map[string][]string
	Capacity  map[[2]string]int
}

// NewFlowNetwork initializes a new network
func NewFlowNetwork() *FlowNetwork {
	return &FlowNetwork{
		Adjacency: make(map[string][]string),
		Capacity:  make(map[[2]string]int),
	}
}

// AddEdge adds a directed edge with capacity
func (fn *FlowNetwork) AddEdge(from, to string, capacity int) {
	key := [2]string{from, to}

	fn.Capacity[key] += capacity
	fn.Adjacency[from] = append(fn.Adjacency[from], to)
	fn.Adjacency[to] = append(fn.Adjacency[to], from) // reverse edge for residual graph
}

// Convert room -> split nodes
func nodeIn(room string) string  { return room + "|in" }
func nodeOut(room string) string { return room + "|out" }

func isRoomNode(node string) bool {
	return strings.HasSuffix(node, "|in")
}

func extractRoom(node string) string {
	return strings.TrimSuffix(node, "|in")
}

func splitRoom(room string, ants int, start, end string) int {
	if room == start || room == end {
		return ants
	}
	return 1
}

func CreateResidualGraph(farm *kuutype.FarmAnts) (*FlowNetwork, error) {
	network := NewFlowNetwork()

	// Split rooms into in/out nodes
	for _, room := range farm.Rooms {

		capacity := splitRoom(room, farm.NumAnts, farm.StartRoom, farm.EndRoom)

		in := nodeIn(room)
		out := nodeOut(room)

		network.AddEdge(in, out, capacity)
	}

	// Connect tunnels
	for room, neighbors := range farm.Links {
		for _, neighbor := range neighbors {

			from := nodeOut(room)
			to := nodeIn(neighbor)

			network.AddEdge(from, to, 1)
		}
	}

	return network, nil
}

func bfsFindPath(graph map[string][]string, capacity map[[2]string]int, source, sink string) map[string]string {

	parent := map[string]string{source: ""}
	queue := []string{source}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current] {

			if _, visited := parent[next]; visited {
				continue
			}

			if capacity[[2]string{current, next}] <= 0 {
				continue
			}

			parent[next] = current

			if next == sink {
				return parent
			}

			queue = append(queue, next)
		}
	}

	return nil
}

func runEdmondsKarp(graph map[string][]string, capacity map[[2]string]int, source, sink string) {

	for {

		parent := bfsFindPath(graph, capacity, source, sink)

		if parent == nil {
			return
		}

		flow := 1

		for node := sink; node != source; node = parent[node] {

			prev := parent[node]

			capacity[[2]string{prev, node}] -= flow
			capacity[[2]string{node, prev}] += flow
		}
	}
}

func findFlowPath(graph map[string][]string, cap map[[2]string]int, original map[[2]string]int, source, sink string) (map[string]string, bool) {

	parent := map[string]string{source: ""}
	queue := []string{source}

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		for _, next := range graph[current] {

			if _, visited := parent[next]; visited {
				continue
			}

			key := [2]string{current, next}

			if original[key] > cap[key] {

				parent[next] = current

				if next == sink {
					return parent, true
				}

				queue = append(queue, next)
			}
		}
	}

	return nil, false
}

func reconstructPath(parent map[string]string, source, sink string) []string {

	var path []string

	for node := sink; node != ""; node = parent[node] {
		path = append([]string{node}, path...)
	}

	return path
}

func removePathFlow(path []string, capacity map[[2]string]int) {

	for i := 0; i < len(path)-1; i++ {

		from := path[i]
		to := path[i+1]

		capacity[[2]string{from, to}]++
	}
}

func convertNodesToRooms(path []string) []string {

	var rooms []string

	for _, node := range path {

		if isRoomNode(node) {

			room := extractRoom(node)
			rooms = append(rooms, room)
		}
	}

	return rooms
}

func extractAllPaths(graph map[string][]string, cap map[[2]string]int, original map[[2]string]int, source, sink string) [][]string {

	var paths [][]string

	for {

		parent, found := findFlowPath(graph, cap, original, source, sink)

		if !found {
			break
		}

		rawPath := reconstructPath(parent, source, sink)

		removePathFlow(rawPath, cap)

		roomPath := convertNodesToRooms(rawPath)

		paths = append(paths, roomPath)
	}

	return paths
}

func calculateTurns(paths [][]string, ants int) int {

	lengths := make([]int, len(paths))
	for i, p := range paths {
		lengths[i] = len(p) - 1
	}

	load := make([]int, len(paths))
	maxTurns := 0

	for ant := 0; ant < ants; ant++ {

		best := 0

		for i := 1; i < len(paths); i++ {
			if load[i]+lengths[i] < load[best]+lengths[best] {
				best = i
			}
		}

		if finish := load[best] + lengths[best]; finish > maxTurns {
			maxTurns = finish
		}

		load[best]++
	}

	return maxTurns
}

func FindPaths(farm *kuutype.FarmAnts) ([][]string, error) {

	network, err := CreateResidualGraph(farm)
	if err != nil {
		return nil, err
	}

	originalCapacity := make(map[[2]string]int)

	for edge, cap := range network.Capacity {
		originalCapacity[edge] = cap
	}

	source := nodeIn(farm.StartRoom)
	sink := nodeOut(farm.EndRoom)

	runEdmondsKarp(network.Adjacency, network.Capacity, source, sink)

	allPaths := extractAllPaths(network.Adjacency, network.Capacity, originalCapacity, source, sink)

	if len(allPaths) == 0 {
		return nil, fmt.Errorf("invalid data format: no path between start and end")
	}

	bestTurns := -1
	var bestPaths [][]string

	for i := 1; i <= len(allPaths); i++ {

		turns := calculateTurns(allPaths[:i], farm.NumAnts)

		if bestTurns == -1 || turns < bestTurns {

			bestTurns = turns
			bestPaths = make([][]string, i)

			copy(bestPaths, allPaths[:i])

		} else {
			break
		}
	}

	return bestPaths, nil
}