package textutil

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Create a bufio.Scanner from a file name.
//
// Usage:
//
// 	scanner, close := NewFileScanner("p054_poker.txt")
// 	defer close()
// 	for scanner.Scan() {
// 		line := scanner.Text()
//	}
func NewFileScanner(filename string) (*bufio.Scanner, func()) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file), func() {
		file.Close()
	}
}

func ReadMatrix(filename string, separator string) ([][]int64, int, int) {
	var ret [][]int64

	scanner, close := NewFileScanner(filename)
	defer close()

	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Split(line, separator)
		nums := make([]int64, len(numStrings))

		for i, v := range numStrings {
			nums[i], _ = strconv.ParseInt(v, 10, 64)
		}
		ret = append(ret, nums)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret, len(ret), len(ret[0])
}
