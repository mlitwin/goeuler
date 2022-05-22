package algo

import (
	"testing"
)

func TestGridDag(t *testing.T) {
	var m [][]int64 = [][]int64{[]int64{0, 0, 0}, []int64{0, 1, 0}, []int64{0, 0, 0}}
	g := NewGridDag(m)

	if g == nil {
		t.Fatal("Can't construct GridDag")
	}

	g.VisitAllNeighbors(func(i0, j0, i1, j1 int) {
		g.AddEdge(i0, j0, i1, j1)
	})

	wt, path := g.MinPathAStar(0, 0, 2, 2)

	if wt != 0 {
		t.Fatal("Wrong pathness weight", wt, path)
	}
}
