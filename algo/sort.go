package algo

// Not for actual use, but to practice with the concepts

/*
Phase 1
Work with []int
	* Set up test harness
	* Simple methods
	  * Insertion
	  * Selection
	  * Bubble
	* Strong-able methods
	  * Quicksort
	  * Mergesort
	  * https://en.wikipedia.org/wiki/Heapsort

Work with generics
* https://pkg.go.dev/sort#Interface seems pretty good
  * only thing would be if it was worth parameterizing the index type


*/

func InsertionSort(input []int) {
	for i := 1; i < len(input); i++ {
		for j := i - 1; j >= 0; j-- {
			if input[j] < input[j+1] {
				break
			}
			input[j], input[j+1] = input[j+1], input[j]
		}
	}
}

func SelectionSort(input []int) {
	for i := 0; i < len(input)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(input); j++ {
			if input[j] < input[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			input[i], input[minIndex] = input[minIndex], input[i]
		}
	}
}

func BubbleSort(input []int) {
	n := len(input)
	done := false
	for !done {
		done = true
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				input[i-1], input[i] = input[i], input[i-1]
				done = false
			}
		}
		n--
	}
}

type quickSorter struct {
	pivot     func([]int) int                 // find a pivot, maybe messing with the input array
	partition func([]int, int) ([]int, []int) // split into 2 arrays around the pivot value
}

func (q quickSorter) sort(array []int) {
	if len(array) <= 1 {
		return
	}
	pivotIndex := q.pivot(array)
	left, right := q.partition(array, pivotIndex)
	q.sort(left)
	q.sort(right)
}

func lomutoPartition(a []int, pivotIndex int) ([]int, []int) {
	var leftEnd int
	last := pivotIndex // must be len(a) - 1 for lumuto
	pivot := a[last]
	for j := 0; j < last; j++ {
		if a[j] < pivot {
			a[leftEnd], a[j] = a[j], a[leftEnd]
			leftEnd++
		}
	}
	a[leftEnd], a[last] = a[last], a[leftEnd]

	return a[:leftEnd], a[leftEnd+1:]
}

func dutchFlagPartition(a []int, pivotIndex int) ([]int, []int) {
	var lowInsert, equalInsert, highInsert int

	last := pivotIndex // must be len(a) - 1 for lumuto
	pivot := a[last]

	highInsert = last

	for equalInsert <= highInsert {
		if a[equalInsert] < pivot {
			a[lowInsert], a[equalInsert] = a[equalInsert], a[lowInsert]
			lowInsert++
			equalInsert++
		} else if pivot < a[equalInsert] {
			a[highInsert], a[equalInsert] = a[equalInsert], a[highInsert]
			highInsert--
		} else {
			equalInsert++

		}
	}

	return a[:lowInsert], a[equalInsert:]

}

func hoarPartition(a []int, pivotIndex int) ([]int, []int) {
	var low, high int

	high = len(a) - 1
	pivot := a[pivotIndex]

	for {
		for a[low] < pivot {
			low++
		}
		for pivot < a[high] {
			high--
		}
		if low >= high {
			break
		}
		a[low], a[high] = a[high], a[low]
		low++
		high--
	}

	return a[:low], a[high+1:]

}

func lumutoPivot(a []int) int {
	return len(a) - 1
}

func lumutoMedianPivot(a []int) int {
	mid := len(a) / 2
	last := len(a) - 1

	if a[mid] < a[0] {
		a[0], a[mid] = a[mid], a[0]
	}
	if a[last] < a[0] {
		a[0], a[last] = a[last], a[0]
	}

	if a[last] < a[mid] {
		a[mid], a[last] = a[last], a[mid]
	}

	return last
}

func QuicksortLomuto(a []int) {
	q := quickSorter{lumutoPivot, lomutoPartition}
	q.sort(a)
}

func QuicksortLomutoDutchFlag(a []int) {
	q := quickSorter{lumutoMedianPivot, dutchFlagPartition}
	q.sort(a)
}

func hoarPivot(a []int) int {
	return (len(a) - 1) / 2
}

func QuicksortHoar(a []int) {
	q := quickSorter{hoarPivot, hoarPartition}
	q.sort(a)
}

type mergeSorter struct {
	buf []int
}

func newMergeSorter(n int) *mergeSorter {
	ret := mergeSorter{}
	ret.buf = make([]int, n)

	return &ret
}

// merge 2 sorted lists a and b into dest
func (m *mergeSorter) merge(dest []int, a []int, b []int) {
	var i, j, n int

	// grab from a and b in order
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			m.buf[n] = a[i]
			i++
			n++
		} else {
			m.buf[n] = b[j]
			j++
			n++
		}
	}

	// any left over in a?
	for i < len(a) {
		m.buf[n] = a[i]
		i++
		n++
	}

	// any left in b?
	for j < len(b) {
		m.buf[n] = b[j]
		j++
		n++
	}

	for k := 0; k < n; k++ {
		dest[k] = m.buf[k]
	}
}

func (m *mergeSorter) mergeSortTopDown(a []int) {
	if len(a) <= 1 {
		return
	}

	mid := len(a) / 2
	left := a[:mid]
	right := a[mid:]

	m.mergeSortTopDown(left)
	m.mergeSortTopDown(right)
	m.merge(a, left, right)
}

func MergeSortTopDown(a []int) {
	m := newMergeSorter(len(a))
	m.mergeSortTopDown(a)
}

func MergeSortBottomUp(a []int) {
	m := newMergeSorter(len(a))

	for n := 1; n < len(a); n *= 2 {
		for k := 0; k+n < len(a); k += 2 * n {
			mid := k + n
			right := k + 2*n
			if right > len(a) {
				right = len(a)
			}
			m.merge(a[k:right], a[k:mid], a[mid:right])
		}
	}
}

func CountSort[T any](a []T, bucket func(T) int, bucketCount int) {
	buckets := make([]int, bucketCount)

	for _, v := range a {
		buckets[bucket(v)]++
	}

	for i := 1; i < len(buckets); i++ {
		buckets[i] += buckets[i-1]
	}

	tmp := make([]T, len(a))

	for i := len(a) - 1; i >= 0; i-- {
		v := a[i]
		b := bucket(v)
		tmp[buckets[b]-1] = v
		buckets[b]--
	}

	copy(a, tmp)
}

func RadixSort(a []int) {
	const base = 10
	var max, min int

	if len(a) > 0 {
		min, max = a[0], a[0]
		for _, v := range a {
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}
	count := max - min

	for exp := 1; exp <= count; exp *= base {
		CountSort(
			a,
			func(v int) int {
				b := (v - min) / exp
				return b % base
			},
			base,
		)
	}
}
