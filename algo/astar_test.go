package algo

import (
	"testing"
)

// Basic grid
type node struct {
	AStarVerexImp
	g [][]node
	n int
	i int
	j int
}

func (n node) IsEnd() bool {
	return n.i == (n.n-1) && n.j == (n.n-1)
}

func (n node) Visit(visit func(neighbor AStarVerex, weight int)) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {

			if i == 0 && j == 0 {
				continue
			}

			x := n.i + i
			y := n.j + j

			if x < 0 || x >= n.n || y < 0 || y >= n.n {
				continue
			}

			dy := n.n - y - 1

			visit(&n.g[x][y], x*dy*dy)
		}
	}
}

func (n node) Heuristic() int {
	return 0
}

func TestAStarBasic(t *testing.T) {
	var nodes = make([][]node, 10)
	for i := range nodes {
		nodes[i] = make([]node, 10)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			nodes[i][j].g = nodes
			nodes[i][j].n = 10
			nodes[i][j].i = i
			nodes[i][j].j = j
		}
	}

	wt, path := MinPathAStar(&nodes[0][0])

	if wt != 0 {
		t.Fatal("Wrong weight", wt)
	}

	if len(path) != 19 {
		t.Fatal("Wrong path", path)
	}
}
