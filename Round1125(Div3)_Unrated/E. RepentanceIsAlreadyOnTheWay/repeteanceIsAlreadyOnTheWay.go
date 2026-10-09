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

func precomp() {}

func solve() {
	n := in.rInt()
	a := readInt64s(n)
	b := readInt64s(n)
	C := func(x, y int64) int64 {
		if x == y {
			return 2
		}
		return 1
	}
	dp := make([][2][2]int64, n)
	for i := range 2 {
		for j := range 2 {
			dp[n-1][i][j] = C(a[n-1], b[n-1])
		}
	}
	for i := n - 2; i >= 0; i-- {
		val1 := C(a[i], b[i]) + C(b[i], a[i+1]) + dp[i+1][0][0]
		val2 := C(a[i], b[i+1]) + dp[i+1][1][1] + C(a[i+1], b[i])
		dp[i][0][0] = max64(val1, val2)
		dp[i][0][1] = C(a[i], b[i+1]) + dp[i+1][1][1] + C(a[i+1], b[i])
		val3 := C(b[i], a[i]) + C(a[i], b[i+1]) + dp[i+1][1][0]
		val4 := C(b[i], a[i+1]) + dp[i+1][0][1] + C(b[i+1], a[i])
		dp[i][1][0] = max64(val3, val4)
		dp[i][1][1] = C(b[i], a[i+1]) + dp[i+1][0][1] + C(b[i+1], a[i])
	}
	fmt.Fprintln(out, dp[0][0][0])
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

// Solution by https://codeforces.com/profile/maspy

/*
The algorithm utilizes dynamic programming to determine the maximum score achievable by traversing
two parallel sequences of values, scoring 1 point for any adjacent transition and 2 points if the
adjacent transition shares identical values. It maintains states `dp[i][entry_state][exit_state]`
defining sequences over the suffix starting at index $i$, where the state flags differentiate
between straight segments directly returning along the same parallel line (`0,0` / `1,0` etc) and
diagonal cross-segment configurations. By transitioning backward iteratively from $N-2$ down to
$0$, it optimally selects between linking consecutive items vertically versus making optimal
crisscross horizontal exchanges to securely accumulate points evaluated by lambda `C(x, y)`. The
time complexity translates strictly to $O(N)$ derived from evaluating fixed contiguous state
transitions iteratively over $N$ items, while the space complexity utilizes $O(N)$ for capturing
and memoizing the array inputs and tracking configurations within the dimensioned dynamic
programming table.
*/
