package arith

// A factoradic base number, supporting conversion to a permutation.
// https://en.wikipedia.org/wiki/Factorial_number_system
type Factoradic struct {
	Value  int64
	Digits []int64
}

// Create the factoradic representation of n.
func NewFactoradic(n int64) *Factoradic {
	f := Factoradic{n, nil}
	var place int64 = 1
	for n > 0 {
		rem := n % place
		f.Digits = append([]int64{rem}, f.Digits...)
		n /= place
		place++
	}
	return &f
}

// Convert the factoradic to a permutation of n items.
func (f Factoradic) Permutation(n int64) []int64 {
	var digits []int64
	var perm []int64
	var i int64
	for i = 0; i < n; i++ {
		digits = append(digits, i)
	}

	for i = 0; i < n; i++ {
		ind := i - (n - int64(len(f.Digits)))
		var choice int64
		var digitIndex int64
		if ind < 0 {
			choice = i
			digitIndex = 0
		} else {
			digitIndex = f.Digits[ind]
			choice = digits[digitIndex]
		}
		digits = append(digits[:digitIndex], digits[digitIndex+1:]...)
		perm = append(perm, choice)
	}

	return perm
}
