package algo

import (
	"math/rand"
	"sort"
	"testing"
)

var r = rand.New(rand.NewSource(10))

type sortCase struct {
	given    []int
	expected []int
}

func newsortCase(size int, n int) *sortCase {
	ret := &sortCase{}
	for i := 0; i < size; i++ {
		v := rand.Intn(n)
		ret.given = append(ret.given, v)
		ret.expected = append(ret.expected, v)
	}
	sort.Ints(ret.expected)

	return ret
}

func (s *sortCase) scratch() []int {
	ret := make([]int, len(s.given))
	copy(ret, s.given)
	return ret
}

func (s *sortCase) test(result []int) bool {
	if len(result) != len(s.expected) {
		return false
	}
	for i, v := range result {
		if v != s.expected[i] {
			return false
		}
	}
	return true
}

var cases []*sortCase

func TestInit(t *testing.T) {
	const N = 100
	for i := 0; i < N; i++ {
		c := newsortCase(10, 5)
		cases = append(cases, c)
	}

	for i := 0; i < N; i++ {
		c := newsortCase(10, 2)

		cases = append(cases, c)
	}

	c := newsortCase(10, 1)
	cases = append(cases, c)

}

func doTestSort(t *testing.T, sort func([]int)) {
	for _, c := range cases {
		s := c.scratch()
		sort(s)
		if !c.test(s) {
			t.Fatal(s, c)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	doTestSort(t, InsertionSort)
}

func TestSelectionSort(t *testing.T) {
	doTestSort(t, SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	doTestSort(t, BubbleSort)
}

func TestQuicksortLomuto(t *testing.T) {
	doTestSort(t, QuicksortLomuto)
}

func TestQuicksortLomutoDutchFlag(t *testing.T) {
	doTestSort(t, QuicksortLomutoDutchFlag)
}

func TestQuicksortHoar(t *testing.T) {
	doTestSort(t, QuicksortHoar)
}
