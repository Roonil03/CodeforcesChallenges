package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p, r := make([]int, n), make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	x, y = d.Find(x), d.Find(y)
	if x == y {
		return
	}
	if d.rank[x] < d.rank[y] {
		d.parent[x] = y
	} else {
		d.parent[y] = x
		if d.rank[x] == d.rank[y] {
			d.rank[x]++
		}
	}
}

func precomp() {

}

func solve() {
	n := in.rInt()
	Q := 0
	Ask := func(seq []int) string {
		Q += len(seq)
		fmt.Fprintf(out, "? %d", len(seq))
		for _, v := range seq {
			fmt.Fprintf(out, " %d", v+1)
		}
		fmt.Fprintln(out)
		out.Flush()
		return in.rString()
	}

	order := make([]int, n)
	for i := range n {
		order[i] = i
	}
	known := make([][]int, n)
	mat := make([][]int, n)
	for i := range n {
		known[i] = make([]int, n)
		mat[i] = make([]int, n)
		for j := range n {
			mat[i][j] = 1
		}
		mat[i][i] = 0
	}
	d := NewDSU(n)
	for it := 0; ; it++ {
		g := make([][]int, n)
		cc := 0
		ck := 0
		for i := range n {
			for j := i + 1; j < n; j++ {
				if d.Find(i) == d.Find(j) && known[i][j] == 0 && mat[i][j] == 1 {
					mat[i][j] = 0
					mat[j][i] = 0
				}
				cc += mat[i][j]
				ck += known[i][j]
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
		if Q+2*(cc-ck) <= 30*n {
			for i := range n {
				for j := i + 1; j < n; j++ {
					if mat[i][j] == 1 && known[i][j] == 0 {
						got := Ask([]int{i, j})
						if got == "11" {
							mat[i][j] = 0
							mat[j][i] = 0
						} else {
							known[i][j] = 1
							known[j][i] = 1
						}
					}
				}
			}
			fmt.Fprintln(out, "!")
			for i := range n {
				for j := i + 1; j < n; j++ {
					if mat[i][j] == 1 {
						fmt.Fprintf(out, "%d %d\n", i+1, j+1)
					}
				}
			}
			out.Flush()
			break
		}
		rand.Shuffle(len(order), func(i, j int) {
			order[i], order[j] = order[j], order[i]
		})
		st := rand.Intn(n)
		was := make([]bool, n)
		que := make([]int, 0, n)
		var Dfs func(int)
		Dfs = func(v int) {
			que = append(que, v)
			was[v] = true
			rand.Shuffle(len(g[v]), func(i, j int) {
				g[v][i], g[v][j] = g[v][j], g[v][i]
			})
			for _, j := range g[v] {
				if !was[j] && mat[v][j] == 1 {
					Dfs(j)
				}
			}
		}
		Dfs(st)
		order = que
		m := n
		got := Ask(order[:m])
		for i := range m {
			if got[i] == '1' {
				for j := range m {
					if got[j] == '1' {
						mat[order[i]][order[j]] = 0
					}
				}
			} else {
				cand := make([]int, 0)
				for j := range i {
					if got[j] == '1' {
						if mat[order[i]][order[j]] == 1 {
							cand = append(cand, order[j])
						}
					}
				}
				if len(cand) == 1 {
					known[order[i]][cand[0]] = 1
					known[cand[0]][order[i]] = 1
					d.Union(order[i], cand[0])
				}
			}
		}
	}
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

// Solution referenced from https://codeforces.com/profile/tourist

/*
The algorithm probabilistically discovers a hidden tree using an interactive subset querying method to trim down possible remaining
edges over successive permutations. It tracks candidates via a boolean adjacency matrix `mat` initialized as a fully connected
complete graph, alongside a DSU tracking confidently known edges (`known`). Iteratively, the algorithm randomly traverses the
potential candidates and queries these full DFS-generated node orderings iteratively. The results inherently allow distinction
amongst connected components: queried subsets marked '1' are internally dense and those components sharing edges can be safely
severed if contradicting spanning tree characteristics, whereas elements returning '0' but definitively pointing at exactly one
valid preceding '1' unmask definite tree edges. Any cycle detected inside via the DSU union tracking gets instantly purged. As
the candidate margin aggressively trims, once the number of required unresolved pairs fits strictly into the remainder of the
permitted bounding cap (e.g. `Q + 2(cc - ck) <= 30 * N`), the exact subset queries process the remaining edges manually to output
the definitive tree sequence. The expected time complexity operates at roughly $O(N^2 \log N)$ heavily regulated by bounding
constraints on query volume pruning, while the space complexity is strictly $O(N^2)$ due to maintaining symmetric interaction
matrices indicating candidacies and definitively identified edges.
*/
