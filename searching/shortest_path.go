package searching

import (
	"github.com/aponte411/data-structures/linear/queue"
	"sort"
)

// TODO: import graph from data-structures/graph

func FindAShortestPath(gph map[string][]string, src, dst string) []string {
	// Raise error if src or dst are not in edges
	paths := make(map[string]string, 0)
	paths[src] = "#"
	que := new(queue.QueueLinkedList)
	que.Enqueue(src)
	for !que.IsEmpty() {
		node := que.Dequeue().(string)
		if node == dst {
			// reconstruct path from dst node
			return reconstructPath(paths, node)
		}
		for i := range gph[node] {
			child := gph[node][i]
			if _, ok := paths[child]; !ok {
				paths[child] = node
				que.Enqueue(child)
			}
		}
	}
	return nil
}

func reconstructPath(paths map[string]string, end string) []string {
	shortest_path := make([]string, 0)
	node := end
	for node != "#" {
		shortest_path = append(shortest_path, node)
		node = paths[node]

	}
	sort.Sort(sort.Reverse(sort.StringSlice(shortest_path)))
	return shortest_path
}
