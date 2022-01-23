package arith

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
