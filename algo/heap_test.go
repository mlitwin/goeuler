package algo

import (
	"math/rand"
	"testing"
)

func TestHeapCreate(t *testing.T) {
	h := NewHeap()
	if h.Len() != 0 {
		t.Fatal("Could not create Heap")
	}
}

func TestHeapPush(t *testing.T) {
	h := NewHeap()
	h.Push(0, 1)
	if h.Len() != 1 {
		t.Fatal("Could not create Push")
	}

	h.Pop()

	if h.Len() != 0 {
		t.Fatal("Could not Pop a Push")
	}
}

func TestHeapOrderedPushPops(t *testing.T) {
	h := NewHeap()

	for i := 0; i < 10; i++ {
		h.Push(i, i)
		if !h.Validate() {
			t.Fatal("Could not do multiple Push")
		}
	}

	for i := 0; i < 10; i++ {
		h.Pop()
		if !h.Validate() {
			t.Fatal("Could not do multiple Pop")
		}
	}
}

func doPushPops(count int, r *rand.Rand, n int, t *testing.T) {
	h := NewHeap()
	var nodes []*HeapNode

	for i := 0; i < count; i++ {
		var priority int
		if n == 0 {
			priority = r.Int()
		} else {
			priority = r.Intn(n)
		}
		node := h.Push(i, priority)
		nodes = append(nodes, node)
		if !h.Validate() {
			t.Fatal("Could not do multiple Push")
		}
	}

	// Twiddle the priorities
	for _, v := range nodes {
		var priority int
		if n == 0 {
			priority = r.Int()
		} else {
			priority = r.Intn(n)
		}
		if priority < v.priority {
			h.Decrease(v, priority)
		}
		if !h.Validate() {
			t.Fatal("Could not Decrease")
		}
	}

	for i := 0; i < count; i++ {
		h.Pop()
		if !h.Validate() {
			t.Fatal("Could not do multiple Pop")
		}
	}
}

func TestHeapRandomizedPushPops(t *testing.T) {

	doPushPops(100, rand.New(rand.NewSource(2)), 0, t)
	doPushPops(100, rand.New(rand.NewSource(2)), 10, t)
	doPushPops(10, rand.New(rand.NewSource(2)), 1, t)
}
