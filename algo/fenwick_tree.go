package algo

import (
	"fmt"
)

// https://www.topcoder.com/thrive/articles/Binary%20Indexed%20Trees
// https://cp-algorithms.com/data_structures/fenwick.html

type FenwickTree struct {
	highIndex int64
	tree      map[int64]int64
}

func NewFenwickTree(highIndex int64) *FenwickTree {
	return &FenwickTree{highIndex, make(map[int64]int64)}
}

func lsval(idx int64) int64 {
	return idx & (-idx)
}

func (f FenwickTree) Read(idx int64) (ret int64) {
	for idx > 0 {
		ret += f.tree[idx]
		idx -= lsval(idx)
	}

	return
}

func (f *FenwickTree) Update(idx int64, val int64) {
	for idx <= f.highIndex {
		f.tree[idx] += val
		idx += lsval(idx)
	}
}

func (f FenwickTree) Debug() {
	fmt.Println(f.highIndex, f.tree)
}
