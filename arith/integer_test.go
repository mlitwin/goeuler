package arith

import (
	"math/big"
	"testing"
)

func TestIntModP(t *testing.T) {
	m := NewIntModM(2)
	if m.m != 2 {
		t.Fatal("Failed to NewIntModM", m)
	}
	var x int64
	m.Let(&x, 3)
	if x != 1 {
		t.Fatal("Failed to Let", 3, m)
	}
}

func TestBigInt(t *testing.T) {
	b := NewBigInt()
	v := big.NewInt(2)
	x := big.NewInt(0)
	b.Let(x, 3)

	c := b.Cmp(*x, *v)

	if c != 1 {
		t.Fatal("Failed Let BigInt")
	}

	b.Neg(x, *v)
	c = b.Cmp(*x, *v)

	if c != -1 {
		t.Fatal("Failed Neg BigInt")
	}
}
