package algos

import "math"

type IndexedPQ struct {
	nodes []int // The heap containing node IDs
	dist  []int // Reference to the distances array
	pos   []int // pos[nodeID] = index in 'nodes' slice; -1 if not in heap
}

func NewIndexedPQ(size int, dists []int) *IndexedPQ {
	pos := make([]int, size)
	for i := range pos {
		pos[i] = -1
	}
	return &IndexedPQ{
		nodes: make([]int, 0, size),
		dist:  dists,
		pos:   pos,
	}
}

func (pq *IndexedPQ) PushOrUpdate(node int) {
	if pq.pos[node] == -1 {
		// Push new node
		pq.pos[node] = len(pq.nodes)
		pq.nodes = append(pq.nodes, node)
		pq.up(len(pq.nodes) - 1)
	} else {
		// Update existing node (it can only go "up" in Dijkstra)
		pq.up(pq.pos[node])
	}
}

func (pq *IndexedPQ) Pop() int {
	root := pq.nodes[0]
	last := len(pq.nodes) - 1

	pq.swap(0, last)
	pq.pos[root] = -1
	pq.nodes = pq.nodes[:last]

	if len(pq.nodes) > 0 {
		pq.down(0)
	}
	return root
}

func (pq *IndexedPQ) up(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if pq.dist[pq.nodes[i]] >= pq.dist[pq.nodes[p]] {
			break
		}
		pq.swap(i, p)
		i = p
	}
}

func (pq *IndexedPQ) down(i int) {
	n := len(pq.nodes)
	for {
		left := 2*i + 1
		if left >= n || left < 0 {
			break
		}
		smallest := left
		if right := left + 1; right < n && pq.dist[pq.nodes[right]] < pq.dist[pq.nodes[left]] {
			smallest = right
		}
		if pq.dist[pq.nodes[i]] <= pq.dist[pq.nodes[smallest]] {
			break
		}
		pq.swap(i, smallest)
		i = smallest
	}
}

func (pq *IndexedPQ) swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
	pq.pos[pq.nodes[i]] = i
	pq.pos[pq.nodes[j]] = j
}

func DijkstraIPQ(graph *Graph, source int) []int {
	n := len(graph.edges)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	ipq := NewIndexedPQ(n, dist)
	ipq.PushOrUpdate(source)

	for len(ipq.nodes) > 0 {
		u := ipq.Pop()

		for _, edge := range graph.edges[u] {
			if newDist := dist[u] + edge.weight; newDist < dist[edge.to] {
				dist[edge.to] = newDist
				ipq.PushOrUpdate(edge.to)
			}
		}
	}
	return dist
}
