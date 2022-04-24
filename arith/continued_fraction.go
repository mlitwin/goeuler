package arith

import "math"

/*
https://en.wikipedia.org/wiki/Quadratic_irrational_number
*/

// Operations on (a+b*sqrt(D))/c
type RationalSurd struct {
	D     int64   // surd
	sqrtD float64 // sqrt(D)
}

// Instance of (a+b*sqrt(D))/c
type RationalSurdValue struct {
	a, b, c int64
}

func NewRationalSurd(D int64) (*RationalSurd, *RationalSurdValue) {
	sqrtd := math.Sqrt(float64(D))
	surd := RationalSurd{
		D:     D,
		sqrtD: sqrtd,
	}
	v := RationalSurdValue{0, 1, 1}

	return &surd, &v
}

// floor(d)
func (s *RationalSurd) IntFloor(d *RationalSurdValue) int64 {
	f := (float64(d.a) + float64(d.b)*s.sqrtD) / float64(d.c)
	return int64(f)
}

// 1/d
func (s *RationalSurd) Invert(d *RationalSurdValue) {
	a := d.a * d.c
	b := -d.b * d.c
	c := d.a*d.a - s.D*d.b*d.b

	div := GCD(GCD(a, c), GCD(b, c))

	d.a = a / div
	d.b = b / div
	d.c = c / div
}

// Next term in continued fraction
func (s *RationalSurd) NextCFTerm(d *RationalSurdValue) int64 {
	i := s.IntFloor(d)
	d.a -= i * d.c
	s.Invert(d)

	return i
}

// Next convergent for continued fraction
// start with an empty []RationalFraction slice. Convergent will be the last
// in the returned slice. Slice is kept for recurrence relation.
func (s *RationalSurd) NextCFConvergent(cur []RationalFraction, a int64) []RationalFraction {
	ret := make([]RationalFraction, 0, 2)

	cur0 := NewRationalFraction(0, 1)
	cur1 := NewRationalFraction(1, 0)

	if len(cur) > 0 {
		*cur0 = cur[0]
		*cur1 = cur[1]
	}

	next := NewRationalFraction(a*cur1.A+cur0.A, a*cur1.B+cur0.B)

	return append(ret, *cur1, *next)
}
