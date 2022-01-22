package arith

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
