package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = NewFastScanner()
	out = bufio.NewWriter(os.Stdout)
)

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (fs *FastScanner) rInt() int {
	sign, val := 1, 0
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		if err != nil {
			return 0
		}
		c, err = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return sign * val
}

func (fs *FastScanner) rInt64() int64 {
	sign, val := int64(1), int64(0)
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		if err != nil {
			return 0
		}
		c, err = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int64(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return sign * val
}

func (fs *FastScanner) rString() string {
	buf := make([]byte, 0, 16)
	c, err := fs.r.ReadByte()
	for c == ' ' || c == '\n' || c == '\r' || c == '\t' {
		if err != nil {
			return ""
		}
		c, err = fs.r.ReadByte()
	}
	for c != ' ' && c != '\n' && c != '\r' && c != '\t' {
		buf = append(buf, c)
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return string(buf)
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxf32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func maxf64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minf32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func minf64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func absf32(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func absf64(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func readInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = in.rInt()
	}
	return a
}

func readInt64s(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = in.rInt64()
	}
	return a
}

func printSlice(arr []int) {
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func precomp() {

}

func solve() {
	mod := 998244353
	powMod := func(a, b int) int {
		res := 1
		a %= mod
		for b > 0 {
			if b%2 == 1 {
				res = (res * a) % mod
			}
			a = (a * a) % mod
			b /= 2
		}
		return res
	}
	dominatorTree := func(n int, edges [][]int, root int) []int {
		g := make([][]int, n)
		rg := make([][]int, n)
		for _, e := range edges {
			g[e[0]] = append(g[e[0]], e[1])
			rg[e[1]] = append(rg[e[1]], e[0])
		}
		par := make([]int, n)
		ord := make([]int, n)
		for i := 0; i < n; i++ {
			ord[i] = -1
		}
		vertex := make([]int, 0, n)
		var dfs func(int)
		dfs = func(v int) {
			ord[v] = len(vertex)
			vertex = append(vertex, v)
			for _, to := range g[v] {
				if ord[to] == -1 {
					par[to] = v
					dfs(to)
				}
			}
		}
		dfs(root)
		sdom := make([]int, n)
		idom := make([]int, n)
		for i := range n {
			sdom[i] = i
			idom[i] = -1
		}
		if len(vertex) == 0 {
			return idom
		}
		bucket := make([][]int, n)
		anc := make([]int, n)
		mi := make([]int, n)
		for i := range n {
			anc[i] = -1
			mi[i] = i
		}
		var find func(int) int
		find = func(v int) int {
			if anc[v] == -1 {
				return v
			}
			b := find(anc[v])
			if ord[sdom[mi[anc[v]]]] < ord[sdom[mi[v]]] {
				mi[v] = mi[anc[v]]
			}
			anc[v] = b
			return b
		}
		us := make([]int, n)
		for i := len(vertex) - 1; i >= 1; i-- {
			w := vertex[i]
			for _, v := range rg[w] {
				if ord[v] == -1 {
					continue
				}
				find(v)
				if ord[sdom[mi[v]]] < ord[sdom[w]] {
					sdom[w] = sdom[mi[v]]
				}
			}
			bucket[sdom[w]] = append(bucket[sdom[w]], w)
			p := par[w]
			for _, v := range bucket[p] {
				find(v)
				us[v] = mi[v]
			}
			bucket[p] = bucket[p][:0]
			anc[w] = p
		}
		for i := 1; i < len(vertex); i++ {
			w := vertex[i]
			if sdom[w] == sdom[us[w]] {
				idom[w] = sdom[w]
			} else {
				idom[w] = idom[us[w]]
			}
		}
		return idom
	}
	H := in.rInt()
	W := in.rInt()
	G := make([][]byte, H)
	for i := range H {
		G[i] = []byte(in.rString())
	}
	que := make([][2]int, 0, H*W)
	for x := range H {
		for y := range W {
			que = append(que, [2]int{x, y})
		}
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	for head := 0; head < len(que); head++ {
		x, y := que[head][0], que[head][1]
		if x < 0 || x >= H || y < 0 || y >= W {
			continue
		}
		if G[x][y] == '0' {
			continue
		}
		u := 0
		if x == 0 || G[x-1][y] == '0' {
			u = 1
		}
		l := 0
		if y == 0 || G[x][y-1] == '0' {
			l = 1
		}
		d := 0
		if x == H-1 || G[x+1][y] == '0' {
			d = 1
		}
		r := 0
		if y == W-1 || G[x][y+1] == '0' {
			r = 1
		}
		if (x+y != 0 && u+l == 2) || (x+y != H+W-2 && r+d == 2) {
			G[x][y] = '0'
			for d2 := range 4 {
				que = append(que, [2]int{x + dx[d2], y + dy[d2]})
			}
		}
	}
	bad := 0
	for i := range H {
		for j := range W {
			if G[i][j] == '0' {
				bad++
			}
		}
	}
	if G[0][0] == '0' {
		ans := (powMod(2, bad) - 1 + mod) % mod
		fmt.Fprintln(out, ans)
		return
	}
	idx := make([][]int, H)
	n := 0
	for i := range H {
		idx[i] = make([]int, W)
		for j := 0; j < W; j++ {
			if G[i][j] == '0' {
				idx[i][j] = -1
			} else {
				idx[i][j] = n
				n++
			}
		}
	}
	edges1 := make([][]int, 0)
	edges2 := make([][]int, 0)
	for x := range H {
		for y := range W {
			if idx[x][y] == -1 {
				continue
			}
			for d2 := range 2 {
				xx := x + dx[d2]
				yy := y + dy[d2]
				if xx >= 0 && xx < H && yy >= 0 && yy < W {
					s := idx[x][y]
					t := idx[xx][yy]
					if s != -1 && t != -1 {
						edges1 = append(edges1, []int{s, t})
						edges2 = append(edges2, []int{t, s})
					}
				}
			}
		}
	}
	par1 := dominatorTree(n, edges1, idx[0][0])
	par2 := dominatorTree(n, edges2, idx[H-1][W-1])
	T := make([][]int, n)
	for i := 1; i < n; i++ {
		p := par1[i]
		if p != -1 {
			T[p] = append(T[p], i)
		}
	}
	lid := make([]int, n)
	rid := make([]int, n)
	for i := 0; i < n; i++ {
		lid[i] = -1
		rid[i] = -1
	}
	timer := 0
	var dfsT func(int)
	dfsT = func(v int) {
		lid[v] = timer
		timer++
		for _, to := range T[v] {
			dfsT(to)
		}
		rid[v] = timer
	}
	dfsT(idx[0][0])
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	var dsuFind func(int) int
	dsuFind = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = dsuFind(parent[i])
		return parent[i]
	}
	merge := func(i, j int) {
		rootI := dsuFind(i)
		rootJ := dsuFind(j)
		if rootI != rootJ {
			if size[rootI] < size[rootJ] {
				rootI, rootJ = rootJ, rootI
			}
			parent[rootJ] = rootI
			size[rootI] += size[rootJ]
		}
	}
	for i := 0; i < n; i++ {
		j := par2[i]
		if j == -1 {
			continue
		}
		if lid[i] != -1 && lid[j] != -1 && lid[i] <= lid[j] && lid[j] < rid[i] {
			merge(i, j)
		}
	}
	ans := (powMod(2, bad) - 1 + mod) % mod
	for r := 0; r < n; r++ {
		if parent[r] == r {
			ans = (ans + powMod(2, size[r]) - 1 + mod) % mod
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	precomp()
	t := in.rInt()
	for t > 0 {
		t--
		solve()
	}
}

// Solution referenced from https://codeforces.com/profile/maspy

/*
The algorithm identifies redundant grid paths by recursively stripping away dead ends using a queue, where any cell (excluding the
absolute start or end) unable to receive paths from the top/left or emit paths to the right/down is marked as blocked. After
computing the number of forced blocked cells, it translates the remaining open cells into a directed acyclic graph and its reverse,
computing the dominator tree for both from the start and end points respectively utilizing the Lengauer-Tarjan algorithm. Subsets
of valid paths form disjoint sets determined by iterating through nodes and merging those where a node's reverse dominator is an
ancestor of the node in the forward dominator tree, meaning they share an inescapable bottleneck sequence. The valid subgrid
configurations answer is formulated combinatorially by accumulating 2^(size) - 1 mapped over each disjoint component size,
outputted modulo 998244353. The time complexity is bounded by O(HW \log(HW)) due to the path compression traversing in the
dominator tree construction, and the space complexity necessitates O(HW) to robustly store the directed graphs, queues, and
disjoint set structures scaled to the grid dimensions.
*/
