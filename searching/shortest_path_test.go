package searching

import (
	"reflect"
	"testing"
)

func TestShortestPath(t *testing.T) {
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
	exp := []string{"Adam", "Amelia", "Jayden"}
	res := FindAShortestPath(graph, src, dst)
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v, got %v", exp, res)
	}
}
