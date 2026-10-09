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
	k := in.rInt64()
	arr := make([][3]int64, n)
	for i := range n {
		for j := range 3 {
			arr[i][j] = in.rInt64()
		}
	}
	ask := func(x int64) bool {
		var ans int64 = 0
		for i := range n {
			a := arr[i][0]
			b := arr[i][1]
			c := arr[i][2]
			if a+b+c >= x {
				continue
			}
			if a > b || b > c || a > c {
				ans += x - (a + b + c)
			} else {
				var need int64 = 1 << 60
				if a != c {
					need = min64(need, b-a+1)
				}
				if a != b {
					need = min64(need, c-b+1)
				}
				ans += 2*need + x - (a + b + c)
			}
			if ans > k {
				return false
			}
		}
		return true
	}
	var l int64 = -1 << 60
	var r int64 = 1 << 61
	for r-l > 1 {
		mid := l + (r-l)/2
		if ask(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Fprintln(out, l)
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

// Solution by https://codeforces.com/profile/abc864197532

/*
The algorithm utilizes binary search to efficiently find the maximum possible minimum value $x$
for the sum of three independently updated variables $(a_i, b_i, c_i)$ across $N$ tuples,
utilizing at most $K$ total operations. The core logic centers on an `ask(x)` function that
determines if it is possible to raise the sum of every tuple to at least $x$ within the allowed
cost limit $K$. For each tuple, if the initial sum $a + b + c$ is already $\ge x$, no cost is
incurred. If a tuple is already asymmetric ($a > b \lor b > c \lor a > c$), it can be freely
incremented without restriction, requiring exactly $x - (a + b + c)$ operations. However, if
the tuple is ordered symmetrically $a \le b \le c$, it cannot be freely incremented. It requires
a specific minimal number of operations `need` to break the symmetry and safely cross the
equality boundaries before it can be freely incremented. Breaking this symmetry requires $2
\times need$ operations, after which the remaining required sum is added. The total calculated
cost is accumulated and validated against $K$. The binary search encompasses extreme bounds to
handle potentially massive target values ranging up to $10^{18}$. The time complexity is
$O(N \log(\text{MAX\_X}))$ per testcase, where $\text{MAX\_X}$ is the search space bound
($\approx 2^{61}$), and the space complexity is $O(N)$ for allocating the tuples.
*/
