package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	if k%2 != 0 {
		fmt.Fprintln(out, -1)
		return
	}
	k /= 2
	if int64(n-1) > k {
		fmt.Fprintln(out, -1)
		return
	}
	ma := int64(n/2) * int64(n-n/2)
	if ma < k {
		fmt.Fprintln(out, -1)
		return
	}
	L := make([]int, 0, n)
	R := make([]int, 0, n)
	a, b := 0, n-1
	for k > 0 {
		d := int64(b - a)
		if k >= d {
			L = append(L, a)
			R = append(R, b)
			a++
			b--
			k -= d
		} else {
			L = append(L, a)
			R = append(R, a+int(k))
			k = 0
		}
	}
	sort.Ints(L)
	sort.Ints(R)
	pos := make([]int, n)
	for i := range n {
		pos[i] = -1
	}
	idx := 0
	for i := 0; i < len(L); i++ {
		pos[L[i]] = idx
		idx++
		pos[R[i]] = idx
		idx++
	}
	for i := n - 1; i >= 0; i-- {
		if pos[i] == -1 {
			pos[i] = idx
			idx++
		}
	}
	for i := 0; i < n-1; i++ {
		fmt.Fprintf(out, "%d %d\n", pos[i]+1, pos[i+1]+1)
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

// Solution referenced from https://codeforces.com/profile/maspy

/*
This algorithm constructs a valid labeled tree forming a path utilizing an interleaved node positioning sequence to precisely
match the target sum `K` for pairwise adjacent node distances modulo traversing effectively. It first aggressively prunes cases
directly returning `-1`: ensuring `K` is even, sufficient to cover simply sequential adjacent distance pairs (`N - 1`), and
fits within the theoretical maximum attainable sum `ma` mapped systematically as paths switching alternately between smallest
and largest ends. Iteratively, bounds (`a` and `b`) collapse inwards matching and subtracting from the remaining sum gap until
completely balanced, gathering positions in two groupings `L` and `R`. After organizing nodes internally according to constraints,
indices construct a structured one-way contiguous path outputting exactly `N - 1` interconnected adjacent edges satisfying the
total distance dynamically. Time complexity bounds efficiently at $O(N \log N)$ mandated strictly by the local sorting of indices
vectors, utilizing an encompassing space footprint strictly scaled linearly to $O(N)$.
*/
