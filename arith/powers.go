package arith

const MaxInt64 int64 = int64((^uint64(0)) >> 1)
const MaxSqrtInt64 int64 = 3037000499
const MaxSquareInt64 int64 = MaxSqrtInt64 * MaxSqrtInt64

// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
func Pow(x int64, n int64) int64 {

	if n == 0 {
		return 1
	}

	var y int64 = 1
	for n > 1 {
		if n%2 == 0 {
			x *= x
			n /= 2
		} else {
			y = x * y
			x *= x
			n = (n - 1) / 2
		}
	}

	return x * y
}

//https://en.wikipedia.org/wiki/Integer_square_root
func IntSqrt(n int64) int64 {
	var l, m, r int64

	if n >= MaxSquareInt64 {
		return MaxSqrtInt64
	}

	r = n + 1

	for l != r-1 {
		// overflow safe m = (l + r) / 2
		m = l/2 + r/2 + (l%2+r%2)/2
		if m <= MaxSqrtInt64 && m*m <= n {
			l = m
		} else {
			r = m
		}
	}

	return l
}
