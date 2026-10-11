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
	u := in.rInt() - 1
	v := in.rInt() - 1
	adj := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a := in.rInt() - 1
		b := in.rInt() - 1
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	pp := make([]int, n)
	sz := make([]int, n)
	var dfs func(int, int)
	dfs = func(p, i int) {
		pp[i] = p
		sz[i] = 1
		for _, j := range adj[i] {
			if j != p {
				dfs(i, j)
				sz[i] ^= sz[j]
			}
		}
	}
	dfs(-1, u)
	if pp[v] == u {
		fmt.Fprintln(out, n%2)
		return
	}
	x := n % 2
	aa := make([]int, 0)
	curr := v
	for pp[curr] != u {
		s := sz[pp[curr]] ^ sz[curr]
		x ^= s
		aa = append(aa, s^1)
		curr = pp[curr]
	}
	get := func(o int) int {
		m := len(aa)
		if aa[o] == aa[o+1] {
			return aa[o] ^ 1
		}
		i, j := o, o+1
		for i >= 0 && j < m && aa[i] != aa[i+1] && aa[j] != aa[j-1] {
			i--
			j++
		}
		if i < 0 || j >= m || (aa[i] == aa[i+1] && aa[j] == aa[j-1]) {
			return 2
		}
		res := 0
		if aa[i] == aa[i+1] {
			res = aa[i] ^ 1
			if (o-i)%2 != 0 {
				for i >= 0 && j < m && aa[i+1] == aa[j] {
					i--
					j++
				}
				if i < 0 || j >= m {
					res = 2
				}
			}
		} else {
			res = aa[j] ^ 1
			if (o-i)%2 != 0 {
				for i >= 0 && j < m && aa[i] == aa[j-1] {
					i--
					j++
				}
				if i < 0 || j >= m {
					res = 2
				}
			}
		}
		return res
	}
	solvePath := func() int {
		m := len(aa)
		if m == 1 {
			return 0
		}
		z := 0
		for i := 0; i < m; i++ {
			z ^= aa[i]
		}
		if m%2 == 0 {
			return get(m/2-1) ^ z
		}
		aVal := get(m/2 - 1)
		bVal := get(m / 2)
		aMex := aVal ^ z
		bMex := bVal ^ z
		if aMex != 0 && bMex != 0 {
			return 0
		}
		if aMex != 1 && bMex != 1 {
			return 1
		}
		return 2
	}
	if solvePath() == x {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, 1)
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

// Solution from https://codeforces.com/profile/rainboy

/*
This algorithm evaluates a 2-player competitive game defined on a tree by converting the sequence of valid game movements along
a specific path into a one-dimensional array of XOR constraints (resembling Nim). It uses Depth-First Search from the target node
$U$ to build parent relationships and recursively calculate the parity of subtree sizes, effectively reducing the entire tree to
sequential path options from the starting node $V$ towards $U$. If the initial starting vertex $V$ is directly adjacent to $U$,
the game concludes in a single check. Otherwise, it extracts the remaining path into a boolean-parity array `aa` where consecutive
XOR differences act as game-theory configurations mapping directly to Grundy numbers. It optimally exploits symmetry avoiding $O(N^2)$
dynamic programming: an even length utilizes checking outward patterns from the exact array midpoint, while an odd length uses `mex`
applied to the two central components coupled with the overall array's XOR sum to calculate optimal responses and determine whether
the starting player possesses a winning strategy. The total Time Complexity is $O(N)$ due to linear tree traversal and symmetric
two-pointer midpoint evaluations, and the Space Complexity is bounded by $O(N)$ for building tree adjacency relations and recursive
call stacks per testcase.
*/
