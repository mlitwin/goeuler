package textutil

import (
	"bufio"
	"os"
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
