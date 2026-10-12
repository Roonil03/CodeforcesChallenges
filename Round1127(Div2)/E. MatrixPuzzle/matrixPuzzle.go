package main

import (
	"bufio"
	"cmp"
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

type Fenwick2D[T cmp.Ordered] struct {
	n, m int
	tree [][]T
	add  func(T, T) T
}

func NewFenwick2D[T cmp.Ordered](n, m int, add func(T, T) T) *Fenwick2D[T] {
	b := make([][]T, n+1)
	for i := range b {
		b[i] = make([]T, m+1)
	}
	return &Fenwick2D[T]{n: n, m: m, tree: b, add: add}
}

func (f *Fenwick2D[T]) Update(x, y int, v T) {
	for i := x; i <= f.n; i += i & -i {
		for j := y; j <= f.m; j += j & -j {
			f.tree[i][j] = f.add(f.tree[i][j], v)
		}
	}
}

func (f *Fenwick2D[T]) Query(x, y int, zero T) T {
	sum := zero
	for i := x; i > 0; i -= i & -i {
		for j := y; j > 0; j -= j & -j {
			sum = f.add(sum, f.tree[i][j])
		}
	}
	return sum
}

func solve() {
	n := in.rInt()
	// arr := readInts(n)
	q := in.rInt()
	a := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			a[i][j] = in.rInt()
		}
	}
	b := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		b[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			b[i][j] = in.rInt()
		}
	}
	d := make([][]int64, n+1)
	for i := 1; i <= n; i++ {
		d[i] = make([]int64, n+1)
	}
	bit := NewFenwick2D(n, n, func(x, y int) int {
		return x + y
	})
	for i := 2; i <= n; i++ {
		for j := 2; j <= n; j++ {
			d[i][j] = int64(b[i][j]-a[i][j]) - int64(b[i-1][j]-a[i-1][j]) - int64(b[i][j-1]-a[i][j-1]) + int64(b[i-1][j-1]-a[i-1][j-1])
			if d[i][j] != 0 {
				bit.Update(i, j, 1)
			}
		}
	}
	add := func(r, c int, val int64) {
		if r >= 2 && r <= n && c >= 2 && c <= n {
			if d[r][c] == 0 && d[r][c]+val != 0 {
				bit.Update(r, c, 1)
			} else if d[r][c] != 0 && d[r][c]+val == 0 {
				bit.Update(r, c, -1)
			}
			d[r][c] += val
		}
	}
	for range q {
		if in.rString() == "U" {
			r1, c1, r2, c2, x := in.rInt(), in.rInt(), in.rInt(), in.rInt(), in.rInt64()
			add(r1, c1, -x)
			add(r1, c2+1, x)
			add(r2+1, c1, x)
			add(r2+1, c2+1, -x)
		} else {
			r1, c1, r2, c2 := in.rInt(), in.rInt(), in.rInt(), in.rInt()
			if r1 == r2 || c1 == c2 {
				fmt.Fprintln(out, "yes")
			} else {
				sum := bit.Query(r2, c2, 0) - bit.Query(r1, c2, 0) - bit.Query(r2, c1, 0) + bit.Query(r1, c1, 0)
				if sum == 0 {
					fmt.Fprintln(out, "yes")
				} else {
					fmt.Fprintln(out, "no")
				}
			}
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
