package arith

import "math/big"

// generic "Integer" operations
// Not a mathematics integer, but a computer science integer - basically has addition/subtraction/multiplication/divistion
type Integer[V any] interface {
	Let(x *V, a int64)
	Set(x *V, a V)
	Neg(x *V, a V)
	Sum(x *V, a V, b V)
	Diff(x *V, a V, b V)
	Mul(x *V, a V, b V)
	Div(x *V, a V, b V)
	Cmp(a V, b V) int
}

// Integer Modulo m, represented in an int64
type IntModM struct {
	m int64
}

func NewIntModM(m int64) *IntModM {
	return &IntModM{m}
}

func (m IntModM) Let(x *int64, a int64) {
	if a >= 0 {
		*x = a % m.m

	} else {
		*x = -a%m.m + m.m
	}
}

func (m IntModM) Set(x *int64, a int64) {
	*x = a
}

func (m IntModM) Neg(x *int64, a int64) {
	*x = -a % m.m
}

// a*b%m with better handling of overflow
// m <= 2^63
// https://www.geeksforgeeks.org/modular-exponentiation-power-in-modular-arithmetic/?ref=lbp
func mulMod(a int64, b int64, n int64) (ret int64) {

	for b > 0 {
		if b%2 != 0 {
			ret = (ret + a) % n
		}

		a = (2 * a) % n
		b /= 2
	}

	return
}

func (m IntModM) Sum(x *int64, a int64, b int64) {
	*x = (a + b) % m.m
}

func (m IntModM) Diff(x *int64, a int64, b int64) {
	*x = (a - b) % m.m
}

func (m IntModM) Mul(x *int64, a int64, b int64) {
	const max32 = 1 >> 32
	if m.m <= max32 {
		*x = (a * b) % m.m
	} else {
		*x = mulMod(a, b, m.m)
	}
}

func (m IntModM) Div(x *int64, a int64, b int64) {
	if b > 0 {
		b = InverseModN(b, m.m)
	} else {
		b = -InverseModN(-b, m.m) + m.m
	}

	*x = (a * b) % m.m
}

func (m IntModM) Cmp(a int64, b int64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	}

	return 0
}

// Integer[big.Int]
type BigInt struct {
}

func NewBigInt() *BigInt {
	return &BigInt{}
}

func (b BigInt) Let(x *big.Int, a int64) {
	x.SetInt64(a)
}

func (b BigInt) Set(x *big.Int, a big.Int) {
	x.Set(&a)
}
func (b BigInt) Neg(x *big.Int, a big.Int) {
	x.Neg(&a)
}
func (bb BigInt) Sum(x *big.Int, a big.Int, b big.Int) {
	x.Add(&a, &b)
}
func (bb BigInt) Diff(x *big.Int, a big.Int, b big.Int) {
	x.Sub(&a, &b)
}
func (bb BigInt) Mul(x *big.Int, a big.Int, b big.Int) {
	x.Mul(&a, &b)
}
func (bb BigInt) Div(x *big.Int, a big.Int, b big.Int) {
	x.Div(&a, &b)
}
func (bv BigInt) Cmp(a big.Int, b big.Int) int {
	return a.Cmp(&b)
}
