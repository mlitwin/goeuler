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
