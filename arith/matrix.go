package arith

// Basic Matrix type
type Matrix[T any] [][]T

func NewMatrix[T any](m int, n int) Matrix[T] {
	r := make([][]T, m)
	for j := 0; j < m; j++ {
		r[j] = make([]T, n)
	}

	return r
}
