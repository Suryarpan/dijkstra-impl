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

	fmt.Println("Shortest distances std:", algos.Dijkstra(graph, 0))
	fmt.Println("Shortest Distances m heap:", algos.DijkstraSSSP(graph, 0))
	fmt.Println("Shortest Distances i pq:", algos.DijkstraIPQ(graph, 0))
}
