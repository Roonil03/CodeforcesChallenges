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

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func AddMod[T Integer](a, b, mod T) T {
	a %= mod
	b %= mod
	res := a + b
	if res >= mod {
		res -= mod
	}
	return res
}

func SubMod[T Integer](a, b, mod T) T {
	a %= mod
	b %= mod
	res := a - b
	if res < 0 {
		res += mod
	}
	return res
}

func MulMod[T Integer](a, b, mod T) T {
	a %= mod
	b %= mod
	return (a * b) % mod
}

func PowerMod[T Integer](base, exp, mod T) T {
	var res T = 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}

func precomp() {

}

func solve() {
	n := in.rInt()
	a := readInts(n)
	sum := 0
	for i := 0; i < n; i++ {
		if a[i] >= 0 {
			sum += a[i]
		}
	}
	if sum > n-1 {
		fmt.Fprintln(out, 0)
		return
	}
	n++
	a = append(a, -1)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	const mod int64 = 998244353
	C := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		C[i] = make([]int64, n+1)
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = AddMod(C[i-1][j], C[i-1][j-1], mod)
		}
	}
	f := make([][]int64, n)
	g := make([][]int64, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int64, n)
		g[i] = make([]int64, n)
	}
	for i := n - 1; i >= 0; i-- {
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				g[x][y] = 0
			}
		}
		g[i][0] = 1
		for j := i; j < n; j++ {
			kLimit := max(a[i]-1, 0)
			for k := 0; k <= kLimit; k++ {
				if g[j][k] == 0 {
					continue
				}
				for p := j + 1; p < n; p++ {
					nk := 0
					if a[i] != -1 {
						nk = k + 1
					}
					term := MulMod(f[j+1][p], C[p-i-1][p-j-1], mod)
					term = MulMod(g[j][k], term, mod)
					g[p][nk] = AddMod(g[p][nk], term, mod)
				}
			}
		}
		req := max(a[i], 0)
		for j := i; j < n; j++ {
			f[i][j] = g[j][req]
		}
	}
	fmt.Fprintln(out, f[0][n-1])
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
The algorithm utilizes dynamic programming to compute valid structural subsets by traversing the sequence backwards (hence
appending a sentinel -1 and reversing the array). It employs a 2D state system represented by `f` and `g`, where `f[i][j]`
maps to entirely resolved segments spanning specific ranges and `g` behaves as an accumulator for recursively building
these sequences across potential bounding splits. Sub-problems scale down by tracking dynamically assembled segments augmented
using precomputed binomial combinations (`C`), essentially integrating combinations iteratively mapped by specific target
constraints dictated by the input limits `a[i]`. The structural combinations inherently simulate branching processes valid
modulo 998244353. The time complexity per testcase is $O(N^4)$ because the core DP iterates through four nested layers
encompassing constraints indices $i$, boundaries $j$, sequence counts $k$, and bounding expansion limits $p$, all proportional
to the sequence size $N$. The space complexity is $O(N^2)$ as it solely allocates memory for managing dynamically updated dual
matrices `f` and `g`, alongside Pascal's triangle mapping structure sizes scalable geometrically up to $N$.
*/
