package kuutype

type FlowNetwork struct {
	Adjacency map[string][]string
	Capacity  map[[2]string]int
}

func NewFlowNetwork() *FlowNetwork {
	return &FlowNetwork{
		Adjacency: make(map[string][]string),
		Capacity:  make(map[[2]string]int),
	}
}

func (fn *FlowNetwork) AddEdge(from, to string, capacity int) {
	key := [2]string{from, to}

	fn.Capacity[key] += capacity
	fn.Adjacency[from] = append(fn.Adjacency[from], to)
	fn.Adjacency[to] = append(fn.Adjacency[to], from)
}