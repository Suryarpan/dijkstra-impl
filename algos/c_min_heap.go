package algos

import "math"

// minHeap is a type-specific binary heap to avoid interface{} overhead
type minHeap struct {
	nodes []Edge
}

func (h *minHeap) Push(nd Edge) {
	h.nodes = append(h.nodes, nd)
	h.up(len(h.nodes) - 1)
}

func (h *minHeap) Pop() Edge {
	n := len(h.nodes) - 1
	res := h.nodes[0]
	h.nodes[0] = h.nodes[n]
	h.nodes = h.nodes[:n]
	h.down(0)
	return res
}

func (h *minHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.nodes[i].weight >= h.nodes[parent].weight {
			break
		}
		h.nodes[i], h.nodes[parent] = h.nodes[parent], h.nodes[i]
		i = parent
	}
}

func (h *minHeap) down(i int) {
	n := len(h.nodes)
	for {
		left := 2*i + 1
		if left >= n || left < 0 {
			break
		}
		smallest := left
		if right := left + 1; right < n && h.nodes[right].weight < h.nodes[left].weight {
			smallest = right
		}
		if h.nodes[i].weight <= h.nodes[smallest].weight {
			break
		}
		h.nodes[i], h.nodes[smallest] = h.nodes[smallest], h.nodes[i]
		i = smallest
	}
}

// DijkstraSSSP calculates the shortest path from a source using the custom heap
func DijkstraSSSP(g *Graph, source int) []int {
	dist := make([]int, len(g.edges))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	hp := &minHeap{nodes: make([]Edge, 0, len(g.edges))}
	hp.Push(Edge{to: source, weight: 0})

	for len(hp.nodes) > 0 {
		curr := hp.Pop()
		u, d := curr.to, curr.weight

		// Early skip for stale priority entries
		if d > dist[u] {
			continue
		}

		for _, edge := range g.edges[u] {
			if newDist := dist[u] + edge.weight; newDist < dist[edge.to] {
				dist[edge.to] = newDist
				hp.Push(Edge{to: edge.to, weight: newDist})
			}
		}
	}
	return dist
}
