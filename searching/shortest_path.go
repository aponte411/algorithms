package searching

// TODO: import graph from data-structures/graph
// and import queue from queue from data-structures/non_linear/queue

func FindAShortestPath(gph map[string][]string, src, dst string) []string {
	// Raise error if src or dst are not in edges
	paths := make(map[string]string, 0)
	que := make([]string, 0)
	que = append(que, src)
	for len(que) != 0 {
		node := que[0]
		que = que[1:]
		if node == dst {
			// reconstruct path from dst node
			return reconstructPath(paths, node)
		}
		for i := range gph[node] {
			child := gph[node][i]
			if _, ok := paths[child]; !ok {
				paths[child] = node
				que = append(que, node)
			}
		}
	}
	return nil
}

func reconstructPath(paths map[string]string, end string) []string {
	shortest_path := make([]string, 0)
	node := end
	for i := 0; i < len(paths); i++ {
		node := paths[node]
		shortest_path = append(shortest_path, node)
	}
	return shortest_path
}
