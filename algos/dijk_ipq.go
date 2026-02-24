package algos

import (
	"math"
)

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
