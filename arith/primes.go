package arith

// Basic primality test
func IsPrime(x int64) bool {
	if x <= 1 {
		return false
	}
	var d int64
	for d = 2; d*d <= x; d++ {
		if x%d == 0 {
			return false
		}
	}
	return true
}

type PrimeFactorization struct {
	n int64
	p int64
}

func NewPrimeFactorization(n int64) *PrimeFactorization {
	return &PrimeFactorization{n, 2}
}

func (f PrimeFactorization) HasValue() bool {
	return f.n > 1
}

func (f *PrimeFactorization) NextValue() (int64, int64) {
	var r int64
	for f.n%f.p != 0 {
		f.p++
	}

	for f.n%f.p == 0 {
		f.n /= f.p
		r++
	}

	return f.p, r
}
