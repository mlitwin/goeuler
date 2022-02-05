package arith

// Return the digits of n base b as a slice
func Digits(n int64, b int64) []int64 {
	var ret []int64

	for n > 0 {
		d := n % b
		ret = append([]int64{d}, ret...)
		n /= b
	}

	return ret
}

func ValueOfDigits(n []int64, b int64) (ret int64) {
	i := len(n) - 1
	var place int64 = 1
	for i >= 0 {
		ret += n[i] * place
		i--
		place *= b
	}

	return
}
