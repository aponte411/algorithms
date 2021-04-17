package graphs

import (
	"sort"
)

func topoSort(graph map[string][]string) []string {
	// Store nodes in topological order
	var ordered_nodes []string
	// Keep track of seen nodes
	visited := make(map[string]bool)
	// DFS
	var visitAll func(nodes []string)
	visitAll = func(nodes []string) {
		for _, node := range nodes {
			if !visited[node] {
				visited[node] = true
				visitAll(graph[node])
				ordered_nodes = append(ordered_nodes, node)
			}
		}
	}
	var keys []string
	for node := range graph {
		keys = append(keys, node)
	}
	sort.Strings(keys)
	visitAll(keys)
	return ordered_nodes
}
