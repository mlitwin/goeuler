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

func TestIntSqrt(t *testing.T) {
	var k int64
	for k = 0; k < 20; k++ {
		s := IntSqrt(k * k)
		if s != k {
			t.Fatal("Bad IntSqrt of k^2", k, s)
		}
	}

	for k = 1; k < 20; k++ {
		s := IntSqrt(k*k + 1)
		if s != k {
			t.Fatal("Bad IntSqrt of k^2+1", k, s)
		}
	}

	k = IntSqrt(MaxInt64)
	if MaxSqrtInt64 != k {
		t.Fatal("Bad IntSqrt of MaxInt64", k)
	}

	k = IntSqrt(MaxSquareInt64)
	if MaxSqrtInt64 != k {
		t.Fatal("Bad IntSqrt of MaxSquareInt64", k)
	}

	s := MaxSqrtInt64 - 1
	k = IntSqrt(s * s)
	if s != k {
		t.Fatal("Bad IntSqrt of MaxSquareInt64 - 1", k, s, s-k)
	}
}

func TestPowOf(t *testing.T) {
	m := NewIntModM(7)
	var r, x int64

	m.Let(&x, 2)

	r = PowOf[int64](*m, x, 6)

	if r != 1 {
		t.Fatal("Bad 2^6 mod 7", r)
	}
	N := Pow(10, 10)

	var m1 = NewIntModM(N)

	r = PowOf[int64](*m1, 17, 17)

	if r != 6336764177 {
		t.Fatal("Bad 17^17 mod 10^10", r)
	}

}
