// TLE ):
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Graph struct {
	N                  int
	cost, tin, tout, A []int
	adj, up            [][]int
	L                  uint
	timer              int
}

func NewGraph(N int) *Graph {
	return &Graph{N, make([]int, N), make([]int, N), make([]int, N), make([]int, 2*N), make([][]int, N), make([][]int, N), 0, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) DFS(node, p int) {
	g.tin[node] = g.timer
	g.A[g.timer] = node
	g.timer++
	g.up[node][0] = p
	for i := uint(1); i <= g.L; i++ {
		g.up[node][i] = g.up[g.up[node][i-1]][i-1]
	}
	for _, to := range g.adj[node] {
		if to != p {
			g.DFS(to, node)
		}
	}
	g.tout[node] = g.timer
	g.A[g.timer] = node
	g.timer++
}

func (g *Graph) IsAncestor(a, b int) bool {
	return g.tin[a] <= g.tin[b] && g.tout[a] >= g.tout[b]
}

func (g *Graph) LCA(a, b int) int {
	if g.IsAncestor(a, b) {
		return a
	}
	if g.IsAncestor(b, a) {
		return b
	}
	for i := int(g.L); i >= 0; i-- {
		if !g.IsAncestor(g.up[a][i], b) {
			a = g.up[a][i]
		}
	}
	return g.up[a][0]
}

func (g *Graph) PrepareLCA(root int) {
	g.L = 1
	for (1 << g.L) <= g.N {
		g.L++
	}
	for i := 0; i < g.N; i++ {
		g.up[i] = make([]int, g.L+1)
	}
	g.timer = 0
	g.DFS(root, root)
}

type Query struct {
	idx, l, r, lca int
}

type DS struct {
	active map[int]bool
	frec   map[int]int
	total  int
}

func (d *DS) toggle(idx int, cost []int) {
	if d.active[idx] {
		d.frec[cost[idx]]--
		if d.frec[cost[idx]] == 0 {
			d.total--
		}
	} else {
		if d.frec[cost[idx]] == 0 {
			d.total++
		}
		d.frec[cost[idx]]++
	}
	d.active[idx] = !d.active[idx]
}

func (d *DS) query() int {
	return d.total
}

type Q []Query

func (a Q) Len() int      { return len(a) }
func (a Q) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Q) Less(i, j int) bool {
	blockSize := 250
	if a[i].l/blockSize != a[j].l/blockSize {
		return a[i].l < a[j].l
	}
	if ((a[i].l / blockSize) & 1) != 0 {
		return a[i].r < a[j].r
	}
	return a[i].r < a[j].r
}

func SolveMo(g *Graph, queries []Query) []int {
	ans := make([]int, len(queries))
	// blockSize := int(math.Sqrt(float64(g.N)))
	sort.Sort(Q(queries))
	i, j := 0, -1 // [i, j)
	ds := DS{make(map[int]bool), make(map[int]int), 0}
	for id, q := range queries {
		for ; i < q.l; i++ {
			ds.toggle(g.A[i], g.cost)
		}
		for i > q.l {
			i--
			ds.toggle(g.A[i], g.cost)
		}
		for j < q.r {
			j++
			ds.toggle(g.A[j], g.cost)
		}
		for ; j > q.r; j-- {
			ds.toggle(g.A[j], g.cost)
		}
		if q.lca != q.l && q.lca != q.r {
			ds.toggle(q.lca, g.cost)
		}
		ans[id] = ds.query()
		if q.lca != q.l && q.lca != q.r {
			ds.toggle(q.lca, g.cost)
		}
	}
	return ans
}

func main() {
	cin := NewReader()
	N, M := cin.Int(), cin.Int()
	graph := NewGraph(N)

	graph.cost = cin.IntArray(N)

	for i := 0; i < N-1; i++ {
		u, v := cin.Int()-1, cin.Int()-1
		graph.AddEdge(u, v)
	}

	graph.PrepareLCA(0)

	queries := make([]Query, M)

	for i := 0; i < M; i++ {
		u, v := cin.Int()-1, cin.Int()-1
		if graph.tin[u] > graph.tin[v] {
			u, v = v, u
		}
		lca := graph.LCA(u, v)
		queries[i].lca = lca
		queries[i].idx = i

		if u == lca {
			queries[i].l = graph.tin[u]
			queries[i].r = graph.tin[v]
		} else {
			queries[i].l = graph.tout[u]
			queries[i].r = graph.tin[v]
		}
	}
	ans := SolveMo(graph, queries)
	for _, it := range ans {
		fmt.Println(it)
	}
}

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

func (r *Reader) IntArray(len int) []int {
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
