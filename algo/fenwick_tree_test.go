package algo

import (
	"math/rand"
	"testing"
)

func sumRuns(t *testing.T, n int64) {
	f := NewFenwickTree(n)
	perm := rand.Perm(int(n + 1))

	var i int64
	for i = 0; i <= n; i++ {
		index := int64(perm[int(i)])
		f.Update(index, 1)
	}

	for i = 0; i <= n; i++ {
		sum := f.Read(i)
		if sum != i+1 {
			t.Fatal("bad run of", n, " at ", i, " got ", sum)
		}
	}

}

func TestFenwickTree(t *testing.T) {
	f := NewFenwickTree(10)

	r := f.Read(0)
	if r != 0 {
		t.Fatal("bad init")
	}

	r = f.Read(1)
	if r != 0 {
		t.Fatal("bad init")
	}

	f.Update(1, 1)

	r = f.Read(1)
	if r != 1 {
		f.Debug()
		t.Fatal("bad read", r)
	}

	f.Update(2, 1)

	r = f.Read(1)
	if r != 1 {
		f.Debug()
		t.Fatal("bad nd 2read", r)
	}

	r = f.Read(2)
	if r != 2 {
		f.Debug()
		t.Fatal("bad nd 2read", r)
	}

	sumRuns(t, 10)
}
