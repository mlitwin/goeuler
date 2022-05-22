package arith

import (
	"testing"
)

func DoTestOneCombinations(t *testing.T, n int64, k int64) {
	c := NewCombinations(n, k)

	var count int64
	for ; c.HasValue(); c.NextValue() {
		_ = c.GetValue()
		count++
	}
	cc := C(n, k)
	if count != cc {
		t.Fatal("Wrong iteraton count", n, k, count, "expected", cc)
	}
}
func TestCombinations(t *testing.T) {
	DoTestOneCombinations(t, 3, 2)
	DoTestOneCombinations(t, 10, 5)
}
