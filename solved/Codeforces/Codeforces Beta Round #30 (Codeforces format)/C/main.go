package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Target struct {
	x, y, t int
	p       float64
}

func dist(a, b Target) float64 {
	x := float64(a.x - b.x)
	y := float64(a.y - b.y)
	return math.Sqrt(x*x + y*y)
}

func Go(data []Target, i, last int, dp [][]float64) float64 {
	if i == len(data) {
		return float64(0)
	}
	if dp[i][last] != -1 {
		return dp[i][last]
	}
	best := Go(data, i+1, last, dp)
	if last >= len(data) || dist(data[i], data[last]) <= float64(data[i].t-data[last].t) {
		best = math.Max(best, Go(data, i+1, i, dp)+data[i].p)
	}
	dp[i][last] = best
	return best
}

func main() {
	cin := NewReader()
	n := cin.Int()
	data := make([]Target, n)
	for i := 0; i < n; i++ {
		data[i].x = cin.Int()
		data[i].y = cin.Int()
		data[i].t = cin.Int()
		data[i].p = cin.Float64()
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].t < data[j].t
	})

	dp := make([][]float64, n+2)
	for i := range dp {
		dp[i] = make([]float64, n+2)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	fmt.Printf("%.10f\n", Go(data, 0, n+1, dp))
}

/* TEMPLATE */

type Reader struct {
	scanner *bufio.Scanner
	Ended   bool
}

func NewReader() *Reader {
	var reader Reader
	reader.scanner = bufio.NewScanner(os.Stdin)
	reader.scanner.Split(bufio.ScanWords)
	reader.Ended = false
	return &reader
}

func (r *Reader) Int() int {
	tmp := r.Next()
	if !r.Ended {
		res, _ := strconv.Atoi(tmp)
		return res
	}
	return 0
}

func (r *Reader) Float64() float64 {
	tmp := r.Next()
	if !r.Ended {
		res, _ := strconv.ParseFloat(tmp, 64)
		return res
	}
	return 0
}

func (r *Reader) Array(len int) []int {
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[i] = r.Int()
	}
	return result
}

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
