// Package algo comprises algorithms conceptually more complicated than those in package arith
// These will generally be graph theoretical, sorting and searching, etc., rather than more
// conceptually arithmetic ones. algo can use arith, but not conversely.
package algo

// AStarGraph defines the interface needed by the caller
// They need a graph object which traffics in vertex id's, which need to be comparable
// in  order to be able to store auxilliary information about vertices.
/*

A `MinPathAStar[ID comparable, W Numeric](g AStarGraph[ID,W], start ID, end ID) (W, []ID) ` function. Takes an `AStarGraph`, start and end vertex, returns the min weight, and the path.

An interesting design question here is how to handle the auxiliary data the algorithm needs to store about each vertex. Here we require the `AStarGraph` interface to traffic with a comparable `ID` for each vertex, so the algorithm can use that as a key to an (internal) map.

Another way to go would be to require the `AStarGraph` to be able to store (and produce) the auxiliary data itself. It seemed like most implementation would end up with some kind of map anyway, which is why I didn't go this route.
*/
type AStarGraph[ID comparable, W Numeric] interface {
	Heuristic(v ID) W
	Visit(v ID, visit func(neighbor ID, weight W))
}

type aStarVerexState[ID comparable, W Numeric] struct {
	v        ID
	visited  bool
	score    W
	heapNode *HeapNode[ID, W]
	prev     *aStarVerexState[ID, W]
}

func (s aStarVerexState[ID, W]) wouldBeBetter(score W) bool {
	return !s.visited || score < s.score
}

func (s *aStarVerexState[ID, W]) setScore(score W) {
	s.score = score
	s.visited = true
}

type minPathAStarImp[ID comparable, W Numeric] struct {
	g     AStarGraph[ID, W]
	start ID
	end   ID

	vertexState map[ID]*aStarVerexState[ID, W]
	openSet     *Heap[ID, W]
}

// getState is a convenience method to get vertex state *, with autovivification
func (imp minPathAStarImp[ID, W]) getState(v ID) *aStarVerexState[ID, W] {
	ret := imp.vertexState[v]
	if nil == ret { // vivify
		ret = &aStarVerexState[ID, W]{}
		ret.v = v
		imp.vertexState[v] = ret
	}

	return ret
}

func (imp minPathAStarImp[ID, W]) isEnd(v ID) bool {
	return v == imp.end
}

// Backtrack to find the actual path that was used
func (imp minPathAStarImp[ID, W]) backtrackPath(v ID) []ID {
	var path []ID
	s := imp.getState(v)

	for s != nil {
		path = append([]ID{s.v}, path...)
		s = s.prev
	}

	return path
}

func newminPathAStarImp[ID comparable, W Numeric](g AStarGraph[ID, W], start ID, end ID) minPathAStarImp[ID, W] {
	imp := minPathAStarImp[ID, W]{g, start, end, nil, nil}
	imp.vertexState = make(map[ID]*aStarVerexState[ID, W])
	imp.openSet = NewHeap[ID, W]()

	return imp
}

func MinPathAStar[ID comparable, W Numeric](g AStarGraph[ID, W], start ID, end ID) (W, []ID) {
	imp := newminPathAStarImp(g, start, end)
	imp.openSet.Push(start, 0)

	for imp.openSet.Len() > 0 {
		cur := imp.openSet.Pop()
		curState := imp.getState(cur.value)
		curState.visited = true

		if imp.isEnd(cur.value) {
			return curState.score, imp.backtrackPath(cur.value)
		}

		g.Visit(cur.value, func(neighbor ID, weight W) {
			score := curState.score + weight
			nState := imp.getState(neighbor)

			if nState.wouldBeBetter(score) {
				nState.setScore(score)
				remainingPathEstimate := score + g.Heuristic(neighbor)
				nState.prev = curState
				nState.heapNode = imp.openSet.Upsert(neighbor, remainingPathEstimate, nState.heapNode)
			}
		})

	}

	return 0, nil
}
