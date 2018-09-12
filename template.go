package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	cin := NewReader()
}

/// TEMPLATE

// Reader is a fast input for programming competitions
type Reader struct {
	scanner *bufio.Scanner
	Ended   bool
}

// NewReader ...
func NewReader() *Reader {
	var reader Reader
	reader.scanner = bufio.NewScanner(os.Stdin)
	reader.scanner.Split(bufio.ScanWords)
	reader.Ended = false
	return &reader
}

// NextInt ...
func (r *Reader) NextInt() int {
	tmp := r.Next()
	if !r.Ended {
		res, _ := strconv.Atoi(tmp)
		return res
	}
	return 0
}

// Next ...
func (r *Reader) Next() string {
	if r.scanner.Scan() {
		return r.scanner.Text()
	}
	r.Ended = true
	return ""
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
