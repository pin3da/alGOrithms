package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Road struct {
	from, to, time int
}

func isPossible(demand []int, stock []int, roads []Road, maxTime int64) bool {
	source := len(demand) + len(stock)
	sink := source + 1
	graph := NewDinic(sink + 10)

	target := int64(0)
	for _, t := range demand {
		target += int64(t)
	}

	for i, d := range stock {
		graph.AddEdge(source, i, d)
	}
	for i, d := range demand {
		graph.AddEdge(i+len(stock), sink, d)
	}

	for _, road := range roads {
		if int64(road.time) <= maxTime {
			graph.AddEdge(road.from, road.to+len(stock), demand[road.to])
		}
	}

	flow := graph.MaxFlow(source, sink)
	return flow == target
}

func main() {
	cin := NewReader()
	P := cin.Int()
	R := cin.Int()
	Q := cin.Int()
	demand := cin.Array(P)
	stock := cin.Array(R)
	roads := make([]Road, Q)

	for i := range roads {
		r := &roads[i]
		r.to = cin.Int() - 1
		r.from = cin.Int() - 1
		r.time = cin.Int()
	}

	if !isPossible(demand, stock, roads, 2000000) {
		fmt.Println(-1)
	} else {
		lo := int64(0)
		hi := int64(20000000)
		for lo < hi {
			mid := (lo + hi) >> 1
			if isPossible(demand, stock, roads, mid) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		fmt.Println(lo)
	}
}

// Dinic

type Edge struct {
	from, to  int
	cap, flow int64
}

type Dinic struct {
	N     int
	edges []Edge
	g     [][]int
	d, pt []int
}

func NewDinic(N int) *Dinic {
	return &Dinic{N, make([]Edge, 0), make([][]int, N), make([]int, N), make([]int, N)}
}

func (g *Dinic) AddEdge(from, to, cap int) {
	g.edges = append(g.edges, Edge{from, to, int64(cap), int64(0)})
	g.g[from] = append(g.g[from], len(g.edges)-1)
	g.edges = append(g.edges, Edge{to, from, int64(0), int64(0)})
	g.g[to] = append(g.g[to], len(g.edges)-1)
}

func (g *Dinic) bfs(s, t int) bool {
	cur := []int{s}
	next := []int{}

	inf := g.N + 10
	for i := range g.d {
		g.d[i] = inf
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
	return g.d[t] != inf
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
				e.flow += pushed
				oe.flow -= pushed
				return pushed
			}

		}
	}
	return 0
}

func (g *Dinic) MaxFlow(source, sink int) int64 {
	total := int64(0)
	for g.bfs(source, sink) {
		for i := range g.pt {
			g.pt[i] = 0
		}
		for flow := g.dfs(source, sink, int64(1<<60)); flow > 0; {
			total += flow
			flow = g.dfs(source, sink, int64(1<<60))
		}
	}
	return total
}

// TEMPLATE

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
