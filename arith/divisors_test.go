package arith

import (
	"testing"
)

func divisorsSum(t *testing.T, n int64) (ret int64) {
	f := NewDivisors(n)

	for f.HasValue() {
		d := f.NextValue()
		if n%d != 0 {
			t.Fatal(d, "not a divisor of", n)
		}
		ret += d
	}

	return
}

func simpleDivisorSum(n int64) (ret int64) {
	var i int64
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			ret += i
		}
	}
	return
}

func TestDivisors(t *testing.T) {
	var n int64
	for n = 1; n <= 100; n++ {
		simp := simpleDivisorSum(n)
		got := divisorsSum(t, n)
		if got != simp {
			t.Fatal("Wrong divisors for", n, "expected", simp, "got", got)
		}
	}
}
