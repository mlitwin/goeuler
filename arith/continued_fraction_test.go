package arith

import (
	"testing"
)

func testConvergent(t *testing.T, D int64, n int, expected RationalFraction) {
	cf, d := NewRationalSurd(D)
	var f []RationalFraction

	for i := 0; i <= n; i++ {
		a := cf.NextCFTerm(d)
		f = cf.NextCFConvergent(f, a)
	}
	v := f[len(f)-1]
	if expected != v {
		t.Fatal("Wrong convergent", D, n, v)
	}
}

func testConvergentOf(t *testing.T, D int64, n int, expected RationalFraction) {
	cf, d := NewRationalSurd(D)
	dl := NewDigitList(10)
	var f [][]int64

	for i := 0; i <= n; i++ {
		a := cf.NextCFTerm(d)
		f = NextCFConvergentOf[[]int64](*dl, f, a)
	}

	v := NewRationalFraction(dl.ValueOfDigits(f[2]), dl.ValueOfDigits(f[3]))
	if expected != *v {
		t.Fatal("Wrong convergent", D, n, *v)
	}
}

func TestCF(t *testing.T) {
	testConvergent(t, 2, 10, RationalFraction{8119, 5741})
	testConvergentOf(t, 2, 10, RationalFraction{8119, 5741})
	testConvergent(t, 3, 10, RationalFraction{989, 571})
	testConvergentOf(t, 3, 10, RationalFraction{989, 571})
}
