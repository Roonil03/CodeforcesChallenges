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
	a := make([]int64, n)
	for i := range n {
		a[i] = in.rInt64()
	}
	g1 := func(arr []int64) bool {
		st := make([]int, 0, n)
		imos := make([]int, n+1)
		for i := range n {
			for len(st) > 0 && arr[st[len(st)-1]] < arr[i] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 && arr[st[len(st)-1]] == arr[i] {
				imos[st[len(st)-1]]++
				imos[i+1]--
			}
			st = append(st, i)
		}
		for i := range n {
			imos[i+1] += imos[i]
		}
		for i := range n {
			if imos[i] == 0 {
				return false
			}
		}
		return true
	}
	if g1(a) {
		fmt.Fprintln(out, 0)
		return
	}
	maxA := a[0]
	for i := 1; i < n; i++ {
		if a[i] > maxA {
			maxA = a[i]
		}
	}
	type pair struct {
		u int
		v int64
	}
	res := []pair{{0, maxA}, {n - 1, maxA}}
	g2 := func(u int, v int64) {
		b := make([]int64, n)
		copy(b, a)
		b[u] = v
		if g1(b) {
			res = []pair{{u, v}}
		}
	}
	for range 2 {
		g2(0, maxA)
		pos := 0
		for i := range n {
			if a[i] == maxA {
				pos = i
			}
		}
		if pos != n-1 {
			v := a[pos+1]
			for i := pos + 1; i < n; i++ {
				if a[i] > v {
					v = a[i]
				}
			}
			g2(pos, v)
		}
		g2(pos, 1)
		for i := range res {
			res[i].u = n - 1 - res[i].u
		}
		for i := 0; i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	fmt.Fprintln(out, len(res))
	for _, p := range res {
		fmt.Fprintln(out, p.u+1, p.v)
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

// Solution by https://codeforces.com/profile/Nachia

/*
This algorithm finds the minimum number of operations to make an array "good", which means every position in the array is covered
by an interval bounded by equal elements with no strictly larger elements between them. A monotonic stack algorithm is used inside
`isGood` to efficiently track these valid pairs and build a difference array (`imos`) that computes the prefix sum of their
intervals; if the interval coverage spans all items, the array is considered good. If it's not initially good, the heuristic logic
explores a maximum of two operations: either setting both extreme ends (index 0 and N-1) to the global maximum of the array
(which will always satisfy the condition), or testing a few edge-case single operation adjustments. The single adjustments check
placing the array's maximum at the ends or targeting the index of the existing maximum value, all handled systematically forwards
and backwards using symmetry. The overall Time Complexity is O(N) per test case since checking validity (stack push/pops and
difference array evaluations) runs in linear time, and the modifications run a constant O(1) number of times. The Space Complexity
is O(N) to accommodate the monotonic stack, difference tracking arrays, and temporary array copies during tests.
*/
