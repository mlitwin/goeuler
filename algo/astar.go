// Package algo comprises algorithms conceptually more complicaticated than those in package arith
// These will generally be graph theoretical, sorting and searching, etc., rather than more
// conceptually arithmetic ones. algo can use arith, but not conversely.
package algo

// AStarGraph defines the interface needed by the caller
// They need a graph object which traffics in verticies.
// In order to be able to store auxilliary information about verticies
// there also needs to be a comparable vertex ID
type AStarGraph[V any, ID comparable, W Numeric] interface {
	GetId(v *V) ID
	Heuristic(v *V) W
	Visit(v *V, visit func(neighbor *V, weight W))
}

type aStarVerexState[V any, W Numeric] struct {
	v *V
	visited bool
	CurPathWeight  W
	heapNode       *HeapNode[*V,W]
	prev *aStarVerexState[V,W]
}

func (s aStarVerexState[V, W]) wouldBeBetter(score W) bool {
	return !s.visited || score < s.CurPathWeight
}

type minPathAStarImp[V any, ID comparable, W Numeric] struct {
	g AStarGraph[V,ID,W]
	start *V
	end *V

	endId ID
	vertexState map[ID]*aStarVerexState[V, W]
	openSet *Heap[*V,W]
}


// A map[]*aStarVerexState
// neededing this function proves I still don't understand something about
// golang maps
func (imp minPathAStarImp[V,ID,W]) getState(v *V) *aStarVerexState[V,W] {
	id := imp.g.GetId(v)
	ret := imp.vertexState[id]
	if nil == ret {
		ret = &aStarVerexState[V,W]{}
		ret.v = v
		imp.vertexState[id] = ret
	}

	return ret
}

func (imp minPathAStarImp[V,ID,W]) isEnd(v *V) bool {
	return imp.g.GetId(v) == imp.endId
}



// Backtrack to find the actual path that was used
func (imp minPathAStarImp[V,ID,W]) backtrackPath(v *V) []*V {
	var path []*V
	s := imp.getState(v)

	for s != nil {
		path = append([]*V{s.v}, path...)
		s = s.prev
	}

	return path
}

func newminPathAStarImp[V any, ID comparable, W Numeric](g AStarGraph[V,ID,W], start *V, end *V) minPathAStarImp[V,ID,W] {
	imp := minPathAStarImp[V,ID,W]{g, start, end,  g.GetId(end), nil, nil}
	imp.vertexState =  make(map[ID]*aStarVerexState[V,W])
	imp.openSet = NewHeap[*V,W]()


	return imp
}

func MinPathAStar[V any, ID comparable, W Numeric](g AStarGraph[V,ID,W], start *V, end *V) (W, []*V) {
	imp := newminPathAStarImp[V,ID,W](g, start, end)
	imp.openSet.Push(start, 0)

	for imp.openSet.Len() > 0 {
		cur := imp.openSet.Pop()
		curState := imp.getState(cur.value)
		curState.visited = true

		if imp.isEnd(cur.value) {
			return curState.CurPathWeight, imp.backtrackPath(cur.value)
		}

		g.Visit(cur.value, func(neighbor *V, weight W) {
			score := curState.CurPathWeight + weight
			nState := imp.getState(neighbor)

			if nState.wouldBeBetter(score) {
				remainingPathEstimate := score + g.Heuristic(neighbor)
				nState.prev = curState
				nState.CurPathWeight = score
				nState.heapNode = imp.openSet.Upsert(neighbor, remainingPathEstimate, nState.heapNode)
				nState.visited = true
			}
		})

	}

	return 0, nil
}
