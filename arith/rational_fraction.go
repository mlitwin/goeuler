package arith

// A RationalFraction A/B supporting extraction of arbitrary base `NextMantissaDigit()`
type RationalFraction struct {
	A int64
	B int64
}

func NewRationalFraction(a int64, b int64) *RationalFraction {
	r := RationalFraction{a, b}
	return &r
}

// Assuming proper fraction, spit out the next digit base whatever,
// and advance to to remainder
func (r *RationalFraction) NextMantissaDigit(base int64) int64 {
	// A/B = k/base, so k = A*base / B
	num := r.A * base
	div := num / r.B
	rem := num % r.B

	r.A = rem

	return div
}
