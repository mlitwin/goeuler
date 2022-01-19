package algo

import (
	"testing"
	"fmt"
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

func (g grid) GetId(v *vertex) index {
	return v.pos
}

func (g grid) Heuristic(v *vertex) int {
	return 0
}

func (g grid) Visit(v *vertex, visit func(neighbor *vertex, weight int)) {
	n := g.n
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			if i == 0 && j == 0 {
				continue
			}

			neighbor := vertex{index{v.pos.i + i, v.pos.j + j}}

			if neighbor.pos.i < 0 || neighbor.pos.i >= n || neighbor.pos.j < 0 || neighbor.pos.j >= n {
				continue
			}

			wt := 100

			if neighbor.pos.i == 0 || neighbor.pos.j == n -1 {
				wt = 1;
			}

			visit(&neighbor, wt)
		}
	}
}

func TestAStarBasic(t *testing.T) {
	g := grid{10}
	start := vertex{index{0,0}}
	end := vertex{index{9,9}}

	wt, path := MinPathAStar[vertex, index, int](&g, &start, &end)

	if wt != 17 {
		t.Fatal("Wrong weight", wt)
	}

	if len(path) != 18 {
		for _,v:= range(path) {
			fmt.Println(v)
		}
		t.Fatal("Wrong path", len(path), path, wt)
	}


}
