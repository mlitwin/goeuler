package algo

import (
	"testing"
)

func TestGridDag(t *testing.T) {
	g := NewGridDag()

	if g == nil {
		t.Fatal("Can't construct GridDag")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var w int64
			if i == 1 && j == 1 {
				w++
			}
			g.AddVertex(i, j, w)
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			n := GridNeighbors(i, j, 3, 3)
			for _, v := range n {
				g.AddEdge(i, j, v.I, v.J)
			}
		}
	}

	wt, path := g.MinPathAStar(0, 0, 2, 2)

	if wt != 0 {
		t.Fatal("Wrong pathness weight", wt, path)
	}
}
