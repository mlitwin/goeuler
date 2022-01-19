package algo


type AStarGraph[V any, ID comparable, W Numeric] interface {
	GetId(v *V) ID
	Heuristic(v *V) W
	Visit(v *V, visit func(neighbor *V, weight W))
}

type AStarVerexState[V any, W Numeric] struct {
	v *V
	visited bool
	CurPathWeight  W
	heapNode       *HeapNode[*V,W]
	prev *AStarVerexState[V,W]
}

func (s AStarVerexState[V, W]) wouldBeBetter(score W) bool {
	if !s.visited {
		return true
	}

	return score < s.CurPathWeight
}

type minPathAStarImp[V any, ID comparable, W Numeric] struct {
	g AStarGraph[V,ID,W]
	start *V
	end *V
	vertexState map[ID]*AStarVerexState[V, W]
	openSet *Heap[*V,W]

}

func (imp minPathAStarImp[V,ID,W]) getState(v *V) *AStarVerexState[V,W] {
	id := imp.g.GetId(v)
	ret := imp.vertexState[id]
	if nil == ret {
		ret = &AStarVerexState[V,W]{}
		imp.vertexState[id] = ret
	}

	return ret
}


// Find the actual path that was used
func (imp minPathAStarImp[V,ID,W]) backtrackPath(v *V) []*V {
	var path []*V
	s := imp.getState(v)


	for s != nil {
		//fmt.Println(s, s.prev)
		path = append([]*V{s.v}, path...)
		s = s.prev
	}

	return path
}

func MinPathAStar[V any, ID comparable, W Numeric](g AStarGraph[V,ID,W], start *V, end *V) (W, []*V) {
	imp := minPathAStarImp[V,ID,W]{g, start, end, nil, nil}
	imp.vertexState =  make(map[ID]*AStarVerexState[V,W])
	imp.openSet = NewHeap[*V,W]()

	imp.openSet.Push(start, 0)


	endId := g.GetId(end)

	for imp.openSet.Len() > 0 {
		cur := imp.openSet.Pop()
		curv := cur.value
		curId := g.GetId(curv)
		curState := imp.getState(curv)
		curState.visited = true
		curState.v = curv


		if curId == endId {
					fmt.Println("backtrack", curId, curState.CurPathWeight)

			return curState.CurPathWeight, imp.backtrackPath(curv)
		}

		fmt.Println("Visiting from ", curId)

		g.Visit(curv, func(neighbor *V, weight W) {
			score := curState.CurPathWeight + weight
			nState := imp.getState(neighbor)
			//fmt.Println(curId, neighbor)

			if nState.wouldBeBetter(score) {
				remainingPathEstimate := score + g.Heuristic(neighbor)
				nState.prev = curState
				nState.CurPathWeight = score
				nState.heapNode = imp.openSet.Upsert(neighbor, remainingPathEstimate, nState.heapNode)
				nState.visited = true
				nState.v = neighbor
				fmt.Println("adding", neighbor, weight, score, curState, nState)

			}
		})

	}

	return 0, nil
}
