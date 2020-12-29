package searching

import (
	"fmt"
	"testing"
)

func TestShortestPath(t *testing.T) {
	graph := map[string][]string{
		"": []string{},
		"": []string{},
		"": []string{},
	}
	src := "David"
	dst := "Pati"
	res := FindAShortestPath(graph, src, dst)
	fmt.Println(res)
}
