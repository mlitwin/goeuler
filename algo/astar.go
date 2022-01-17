package algo

type AStarVerexState struct {
	CurPathWeight  int
	heapNode       *HeapNode
	previousVertex AStarVerex
}

func (s AStarVerexState) wouldBeBetter(score int) bool {
	if s.heapNode == nil {
		return true
	}

	return score < s.CurPathWeight
}

type AStarVerex interface {
	GetState() *AStarVerexState
	IsEnd() bool
	Visit(visit func(neighbor AStarVerex, weight int))
	Heuristic() int
}

type AStarVerexImp struct {
	state AStarVerexState
}

func (v *AStarVerexImp) GetState() *AStarVerexState {
	return &v.state
}

func getHeapVertex(node *HeapNode) AStarVerex {
	value := node.value
	vertex := value.(AStarVerex)
	return vertex
}

// Find the actual path that was used
func backtrackPath(s AStarVerex) []AStarVerex {
	var path []AStarVerex
	path = append(path, s)

	for s != nil {
		path = append([]AStarVerex{s}, path...)
		s = s.GetState().previousVertex
	}

	return path
}

func MinPathAStar(start AStarVerex) (int, []AStarVerex) {
	openSet := NewHeap()
	state := start.GetState()
	state.heapNode = openSet.Push(start, 0)

	for openSet.Len() > 0 {
		cur := getHeapVertex(openSet.Pop())

		curState := cur.GetState()

		if cur.IsEnd() {
			return curState.CurPathWeight, backtrackPath(cur)
		}

		cur.Visit(func(neighbor AStarVerex, weight int) {
			score := curState.CurPathWeight + weight
			nState := neighbor.GetState()
			if nState.wouldBeBetter(score) {
				remainingPathEstimate := score + neighbor.Heuristic()
				nState.previousVertex = cur
				nState.CurPathWeight = score
				nState.heapNode = openSet.Upsert(neighbor, remainingPathEstimate, nState.heapNode)
			}
		})
	}

	return 0, nil
}
