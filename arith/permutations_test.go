package arith

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	perms := NewPermutations(3)
	if nil == perms {
		t.Fatal("Can't construct a permutation")
	}
	var count = 0
	for perms.HasValue() {
		count++
		perms.GetValue()
		perms.NextValue()
	}
	if count != 6 {
		t.Fatal("Wrong number of permutations of 3")
	}
}
