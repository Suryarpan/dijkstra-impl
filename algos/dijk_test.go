package algos

import (
	"math/rand"
	"testing"
)

func generateRandomGraph(V, E int, maxWeight int) *Graph {
	g := NewGraph(uint(V))

	// 1. Ensure Connectivity (Simple path: 0->1->2...->V-1)
	// This prevents Dijkstra from finishing early due to unreachable nodes.
	for i := 0; i < V-1; i++ {
		weight := rand.Intn(maxWeight) + 1
		g = g.AddEdge(i, i+1, weight)
	}

	// 2. Fill remaining edges to reach target density E
	// Use a map to avoid duplicate edges between the same two nodes
	edgesAdded := V - 1
	for edgesAdded < E {
		u := rand.Intn(V)
		v := rand.Intn(V)

		// Avoid self-loops and simple duplicates for cleaner benchmarking
		if u != v {
			weight := rand.Intn(maxWeight) + 1
			g = g.AddEdge(u, v, weight)
			edgesAdded++
		}
	}

	return g
}

func BenchmarkDijkstra(b *testing.B) {
	graph := generateRandomGraph(10000, 1000000, 1000)

	b.Run("Standard_Heap", func(b *testing.B) {
		for b.Loop() {
			Dijkstra(graph, 0)
		}
	})

	b.Run("Custom Min Heap", func(b *testing.B) {
		for b.Loop() {
			DijkstraSSSP(graph, 0)
		}
	})

	b.Run("Indexed_PQ", func(b *testing.B) {
		for b.Loop() {
			DijkstraIPQ(graph, 0)
		}
	})
}
