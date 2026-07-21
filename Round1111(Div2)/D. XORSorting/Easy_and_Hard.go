package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

var mx, mn, mbh []int

func Init() {
	sz := 1 << 21
	if len(mx) < sz {
		mx = make([]int, sz)
		mn = make([]int, sz)
		mbh = make([]int, sz)
	}
}

func precomp() {
	Init()
}

func solve() {
	n := in.rInt()
	q := in.rInt()
	// arr := readInt64s(n)
	s, k := 1, 0
	for s < n {
		s <<= 1
		k++
	}
	for i := range n {
		val := in.rInt()
		id := s + i
		mx[id] = val
		mn[id] = val
		mbh[id] = -1
	}
	for i := n; i < s; i++ {
		val := int(1e9+1) + i
		id := s + i
		mx[id] = val
		mn[id] = val
		mbh[id] = -1
	}
	for i := s - 1; i > 0; i-- {
		l := i << 1
		r := l | 1
		m1 := max(mx[l], mx[r])
		mx[i] = m1
		m2 := min(mn[l], mn[r])
		mn[i] = m2
		bh := max(mbh[l], mbh[r])
		if mx[l] > mn[r] {
			h := k - bits.Len(uint(i))
			if h > bh {
				bh = h
			}
		}
		mbh[i] = bh
	}
	if mbh[1] == -1 {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, 1<<mbh[1])
	}
	for range q {
		id := in.rInt()
		val := in.rInt()
		j := s + id
		mx[j] = val
		mn[j] = val
		j >>= 1
		for j > 0 {
			l := j << 1
			r := l | 1
			m1 := max(mx[l], mx[r])
			mx[j] = m1
			m2 := min(mn[l], mn[r])
			mn[j] = m2
			bh := max(mbh[l], mbh[r])
			if mx[l] > mn[r] {
				h := k - bits.Len(uint(j))
				if h > bh {
					bh = h
				}
			}
			mbh[j] = bh
			j >>= 1
		}
		if mbh[1] == -1 {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, 1<<mbh[1])
		}
	}
	// printSlice(arr)
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
