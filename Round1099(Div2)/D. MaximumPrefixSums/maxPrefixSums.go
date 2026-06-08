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
	s := in.rString()
	a := readInt64s(n)
	c := readInt64s(n)
	for i := 1; i < n; i++ {
		if c[i] < c[i-1] {
			fmt.Fprintln(out, "no")
			return
		}
	}
	l := make([]int64, n)
	r := make([]int64, n)
	const inf = int64(2e18)
	for i := range n {
		r[i] = c[i]
		if i == 0 || c[i] > c[i-1] {
			l[i] = c[i]
		} else {
			l[i] = -inf
		}
	}
	if s[0] == '1' {
		if a[0] > l[0] {
			l[0] = a[0]
		}
		if a[0] < r[0] {
			r[0] = a[0]
		}
	}
	for i := n - 1; i >= 1; i-- {
		if s[i] == '1' {
			reql := int64(-inf)
			if l[i] != -inf {
				reql = l[i] - a[i]
			}
			reqr := r[i] - a[i]
			if reql > l[i-1] {
				l[i-1] = reql
			}
			if reqr < r[i-1] {
				r[i-1] = reqr
			}
		}
	}
	for i := range n {
		if l[i] > r[i] {
			fmt.Fprintln(out, "no")
			return
		}
	}
	fmt.Fprintln(out, "yes")
	b := make([]int64, n)
	b[0] = r[0]
	a[0] = b[0]
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			b[i] = b[i-1] + a[i]
		} else {
			b[i] = r[i]
			a[i] = b[i] - b[i-1]
		}
	}
	for i := range n {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, a[i])
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
