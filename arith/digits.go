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

// Convert digit slice back to `int64`
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

// Count digits in a slice
func HistogramOfDigits(n []int64, b int64) []int64 {
	ret := make([]int64, b)

	for _, d := range n {
		ret[d]++
	}

	return ret
}

func NormalizeDigits(n []int64, b int64) []int64 {
	var ret []int64
	var i int
	var r int64

	for i = len(n) - 1; i >= 0; i-- {
		v := n[i] + r
		r = v / b
		d := v % b
		ret = append([]int64{d}, ret...)
	}

	for r != 0 {
		d := r % b
		r /= b
		ret = append([]int64{d}, ret...)
	}

	return ret
}

// Integer interface

type DigitList struct {
	base int64
}

func NewDigitList(base int64) *DigitList {
	return &DigitList{base}
}

func (d DigitList) normalize(x *[]int64) {
	*x = NormalizeDigits(*x, d.base)
}

func (d DigitList) Let(x *[]int64, a int64) {
	*x = Digits(a, d.base)
}

func (d DigitList) Set(x *[]int64, a []int64) {
	*x = make([]int64, len(a))
	copy(*x, a)
}

func (d DigitList) Neg(x *[]int64, a []int64) {
	d.Set(x, a)
	for i, v := range *x {
		(*x)[i] = -v
	}

	d.normalize(x)
}

func makemax(a []int64, b []int64) []int64 {
	var l int
	if len(a) > len(b) {
		l = len(a)
	} else {
		l = len(b)
	}

	return make([]int64, l)
}

func (d DigitList) Sum(x *[]int64, a []int64, b []int64) {
	t := makemax(a, b)

	for i, v := range a {
		dig_a := len(a) - 1 - i
		t[len(t)-1-dig_a] += v
	}
	for i, v := range b {
		dig_b := len(b) - 1 - i
		(t)[len(t)-1-dig_b] += v
	}

	*x = t
	d.normalize(x)
}

func (d DigitList) Diff(x *[]int64, a []int64, b []int64) {
	t := makemax(a, b)

	for i, v := range a {
		dig_a := len(a) - 1 - i
		t[len(t)-1-dig_a] += v
	}
	for i, v := range b {
		dig_b := len(b) - 1 - i
		t[len(t)-1-dig_b] -= v
	}

	*x = t
	d.normalize(x)

}

func (d DigitList) Mul(x *[]int64, a []int64, b []int64) {
	t := make([]int64, len(a)+len(b))

	for i, v := range a {
		dig_i := len(a) - 1 - i
		for j, w := range b {
			dig_j := len(b) - 1 - j
			dig_x := len(t) - 1 - (dig_i + dig_j)
			t[dig_x] += v * w
		}
	}

	*x = t
	d.normalize(x)

}

func (d DigitList) Div(x *[]int64, a []int64, b []int64) {
	t := make([]int64, len(a))

	*x = t

	d.normalize(x)

	// TBD long division https://en.wikipedia.org/wiki/Long_division
}

func (d DigitList) Cmp(x *[]int64, a []int64, b []int64) int {
	var l int
	if len(a) > len(b) {
		l = len(a)
	} else {
		l = len(b)
	}

	for i := l - 1; i >= 0; i-- {
		var da, db int64

		if i < len(a) {
			da = a[i]
		}

		if i < len(b) {
			db = b[i]
		}

		if da < db {
			return -1
		}

		if db > da {
			return 1
		}
	}

	return 0
}
