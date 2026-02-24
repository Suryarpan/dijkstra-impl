package algos

// Edge representation
type Edge struct {
	to, weight int
}

type Graph struct {
	edges [][]Edge
}

func NewGraph(sz uint) *Graph {
	return &Graph{
		edges: make([][]Edge, sz),
	}
}

func (g *Graph) AddEdge(from, to, weight int) *Graph {
	g.edges[from] = append(g.edges[from], Edge{to: to, weight: weight})
	return g
}
