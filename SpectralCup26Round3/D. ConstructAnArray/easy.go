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
	n := in.rInt()
	m := in.rInt()
	type edge struct {
		to, o int
	}
	g := make([][]edge, n)
	deg := make([][2]int, n)
	for range m {
		o := in.rInt() - 1
		x := in.rInt() - 1
		y := in.rInt() - 1
		g[x] = append(g[x], edge{y, o})
		g[y] = append(g[y], edge{x, o})
		deg[x][o]++
		deg[y][o]++
	}
	a := make([]int, n)
	q := make([]int, 0, n)
	used := make([]bool, n)
	for i := range n {
		if min(deg[i][0], deg[i][1]) == 0 {
			q = append(q, i)
			used[i] = true
		}
	}
	for it := 0; it < len(q); it++ {
		i := q[it]
		if deg[i][1] == 0 {
			a[i] = n - it
		} else {
			a[i] = -(n - it)
		}
		for _, e := range g[i] {
			deg[e.to][e.o]--
			if !used[e.to] && min(deg[e.to][0], deg[e.to][1]) == 0 {
				q = append(q, e.to)
				used[e.to] = true
			}
		}
	}
	if len(q) == n {
		fmt.Fprintln(out, "YES")
		printSlice(a)
	} else {
		fmt.Fprintln(out, "NO")
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
The algorithm functions similarly to Kahn's algorithm for topological sorting to peel off valid vertices iteratively from the
graph. By tracking the degree of both edge types (0 and 1) for every vertex, it identifies nodes that are strictly restricted
to at most one edge type (i.e., `min(deg[i][0], deg[i][1]) == 0`). Since these nodes lack conflicting constraints, they can be
safely assigned an absolute value strictly larger than any unassigned vertex ($N$, then $N-1$, and so forth), determining the
positive or negative sign depending on which edge type it lacks. Once a vertex is processed, we conceptually remove it by
decrementing the matched degree constraints of its neighbors, potentially allowing those neighbors to be added to the queue if
their remaining constraints simplify to a single edge type. If all $N$ vertices are processed, a valid assignment is constructed
and verified; otherwise, unresolved cycles of conflicting edge types exist, making it impossible. The time complexity is
$O(N + M)$ as every vertex and edge is processed at most once, and the space complexity is $O(N + M)$ for maintaining the
adjacency list and degree counts.
*/
