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
	A := in.rInt() - 1
	B := in.rInt() - 1
	tree := make([][]int, N)
	for i := 0; i < N-1; i++ {
		x := in.rInt() - 1
		y := in.rInt() - 1
		tree[x] = append(tree[x], y)
		tree[y] = append(tree[y], x)
	}
	par := make([]int, N)
	for i := range N {
		par[i] = -1
	}
	q := make([]int, 0, N)
	q = append(q, A)
	par[A] = -2
	for head := 0; head < len(q); head++ {
		u := q[head]
		if u == B {
			break
		}
		for _, v := range tree[u] {
			if par[v] == -1 {
				par[v] = u
				q = append(q, v)
			}
		}
	}
	path := make([]int, 0)
	for v := B; v != -2; v = par[v] {
		path = append(path, v)
	}
	onPath := make([]bool, N)
	for _, v := range path {
		onPath[v] = true
	}
	sizes := make([]int, len(path))
	for i, st := range path {
		cnt := 0
		q = q[:0]
		q = append(q, st)
		for head := 0; head < len(q); head++ {
			u := q[head]
			cnt++
			for _, v := range tree[u] {
				if !onPath[v] {
					onPath[v] = true
					q = append(q, v)
				}
			}
		}
		sizes[i] = cnt - 1
	}
	L := len(sizes)
	dp0 := make([][2]int, L)
	dp1 := make([][2]int, L)
	dp := [2][][2]int{dp0, dp1}
	for l := L - 1; l >= 0; l-- {
		cur := l % 2
		nxt := 1 - cur
		for r := 0; r < L; r++ {
			dp[cur][r][0] = 0
			dp[cur][r][1] = 0
		}
		for r := l + 1; r < L; r++ {
			for e := 0; e <= 1; e++ {
				if e == 1 {
					if dp[cur][r][0] == 0 {
						dp[cur][r][1] = 1
					}
				}
				if l+1 < r-1 {
					if dp[nxt][r][(e+sizes[l+1])%2] == 0 {
						dp[cur][r][e] = 1
					}
					if dp[cur][r-1][(e+sizes[r-1])%2] == 0 {
						dp[cur][r][e] = 1
					}
				}
			}
		}
	}
	fmt.Fprintln(out, dp[0][L-1][(sizes[0]+sizes[L-1])%2])
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

// Solution from https://codeforces.com/profile/ksun48

/*
This algorithm is designed to solve a 2-player game played on a tree, breaking it down into a localized path-based dynamic programming
problem. First, the algorithm finds the unique path between two specified starting nodes A and B using a Breadth-First Search. It
proceeds by isolating the components that hang off each node on this specific path (again efficiently calculated via BFS where the
path nodes act as impassable boundaries). The size of each component represents the parity of extra moves available to the players
at that particular point along the path. By interpreting the game entirely over these sizes along the path of length L, it computes
a DP table indicating if the current player has a winning strategy remaining. Since transitions only depend on shifting the active
endpoints inward (l+1 or r-1), memory is massively conserved using a rolling DP array where states map to [l%2][r][e], evaluating
the possible recursive decisions to force a losing state onto the opponent. The overall time complexity is O(N + L²) bounded essentially
by O(N²), as filling the L² DP states dominates the computation where L ≤ N, and the space complexity sits at tightly packed O(N)
memory allocation required to store the tree and the optimized dynamic programming arrays.
*/
