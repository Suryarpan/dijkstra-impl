package main

import (
	"fmt"

	"github.com/Suryarpan/dijkstra-impl/algos"
)

func main() {
	graph := algos.NewGraph(4)
	graph = graph.
		AddEdge(0, 1, 5).
		AddEdge(0, 2, 1).
		AddEdge(1, 3, 1).
		AddEdge(2, 1, 2).
		AddEdge(2, 3, 7)
	fmt.Println("Shortest Distances:", algos.DijkstraIPQ(graph, 0))
}
