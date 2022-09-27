package arith

import (
	"testing"
)

func divisorsSum(n int64) (ret int64) {
	f := NewDivisors(n)

	for f.HasValue() {
		d := f.NextValue()
		ret += d
	}

	return
}

func TestDivisors(t *testing.T) {
	if divisorsSum(28) != 2*28 {
		t.Fatal("Wrong divisors for", 28)
	}
}
