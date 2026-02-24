package algos

import (
	"container/heap"
	"math"
)

// Item is an entry in the priority queue
type Item struct {
	node     int
	priority int
}

// PriorityQueue implements heap.Interface and holds Items
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Dijkstra calculates the shortest distances from a source node
func Dijkstra(graph *Graph, source int) []int {
	dist := make([]int, len(graph.edges))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: source, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		u := current.node
		d := current.priority

		// OPTIMIZATION: If we found a better path already, skip this stale entry
		if d > dist[u] {
			continue
		}

		for _, edge := range graph.edges[u] {
			v := edge.to
			if newDist := dist[u] + edge.weight; newDist < dist[v] {
				dist[v] = newDist
				heap.Push(pq, &Item{node: v, priority: newDist})
			}
		}
	}
	return dist
}
