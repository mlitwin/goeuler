package algo

import (
	"fmt"
	"testing"
)

type index struct {
	i, j int
}

type vertex struct {
	pos index
}

type grid struct {
	n int
}

func (g grid) Heuristic(v index) int {
	return 0
}

func (g grid) Visit(v index, visit func(neighbor index, weight int)) {
	n := g.n
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			if i == 0 && j == 0 {
				continue
			}

			neighbor := index{v.i + i, v.j + j}

			if neighbor.i < 0 || neighbor.i >= n || neighbor.j < 0 || neighbor.j >= n {
				continue
			}

			wt := 100

			if neighbor.i == 0 || neighbor.j == n-1 {
				wt = 1
			}

			visit(neighbor, wt)
		}
	}
}

func TestAStarBasic(t *testing.T) {
	g := grid{10}
	start := index{0, 0}
	end := index{9, 9}

	wt, path := MinPathAStar[index, int](g, start, end)

	if wt != 17 {
		t.Fatal("Wrong weight", wt)
	}

	if len(path) != 18 {
		for _, v := range path {
			fmt.Println(v)
		}
		t.Fatal("Wrong path", len(path), path, wt)
	}

}
