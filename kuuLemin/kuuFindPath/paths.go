package kuufindpath

func findFlowPath(graph map[string][]string, cap map[[2]string]int, original map[[2]string]int, start, end string) (map[string]string, bool) {
	parent := map[string]string{start: ""}
	queue := []string{start}

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

				if next == end {
					return parent, true
				}

				queue = append(queue, next)
			}
		}
	}

	return nil, false
}

func reconstructPath(parent map[string]string, end string) []string {
	var path []string

	for node := end; node != ""; node = parent[node] {
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

func extractAllPaths(graph map[string][]string, cap map[[2]string]int, original map[[2]string]int, start, end string) [][]string {
	var paths [][]string

	for {
		parent, found := findFlowPath(graph, cap, original, start, end)
		if !found {
			break
		}

		rawPath := reconstructPath(parent, end)
		removePathFlow(rawPath, cap)

		roomPath := convertNodesToRooms(rawPath)
		paths = append(paths, roomPath)
	}

	return paths
}