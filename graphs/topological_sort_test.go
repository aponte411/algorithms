package graphs

import (
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	var preRequisites = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	ordered_courses := topoSort(preRequisites)
	if len(ordered_courses) == 0 {
		t.Error("Expected ordered courses to have at least 1 course")
	}
	for order, course := range ordered_courses {
		fmt.Printf("%d:\t%s\n", order+1, course)
	}

}
