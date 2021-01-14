package searching

import (
	"fmt"
	"testing"
)

func BenchmarkShortestPath(b *testing.B) {
	graph := map[string][]string{
		"Min":     []string{"William", "Jayden", "Omar"},
		"William": []string{"Min", "Noam"},
		"Jayden":  []string{"Min", "Amelia", "Ren", "Noam"},
		"Ren":     []string{"Jayden", "Omar"},
		"Amelia":  []string{"Jayden", "Adam", "Miguel"},
		"Adam":    []string{"Amelia", "Sofia", "Miguel", "Lucas"},
		"Miguel":  []string{"Amelia", "Adam", "Liam", "Nathan"},
		"Noam":    []string{"Nathan", "Jayden", "William"},
		"Omar":    []string{"Ren", "Min", "Scott"},
	}
	src := "Jayden"
	dst := "Adam"
	for n := 0; n < b.N; n++ {
		res := FindAShortestPath(graph, src, dst)
		fmt.Println(res)
	}
}
