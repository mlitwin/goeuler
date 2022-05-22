package algo

type GridIndex struct {
	I, J int
}

// Generic Graph / DAG based on a grid
type GridDag struct {
	e map[GridIndex][]*GridIndex // edges
	w map[GridIndex]int64        // vertex weights
}

func NewGridDag(m [][]int64) *GridDag {
	g := GridDag{}
	g.w = make(map[GridIndex]int64)
	g.e = make(map[GridIndex][]*GridIndex)

	for i, row := range m {
		for j, v := range row {
			g.AddVertex(i, j, v)
		}
	}

	return &g
}

func (g *GridDag) AddVertex(i, j int, w int64) {
	index := GridIndex{i, j}
	g.w[index] = w
}

func (g *GridDag) AddEdge(i0, j0 int, i1, j1 int) {
	index0 := GridIndex{i0, j0}
	index1 := GridIndex{i1, j1}

	g.e[index0] = append(g.e[index0], &index1)
}

// Generic Graph interface - here the index is the same type as the vertex

func (g GridDag) GetId(v *GridIndex) GridIndex {
	return *v
}

func (g GridDag) Heuristic(v *GridIndex) int64 {
	return 0
}

func (g GridDag) Visit(v *GridIndex, visit func(neighbor *GridIndex, weight int64)) {
	for _, e := range g.e[*v] {
		w := g.w[*e]
		visit(e, w)
	}
}

func (g GridDag) MinPathAStar(i0, j0 int, i1, j1 int) (int64, []GridIndex) {
	start := GridIndex{i0, j0}
	end := GridIndex{i1, j1}
	wt, path := MinPathAStar[GridIndex, GridIndex, int64](g, &start, &end)
	retpath := make([]GridIndex, len(path))
	for i, v := range path {
		retpath[i] = *v
	}

	return wt, retpath
}

// Convenience utility to visit the potential neighbors of verticies
func (g GridDag) VisitAllNeighbors(rows, cols int, visit func(i0, j0 int, i1, j1 int)) {
	for i0 := 0; i0 < rows; i0++ {
		for j0 := 0; j0 < cols; j0++ {
			for i1 := i0 - 1; i1 <= i0+1; i1++ {
				for j1 := j0 - 1; j1 <= j0+1; j1++ {
					if i1 == i0 && j1 == j0 {
						continue
					}

					if i1 < 0 || j1 < 0 || i1 >= rows || j1 >= cols {
						continue
					}

					visit(i0, j0, i1, j1)
				}
			}
		}
	}

}

// Convenience utility to get the potential neighbors of an element
func GridNeighbors(i0, j0 int, w, h int) (ret []GridIndex) {
	for i := i0 - 1; i <= i0+1; i++ {
		for j := j0 - 1; j <= j0+1; j++ {
			if i == i0 && j == j0 {
				continue
			}

			if i < 0 || j < 0 || i >= w || j >= h {
				continue
			}
			ret = append(ret, GridIndex{i, j})
		}
	}

	return
}
