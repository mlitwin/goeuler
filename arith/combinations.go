package arith

import (
	"math"
)

// n Choose k
// https://cp-algorithms.com/combinatorics/binomial-coefficients.html
func C(n int64, k int64) int64 {
	fn := float64(n)
	fk := float64(k)
	var res float64 = 1

	var i int64
	for i = 1; i <= k; i++ {
		fi := float64(i)
		res = res * (fn - fk + fi) / fi
	}
	return (int64)(math.Floor(res + 0.5))
}
