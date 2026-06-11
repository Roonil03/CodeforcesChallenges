package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = NewFastScanner()
	out = bufio.NewWriter(os.Stdout)

	vis []int
	cd  int
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

func h1(arr []int) bool {
	cd++
	if len(arr) == 0 {
		return true
	}
	vis[arr[0]] = cd
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			if vis[arr[i]] == cd {
				return false
			}
			vis[arr[i]] = cd
		}
	}
	return true
}

func solve() {
	n := in.rInt()
	a := readInts(n)
	m1 := make(map[int]int)
	id := 0
	for i := range n {
		if _, ok := m1[a[i]]; !ok {
			m1[a[i]] = id
			id++
		}
		a[i] = m1[a[i]]
	}
	if len(vis) < id {
		vis = make([]int, max(len(vis)*2, id))
	}
	if h1(a) {
		fmt.Fprintln(out, "YES")
		return
	}
	x := make([]int, id)
	y := make([]int, id)
	for i := range id {
		x[i] = -1
		y[i] = -1
	}
	for i, v := range a {
		if x[v] == -1 {
			x[v] = i
		}
		y[v] = i
	}
	ll := -1
	for i := 0; i < n-1; i++ {
		if y[a[i]] > i && a[i] != a[i+1] {
			ll = i
			break
		}
	}
	rr := -1
	for i := n - 1; i > 0; i-- {
		if x[a[i]] < i && a[i] != a[i-1] {
			rr = i
			break
		}
	}
	mem := make(map[int]bool)
	g1 := func(a1 int) {
		if a1 >= 0 && a1 < n {
			mem[a1] = true
		}
	}
	if ll != -1 && rr != -1 {
		vals := []int{a[ll], a[rr], a[ll+1], a[rr-1]}
		for _, v := range vals {
			g1(x[v] - 1)
			g1(x[v])
			g1(y[v])
			g1(y[v] + 1)
		}
		g1(ll)
		g1(ll + 1)
		g1(rr - 1)
		g1(rr)
	}
	var cands []int
	for i := range mem {
		cands = append(cands, i)
	}
	fg := false
	for i := 0; i < len(cands); i++ {
		for j := i + 1; j < len(cands); j++ {
			u, v := cands[i], cands[j]
			a[u], a[v] = a[v], a[u]
			if h1(a) {
				fg = true
				a[u], a[v] = a[v], a[u]
				break
			}
			a[u], a[v] = a[v], a[u]
		}
		if fg {
			break
		}
	}
	if fg {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
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
