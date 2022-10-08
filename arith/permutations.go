package arith

// Iterate through the permutations
//
//	c := NewPermutations(n)
//
//	for ; c.HasValue(); c.NextValue() {
//		v := c.GetValue()
//	}
// This is not implemented with a fast algorithm,
// we will need https://en.wikipedia.org/wiki/Heap%27s_algorithm
// for that.
type Permutations struct {
	n    int64
	nfac int64
	i    int64
}

func NewPermutations(n int64) *Permutations {
	p := &Permutations{n, Factorial(n), 0}

	return p
}

func (p *Permutations) HasValue() bool {
	return p.i < p.nfac
}

func (p *Permutations) GetValue() []int64 {
	f := NewFactoradic(p.i)
	return f.Permutation(p.n)
}

func (p *Permutations) NextValue() {
	p.i++
}
