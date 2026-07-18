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
	for i := 0; i < m; i++ {
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
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		if min(deg[i][0], deg[i][1]) == 0 {
			q = append(q, i)
			vis[i] = true
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
			if !vis[e.to] && min(deg[e.to][0], deg[e.to][1]) == 0 {
				q = append(q, e.to)
				vis[e.to] = true
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
The algorithm utilizes a topological sorting mechanism akin to Kahn's algorithm to iteratively untangle the dependency graph. By
tracking the counts (degrees) of both edge types (0 and 1) connected to every vertex, it identifies nodes that are strictly
restricted to at most one edge type condition (`min(deg[i][0], deg[i][1]) == 0`). Since these nodes lack conflicting constraints,
they can safely be assigned a magnitude strictly larger than any unassigned vertex (starting from $N$ down to 1), with their sign
determined by whichever edge type they are not restricted by. As each vertex is successfully processed, it is logically removed
by decrementing the matched degree constraints of all its neighbors, potentially allowing those neighbors to be added to the queue
once their remaining constraints simplify to a single edge type. If all $N$ vertices are processed, a valid state assignment is
successfully constructed; otherwise, cyclic conflicting dependencies exist and the configuration is impossible. The time complexity
is $O(N + M)$ as every vertex and edge is processed linearly at most once, and the space complexity is $O(N + M)$ to store the
adjacency lists and the array of degree counts.
*/
