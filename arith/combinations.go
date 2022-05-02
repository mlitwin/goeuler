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

// Iterate through the combinations
// https://en.wikipedia.org/wiki/Combination#Enumerating_k-combinations
//
//	c := NewCombinations(n, k)
//
//	for ; c.HasValue(); c.NextValue() {
//		v := c.GetValue()
//	}
type Combinations struct {
	n, k int64
	comb []int64
}

func NewCombinations(n, k int64) *Combinations {
	c := &Combinations{n, k, make([]int64, k)}
	for i := range c.comb {
		c.comb[i] = int64(i)
	}

	return c
}

func (c *Combinations) HasValue() bool {
	return nil != c.comb
}

func (c *Combinations) GetValue() []int64 {
	return c.comb
}

func (c *Combinations) NextValue() {
	lim := c.n - 1

	for i := c.k - 1; i >= 0; i-- {
		if c.comb[i] < lim {
			c.comb[i]++
			for j := i + 1; j < c.k; j++ {
				c.comb[j] = c.comb[j-1] + 1
			}

			return
		}
		lim = c.comb[i] - 1
	}

	c.comb = nil
}
