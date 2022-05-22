package arith

import (
	"testing"
)

func TestPrimes(t *testing.T) {
	n := int64(7)
	if !IsPrime(n) {
		t.Fatal(n, "is prime")
	}

	if IsPrime(1) {
		t.Fatal(1, "is not prime")
	}
}

func TestPrimeFactorization(t *testing.T) {
	f := NewPrimeFactorization(2 * 2 * 3)
	if !f.HasValue() {
		t.Fatal("Can't initialize PrimeFactorization")
	}
	p, r := f.NextValue()

	if p != 2 || r != 2 {
		t.Fatal("Can't find first factor", p, r)

	}

	if !f.HasValue() {
		t.Fatal("Can't Get next")
	}

	p, r = f.NextValue()

	if p != 3 || r != 1 {
		t.Fatal("Can't find second factor", p, r)

	}

	if f.HasValue() {
		t.Fatal("Bad ending")
	}

}
