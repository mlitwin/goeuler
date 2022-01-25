package algo

import (
	"testing"
)

func TestFenwickTree(t *testing.T) {
	f := NewFenwickTree()

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

	f.Update(200, 200)

	r = f.Read(200)
	if r != 202 {
		f.Debug()
		t.Fatal("bad far read", r)
	}
}
