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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	a -= b
	if a < 0 {
		a += mod
	}
	return a
}

func mul(a, b int) int {
	return int((int64(a) * int64(b)) % mod)
}

func solve() {
	n := in.rInt()
	m := in.rInt()
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			a[i][j] = in.rInt()
			if a[i][j] < 0 {
				a[i][j] += mod
			}
		}
	}
	res := 0
	solveFunc := func(a [][]int) {
		pre := make([][]int, n)
		for i := range n {
			pre[i] = make([]int, m)
			for j := 0; j < m; j++ {
				if i%2 == n%2 || j%2 == m%2 {
					pre[i][j] = 0
				} else {
					pre[i][j] = a[i][j]
				}
			}
		}
		for i := n - 1; i >= 0; i-- {
			for j := m - 2; j >= 0; j-- {
				pre[i][j] = add(pre[i][j], pre[i][j+1])
			}
		}
		for i := n - 2; i >= 0; i-- {
			for j := m - 1; j >= 0; j-- {
				pre[i][j] = add(pre[i][j], pre[i+1][j])
			}
		}
		for i := range n {
			for j := range m {
				now := pre[i][j]
				if i%2 != n%2 && j%2 != m%2 {
					now = sub(now, a[i][j])
				}
				if i%2 == 0 && j%2 == 0 {
					res = add(res, mul(a[i][j], mul(2, now)))
				}
			}
		}
	}
	solveFunc(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
	solveFunc(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
	for i := range n {
		if i%2 == 0 && i%2 != n%2 {
			sum := 0
			for j := m - 1; j >= 0; j-- {
				if j%2 == 0 {
					res = sub(res, mul(a[i][j], mul(2, sum)))
				}
				if j%2 != m%2 {
					sum = add(sum, a[i][j])
				}
			}
		}
	}
	for j := range m {
		if j%2 == 0 && j%2 != m%2 {
			sum := 0
			for i := n - 1; i >= 0; i-- {
				if i%2 == 0 {
					res = sub(res, mul(a[i][j], mul(2, sum)))
				}
				if i%2 != n%2 {
					sum = add(sum, a[i][j])
				}
			}
		}
	}
	for i := range n {
		for j := range m {
			if i%2 == 0 && i%2 != n%2 && j%2 == 0 && j%2 != m%2 {
				res = add(res, mul(a[i][j], a[i][j]))
			}
		}
	}
	if (n+m)&1 == 1 {
		res = sub(0, res)
	}
	fmt.Fprintln(out, res)
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

// Solution from https://codeforces.com/profile/abc864197532

/*
This algorithm evaluates a 2D weighted sum by isolating distinct alternating grid patterns. It
accumulates terms derived from a matrix representation by processing block sums through two core
`solveFunc` invocations—first acting directly on the grid and then effectively reversing it
vertically. The logic efficiently precomputes reverse prefix 2D sums, masked dynamically depending
on matrix dimension parities (`n%2`, `m%2`). It then iterates to sum cross-multiplied pairs
(`a[i][j]` against the filtered block sum `pre[i][j]`) for specifically indexed cells. The later
loops fine-tune these contributions by subtracting over-counted row/col subsets and explicitly
adding squared diagonal terms. Because all 2D prefix aggregations operate in place over isolated
submatrices, the time complexity is tightly O(N * M) per testcase. The overall space complexity
sits at O(N * M) designated for allocating and modifying the matrix grid structure alongside the
identically sized prefix-sum array.
*/
