package searching

import (
	"testing"
)

func TestRotationPoint(t *testing.T) {
	words := []string{
		"grape", "orange", "plum", "radish", "apple",
	}
	res := FindRotationPoint(words)
	if res != 4 {
		t.Errorf("Expected %v, got %v", 4, res)
	}
}
