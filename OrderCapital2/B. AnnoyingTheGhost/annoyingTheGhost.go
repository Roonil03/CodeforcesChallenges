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
	a := readInt64s(n)
	b := readInt64s(n)
	c := make([]int, n)
	fg := true
	for i := range n {
		id := -1
		l, r := 0, n-1
		for l <= r {
			m := (l + r) / 2
			if b[m] >= a[i] {
				id = m
				r = m - 1
			} else {
				l = m + 1
			}
		}
		if id == -1 {
			fg = false
			break
		}
		c[i] = id
	}
	if !fg {
		fmt.Fprintln(out, -1)
		return
	}
	res := 0
	f, freq := make([]bool, n), make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := range n {
			if !f[j] && c[j] > i {
				fg = false
				break
			}
		}
		if !fg {
			break
		}
		for j := range n + 1 {
			freq[j] = 0
		}
		for j := range n {
			if !f[j] {
				freq[c[j]]++
			}
		}
		count, k := 0, -1
		for j := i; j >= 0; j-- {
			count += freq[j]
			if count > i-j+1 {
				fg = false
				break
			}
			if count == i-j+1 && count > 0 && k == -1 {
				k = j
			}
		}
		if !fg {
			break
		}
		d := -1
		if k != -1 {
			for j := n - 1; j >= 0; j-- {
				if !f[j] && c[j] >= k {
					d = j
					break
				}
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				if !f[i] {
					d = j
					break
				}
			}
		}
		temp := 0
		for j := range d {
			if f[j] {
				temp++
			}
		}
		res += temp
		f[d] = true
	}
	if !fg {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, res)
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
