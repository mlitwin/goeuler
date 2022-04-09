package arith

// Basic abstract iterator of int64's with a `HasValue()` / `NextValue()` interface
type Iterator interface {
	HasValue() bool
	NextValue() int64
}

// A generic `Reduce()` to int64 method for Iterator's
func Reduce(i Iterator, f func(int64, int64) int64, start int64) int64 {
	red := start

	for i.HasValue() {
		cur := i.NextValue()
		red = f(cur, red)
	}

	return red
}
