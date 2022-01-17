package arith

import (
	"reflect"
	"testing"
)

func TestNewFactoradic(t *testing.T) {
	f := NewFactoradic(463)
	if f.Value != 463 {
		t.Fatal("Falied to construct")
	}

	if !reflect.DeepEqual(f.Digits, []int64{3, 4, 1, 0, 1, 0}) {
		t.Fatal("Falied to compute", f.Digits)
	}

}

func TestFactoradicPermutation(t *testing.T) {
	var f *Factoradic
	f = NewFactoradic(4)
	p := f.Permutation(3)
	if !reflect.DeepEqual(p, []int64{2, 0, 1}) {
		t.Fatal("Falied to Permutation n=4, 3")
	}

	f = NewFactoradic(0)
	p = f.Permutation(3)
	if !reflect.DeepEqual(p, []int64{0, 1, 2}) {
		t.Fatal("Falied to Permutation n=3, 0", p)
	}

	f = NewFactoradic(2982)
	p = f.Permutation(7)
	if !reflect.DeepEqual(p, []int64{4, 0, 6, 2, 1, 3, 5}) {
		t.Fatal("Falied to Permutation n=7, 2982", p)
	}
}
