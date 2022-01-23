package arith

import (
	"testing"
)

func TestPow(t *testing.T) {
	n := int64(2)
	if Pow(n, 0) != 1 {
		t.Fatal("Bad pow of 0")
	}

	if Pow(n, 1) != n {
		t.Fatal("Bad pow of 1")
	}

	p := Pow(n, 2)
	if p != n*n {
		t.Fatal("Bad pow of 2", p)
	}

	p = Pow(n, 3)
	if p != n*n*n {
		t.Fatal("Bad pow of 3", p)
	}

}
