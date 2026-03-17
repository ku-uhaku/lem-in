package kuufindpath

func bfs(graph map[string][]string, capacity map[[2]string]int, start, end string) map[string]string {
	parent := map[string]string{start: ""}
	queue := []string{start}

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

			if next == end {
				return parent
			}

			queue = append(queue, next)
		}
	}

	return nil
}

func runEdmondsKarp(graph map[string][]string, capacity map[[2]string]int, start, end string) {
	for {
		parent := bfs(graph, capacity, start, end)
		if parent == nil {
			return
		}

		flow := 1

		for node := end; node != start; node = parent[node] {
			prev := parent[node]

			capacity[[2]string{prev, node}] -= flow
			capacity[[2]string{node, prev}] += flow
		}
	}
}
	