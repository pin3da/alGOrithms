package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func markRows(board []string, id int) ([][]int, int) {
	n := len(board)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				id++
			} else {
				row[j] = id
			}
		}
		id++
		result[i] = row
	}

	return result, id
}

func markCols(board []string, id int) ([][]int, int) {
	n := len(board)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if board[i][j] == 'X' {
				id++
			} else {
				result[i][j] = id
			}
		}
		id++
	}

	return result, id
}

// Edge ...
type Edge struct {
	from, to  int
	cap, flow int64
}

// Dinic ...
type Dinic struct {
	N     int
	edges []Edge
	g     [][]int
	d, pt []int
}

// NewDinic ...
func NewDinic(N int) *Dinic {
	return &Dinic{N, make([]Edge, 0), make([][]int, N), make([]int, N), make([]int, N)}
}

// AddEdge ...
func (g *Dinic) AddEdge(from, to, cap int) {
	g.edges = append(g.edges, Edge{from, to, int64(cap), int64(0)})
	g.g[from] = append(g.g[from], len(g.edges)-1)
	g.edges = append(g.edges, Edge{to, from, int64(0), int64(0)})
	g.g[to] = append(g.g[to], len(g.edges)-1)
}

func (g *Dinic) bfs(s, t int) bool {
	cur := []int{s}
	next := []int{}

	for i := range g.d {
		g.d[i] = g.N + 10
	}
	g.d[s] = 0

	for len(cur) > 0 {
		for _, node := range cur {
			for _, id := range g.g[node] {
				e := &g.edges[id]
				if e.flow < e.cap && g.d[e.to] > g.d[e.from]+1 {
					g.d[e.to] = g.d[e.from] + 1
					next = append(next, e.to)
				}
			}
		}
		cur = next
		next = []int{}
	}
	return g.d[t] != (g.N + 10)
}

func (g *Dinic) dfs(node, T int, flow int64) int64 {
	if node == T || flow == 0 {
		return flow
	}
	for ; g.pt[node] < len(g.g[node]); g.pt[node]++ {
		id := g.g[node][g.pt[node]]
		e := &g.edges[id]
		oe := &g.edges[id^1]
		if g.d[e.from] == g.d[e.to]-1 {
			amt := min64(e.cap-e.flow, flow)
			if pushed := g.dfs(e.to, T, amt); pushed > 0 {
				e.flow += amt
				oe.flow -= amt
				return pushed
			}

		}
	}
	return 0
}

// MaxFlow ...
func (g *Dinic) MaxFlow(source, sink int) int64 {
	total := int64(0)
	for g.bfs(source, sink) {
		for i := range g.pt {
			g.pt[i] = 0
		}
		for flow := g.dfs(source, sink, int64(1<<60)); flow > 0; flow = g.dfs(source, sink, int64(1<<60)) {
			total += flow
		}
	}
	return total
}

func main() {
	cin := NewReader()
	for n := cin.NextInt(); !cin.Ended; n = cin.NextInt() {
		board := make([]string, n)
		for i := 0; i < n; i++ {
			board[i] = cin.Next()
		}
		curID := 0
		rowID, curID := markRows(board, curID)
		colID, curID := markCols(board, curID)
		source := curID
		sink := source + 1
		nodes := sink + 1

		graph := NewDinic(nodes)

		rows, cols := make(map[int]bool), make(map[int]bool)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == '.' {
					graph.AddEdge(rowID[i][j], colID[i][j], 1)
					rows[rowID[i][j]] = true
					cols[colID[i][j]] = true
				}

			}
		}

		for r := range rows {
			graph.AddEdge(source, r, 1)
		}

		for c := range cols {
			graph.AddEdge(c, sink, 1)
		}

		fmt.Println(graph.MaxFlow(source, sink))
	}

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
