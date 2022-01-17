package arith

type Divisors struct {
	n int64
	cur int64
}

func NewDivisors(n int64) *Divisors {
	d := Divisors{n:n}
	d.cur = 1

	return &d
}

func (d *Divisors) HasValue() bool {
	return d.cur <= d.n
}

func (d *Divisors) NextValue() int64 {
	ret := d.cur

	d.cur++
	
	for ; d.cur <= d.n; d.cur++{
		if(d.n % d.cur == 0) {
			break;
		}
	}

	return ret
}