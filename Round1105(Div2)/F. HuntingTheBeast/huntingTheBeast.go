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

const mod = 998244353

var (
	fac     []int
	ifac    []int
	queries [][2]int
)

func qpow(a, b int) int {
	a %= mod
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
	}
	return res
}

func C(n, m int) int {
	if m < 0 || m > n {
		return 0
	}
	return fac[n] * ifac[m] % mod * ifac[n-m] % mod
}

func precomp() {
}

func solve() {
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()

	t := in.rInt()
	queries = make([][2]int, t)
	maxn := 0
	for i := 0; i < t; i++ {
		queries[i][0] = in.rInt()
		queries[i][1] = in.rInt()
		if queries[i][0] > maxn {
			maxn = queries[i][0]
		}
	}

	fac = make([]int, maxn+1)
	ifac = make([]int, maxn+1)
	fac[0] = 1
	ifac[0] = 1
	for i := 1; i <= maxn; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	ifac[maxn] = qpow(fac[maxn], mod-2)
	for i := maxn; i >= 1; i-- {
		ifac[i-1] = ifac[i] * i % mod
	}

	for i := 0; i < t; i++ {
		n, m := queries[i][0], queries[i][1]
		k := n - m
		sum := 0
		for j := 0; j <= k; j++ {
			w := C(k, j) * qpow(m+j-1, n-1) % mod
			if (k-j)&1 == 1 {
				sum = (sum + mod - w) % mod
			} else {
				sum = (sum + w) % mod
			}
		}
		ans := C(n, m) * sum % mod
		fmt.Fprintln(out, ans*(n-1)%mod)
	}
}

/*
The algorithm uses the Principle of Inclusion-Exclusion (PIE) alongside combinatorics under modulo 998244353 to answer multiple
offline queries. It precomputes factorials and inverse factorials up to the maximum required N across all queries to support
efficient combinations $C(n, m)$ computations in $O(1)$ time. For each query $(n, m)$, it computes an inclusion-exclusion
alternating sum spanning from $0$ to $k = n-m$. The term calculates permutations or combinations scaled by a modular
exponentiation power $qpow(m+j-1, n-1)$.

Time Complexity: $O(\text{maxn} + T \cdot N \log N)$ where $\text{maxn}$ is the maximum $N$ over all queries. The linear
precomputation requires $O(\text{maxn} + \log \text{mod})$ steps, while processing each query takes $O(k \log N) \approx
O(N \log N)$ operations because of the loop and the nested modular power logic.

Space Complexity: $O(\text{maxn} + T)$ to house the factorials, inverse factorials, and the collection of offline
test cases.
*/
