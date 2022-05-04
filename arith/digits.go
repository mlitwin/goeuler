package arith

//  Return the digits of n base b as a slice
func digits(n int64, b int64) []int64 {
	var ret []int64

	for n > 0 {
		d := n % b
		ret = append([]int64{d}, ret...)
		n /= b
	}

	return ret
}

// Convert digit slice back to `int64`
func valueOfDigits(n []int64, b int64) (ret int64) {
	i := len(n) - 1
	var place int64 = 1
	for i >= 0 {
		ret += n[i] * place
		i--
		place *= b
	}

	return
}

// Return the base b digit list n as a proper base b number, each digit in range, and no leading zeros.
func normalizeDigits(n []int64, b int64) []int64 {
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

	var firstNonzeroIndex int
	for firstNonzeroIndex < len(ret) {
		if ret[firstNonzeroIndex] != 0 {
			break
		}
		firstNonzeroIndex++
	}

	return ret[firstNonzeroIndex:]
}

// Manage integers as lists of digits with a given base.
//
// This is basically an inefficent BigInt, intended as an exercise
// in the basic algorithms of digit based arithmetic, and for problems
// which primarily invove digit manipulation for other reasons.
//
// Conforms to Integer interface
type DigitList struct {
	base int64
}

// Constructor for DigitList
func NewDigitList(base int64) *DigitList {
	return &DigitList{base}
}

func (d DigitList) normalize(x *[]int64) {
	*x = normalizeDigits(*x, d.base)
}

// x = a (int64)
func (d DigitList) Let(x *[]int64, a int64) {
	*x = digits(a, d.base)
}

// x = a
func (d DigitList) Set(x *[]int64, a []int64) {
	*x = make([]int64, len(a))
	copy(*x, a)
}

// x = -a
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

// x = a + b
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

// x = a - b
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

// x = a * b
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

// x = a / b (integer division)
//
// BUG(mlitwin): Not actually implemented.
func (d DigitList) Div(x *[]int64, a []int64, b []int64) {
	t := make([]int64, len(a))

	*x = t

	d.normalize(x)

	// TBD long division https://en.wikipedia.org/wiki/Long_division
}

// Comparison: a X b (-1 means <; 0 means ==; 1 means >)
func (d DigitList) Cmp(a []int64, b []int64) int {
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	return 0
}

// Create digit list representation of n.
func (d DigitList) Digits(n int64) []int64 {
	return digits(n, d.base)
}

// Compute integer value represented by list of digits.
func (d DigitList) ValueOfDigits(n []int64) (ret int64) {

	return valueOfDigits(n, d.base)
}

// Histogram of digit values
func (d DigitList) Histogram(n []int64) (ret []int64) {
	ret = make([]int64, d.base)
	for _, v := range n {
		ret[v]++
	}

	return
}
