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