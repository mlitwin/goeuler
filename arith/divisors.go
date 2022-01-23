package arith

type Divisors struct {
	n   int64
	cur int64
}

func NewDivisors(n int64) *Divisors {
	d := Divisors{n: n}
	d.cur = 1

	return &d
}

func (d *Divisors) HasValue() bool {
	return d.cur <= d.n
}

func (d *Divisors) NextValue() int64 {
	ret := d.cur

	d.cur++

	for ; d.cur <= d.n; d.cur++ {
		if d.n%d.cur == 0 {
			break
		}
	}

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
