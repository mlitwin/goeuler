package arith

// Divisors implements an Iterator through the divisors of n
type Divisors struct {
	n       int64
	sqrtN   int64
	test    int64
	current int64
	saved   int64
}

func NewDivisors(n int64) *Divisors {
	d := Divisors{n: n}
	d.sqrtN = IntSqrt(n)
	d.test = 1
	d.current = 1
	if n > 1 {
		d.saved = n
	}

	return &d
}

func (d *Divisors) HasValue() bool {
	return d.current != 0
}

func (d *Divisors) advance() {
	if d.saved != 0 {
		d.current, d.saved = d.saved, 0
	} else {
		d.test++
		d.current = 0
		for ; d.test <= d.sqrtN; d.test++ {
			if d.n%d.test == 0 {
				d.current = d.test
				d.saved = d.n / d.current
				if d.saved == d.current {
					d.saved = 0
				}
				break
			}
		}

	}
}

func (d *Divisors) NextValue() int64 {

	ret := d.current
	d.advance()

	return ret
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int64) int64 {
	gcd := GCD(a, b)
	return (a * b) / gcd
}

// Modular inverse (or 0 if no inverse)
func InverseModN(a, n int64) int64 {
	var t int64 = 0
	var newt int64 = 1
	var r int64 = n
	var newr int64 = a

	for newr != 0 {
		q := r / newr

		tmp1 := newt
		newt = t - q*newt
		t = tmp1

		tmp2 := newr
		newr = r - q*newr
		r = tmp2
	}

	if r > 1 {
		return 0
	}

	return t + n
}

func Totient(n int64) int64 {
	var result int64 = n
	var i int64
	for i = 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result = result - result/i
		}
	}
	if n > 1 {
		result = result - result/n

	}
	return result
}

// Produce the n'th Farey sequence from the (n-1)th
// https://en.wikipedia.org/wiki/Farey_sequence
//
// You can pass nil to start things off:
//
//	NextFareySequence(nil, 1)
//
// And pass a slice of the previous sequence, if all you are interested in
// is expanding that subsequence.
func NextFareySequence(f []RationalFraction, n int64) []RationalFraction {
	if nil == f {
		return []RationalFraction{{0, 1}, {1, 1}}
	}

	var ret []RationalFraction

	for i, v := range f {
		ret = append(ret, v)
		if i < len(f)-1 {
			q := v.B + f[i+1].B
			if q <= n+1 {
				p := v.A + f[i+1].A
				ret = append(ret, RationalFraction{p, q})
			}
		}
	}

	return ret

}
