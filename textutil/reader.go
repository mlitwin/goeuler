package textutil

import (
	"bufio"
	"os"
)

func NewFileScanner(filename string) (*bufio.Scanner, func()) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return bufio.NewScanner(file), func() {
		file.Close()
	}
}
