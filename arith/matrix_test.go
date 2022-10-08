package arith

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	m := NewMatrix[int](2, 2)

	if 0 != m[1][1] {
		t.Fatal("Failed to construct simple matrix")
	}
}
