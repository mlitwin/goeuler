package arith

import (
	"testing"
)

func TestIntModP(t *testing.T) {
	m := NewIntModM(2)
	if m.m != 2 {
		t.Fatal("Falied to NewIntModM", m)
	}
	var x int64
	m.Let(&x, 3)
	if x != 1 {
		t.Fatal("Falied to Let", 3, m)
	}
}
