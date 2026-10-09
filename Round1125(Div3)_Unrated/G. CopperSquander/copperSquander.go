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
	N := in.rInt()
	M := in.rInt()
	Q := in.rInt()
	type edge struct {
		u, v int
		w    int64
	}
	edges := make([]edge, M)
	var base int64
	for i := range M {
		edges[i].u = in.rInt() - 1
		edges[i].v = in.rInt() - 1
		edges[i].w = in.rInt64()
		base += edges[i].w
	}
	var qsort func(l, r int)
	qsort = func(l, r int) {
		if l >= r {
			return
		}
		p := edges[l+(r-l)/2].w
		i, j := l, r
		for i <= j {
			for edges[i].w < p {
				i++
			}
			for edges[j].w > p {
				j--
			}
			if i <= j {
				edges[i], edges[j] = edges[j], edges[i]
				i++
				j--
			}
		}
		qsort(l, j)
		qsort(i, r)
	}
	qsort(0, M-1)
	parent := make([]int, N)
	for i := range N {
		parent[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}
	union := func(i, j int) bool {
		rI, rJ := find(i), find(j)
		if rI != rJ {
			parent[rI] = rJ
			return true
		}
		return false
	}
	A := make([]int64, 0, N)
	for _, e := range edges {
		if union(e.u, e.v) {
			A = append(A, e.w)
		}
	}
	sz := len(A)
	pref := make([]int64, sz+1)
	for i := 0; i < sz; i++ {
		pref[i+1] = pref[i] + A[i]
	}
	F := func(n int, X int64) int64 {
		rem := int64(N - 1 - n)
		return pref[n] + rem*(rem+1)/2*X
	}
	for i := range Q {
		X := in.rInt64()
		low, high := 0, sz
		for low < high {
			mid := low + (high-low)/2
			if F(mid, X) <= F(mid+1, X) {
				high = mid
			} else {
				low = mid + 1
			}
		}
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, base-F(low, X))
	}
	fmt.Fprintln(out)
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

// Solution from https://codeforces.com/profile/maspy

/*
This algorithm addresses a minimum spanning forest modification problem by leveraging Kruskal's
algorithm and binary search over a convex cost function. Initially, the algorithm reads all edges,
accumulates their total weight (the "base" cost), and sorts them to form a Minimum Spanning Tree
(MST) using a Union-Find disjoint set. The weights of the edges successfully added to the MST are
stored sequentially, naturally in ascending order, and a prefix sum array is built from these
weights. For each query with a penalty value X, the algorithm simulates forming a forest by
selecting the first 'n' cheapest MST edges and complementing the remaining unconnected (N - n)
components using hypothetical edges of weight X. The cost of adding these structural edges equates
to the number of missing edges `(N - 1 - n) * (N - n) / 2` scaled by X. Because the accumulated MST
edges are strictly increasing, this simulated function `F(n)` forms a unimodal/convex sequence.
Thus, a binary search iteratively finds the discrete minimal point `min_cost` in O(log N) time. The
final maximum achievable value for the query is retrieved via `base - min_cost`.

Time Complexity: O(M log M + Q log N), where M is the number of edges (for sorting) and Q is the
number of queries evaluating binary search across N possible edges.

Space Complexity: O(M + N), as the edge lists, Union-Find arrays, and the MST prefix sum structures
require memory proportional to the size of the graph nodes and edges.
*/
