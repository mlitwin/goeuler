package algo

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type HeapNode[V any, P Numeric] struct {
	priority P
	index    int
	value    V
}

type Heap[V any, P Numeric] struct {
	p []*HeapNode[V, P]
}

func (h *Heap[V, P]) Len() int {
	return len(h.p)
}

func (h *Heap[V, P]) Push(x V, priority P) *HeapNode[V, P] {
	n := &HeapNode[V, P]{priority, len(h.p), x}
	h.p = append(h.p, n)
	h.swim(len(h.p) - 1)
	return n
}

func (h *Heap[V, P]) Pop() *HeapNode[V, P] {
	ret := h.p[0]
	last := len(h.p) - 1

	h.p[0] = h.p[last]
	h.p[0].index = 0
	h.p = h.p[:last]

	if last > 0 {
		h.sink(0)
	}

	return ret
}

func (h *Heap[V, P]) Decrease(n *HeapNode[V, P], priority P) {
	if priority > n.priority {
		panic("Min Heap does not support an increase priority operation")
	}
	n.priority = priority
	h.swim(n.index)
}

// Convenience fucnction to Push if new, change priority if exists
func (h *Heap[V, P]) Upsert(x V, priority P, n *HeapNode[V, P]) *HeapNode[V, P] {
	if nil == n {
		return h.Push(x, priority)
	}

	h.Decrease(n, priority)

	return n
}

func NewHeap[V any, P Numeric]() *Heap[V, P] {
	h := Heap[V, P]{}

	return &h
}

func (h *Heap[V, P]) Validate() bool {
	var valid bool = true
	p := h.p
	l := len(p)

	for i := 0; i < len(p); i++ {
		left := left(i)
		if left < l && p[left].priority < p[i].priority {
			valid = false
		}

		right := right(i)
		if right < l && p[right].priority < p[i].priority {
			valid = false
		}
	}

	return valid
}

func (h *Heap[V, P]) swap(i int, j int) {
	t := h.p[i]
	h.p[i] = h.p[j]
	h.p[j] = t

	index := h.p[i].index
	h.p[i].index = h.p[j].index
	h.p[j].index = index
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *Heap[V, P]) sink(i int) {
	p := h.p

	cur := i

	for {
		var candidate int
		l := left(cur)

		if l >= len(p) { // cur is a leaf
			return
		}

		r := right(cur)

		if r >= len(p) { // one child
			candidate = l
		} else { // two children
			if p[l].priority <= p[r].priority {
				candidate = l
			} else {
				candidate = r
			}
		}

		if p[cur].priority <= p[candidate].priority { // cur is good
			return
		}

		h.swap(cur, candidate)
		cur = candidate
	}

}

func (h *Heap[V, P]) swim(i int) {
	p := h.p

	cur := i
	for cur >= 0 {
		up := parent(cur)
		if up < 0 || p[up].priority <= p[cur].priority {
			return
		}
		h.swap(cur, up)
		cur = up
	}
}
