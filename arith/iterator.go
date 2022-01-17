package arith


type Iterator interface {
	HasValue() bool
	NextValue() int64
}

func Reduce(i Iterator, f func(int64, int64) int64, start int64) int64 {
	red := start

	for i.HasValue() {
		cur := i.NextValue()
		red = f(cur, red)
	}

	return red
}