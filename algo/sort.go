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

import (
	"fmt"
)

var _ = fmt.Println

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
	partition func([]int) ([]int, []int)
}

func (q quickSorter) sort(array []int) {
	if len(array) <= 1 {
		return
	}
	left, right := q.partition(array)
	q.sort(left)
	q.sort(right)
}

func lomutoPartition(a []int) ([]int, []int) {
	var leftEnd int
	last := len(a) - 1
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

func dutchFlagPartition(a []int) ([]int, []int) {
	var lowInsert, equalInsert, highInsert int

	last := len(a) - 1
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

func hoarPartition(a []int) ([]int, []int) {
	var low, high int

	high = len(a) - 1
	pivot := a[high/2]

	for {
		fmt.Println(low, high)
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

func QuicksortLomuto(a []int) {
	q := quickSorter{lomutoPartition}
	q.sort(a)
}

func QuicksortLomutoDutchFlag(a []int) {
	q := quickSorter{dutchFlagPartition}
	q.sort(a)
}

func QuicksortHoar(a []int) {
	q := quickSorter{hoarPartition}
	q.sort(a)
}
