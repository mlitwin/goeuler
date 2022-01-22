package arith

import (
	"testing"
)

func TestPrimes(t *testing.T) {
	n := int64(7)
	if !IsPrime(n) {
		t.Fatal(n, "is prime")
	}

}
