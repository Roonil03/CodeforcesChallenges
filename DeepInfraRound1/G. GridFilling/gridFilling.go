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
	m := in.rInt()
	h := in.rInt()
	v := in.rInt()
	h++
	v++
	r := make([]int, h+1)
	c := make([]int, v+1)
	r[h] = n
	c[v] = m
	for i := 1; i < h; i++ {
		r[i] = in.rInt()
	}
	for i := 1; i < v; i++ {
		c[i] = in.rInt()
	}
	a := make([]int, h)
	b := make([]int, v)
	for i := 0; i < h; i++ {
		a[i] = r[i+1] - r[i]
	}
	for i := 0; i < v; i++ {
		b[i] = c[i+1] - c[i]
	}
	minSlice := func(arr []int) int {
		if len(arr) == 0 {
			return 0
		}
		res := arr[0]
		for _, val := range arr {
			if val < res {
				res = val
			}
		}
		return res
	}
	done := make([][]bool, h)
	for i := range done {
		done[i] = make([]bool, v)
	}
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, m)
	}
	if minSlice(a) == 1 && minSlice(b) == 1 {
		f := false
		for i := 0; i < h; i++ {
			if a[i] == 1 {
				if f {
					fmt.Fprintln(out, -1)
					return
				}
				f = true
			}
			if a[i] > 2 {
				f = false
			}
		}
		f = false
		for i := 0; i < v; i++ {
			if b[i] == 1 {
				if f {
					fmt.Fprintln(out, -1)
					return
				}
				f = true
			}
			if b[i] > 2 {
				f = false
			}
		}
	}
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	var seed uint64 = 88172645463325252
	randIntn := func(lim int) int {
		seed ^= seed << 13
		seed ^= seed >> 7
		seed ^= seed << 17
		return int(seed % uint64(lim))
	}
	fill := func(i, j int) {
		done[i][j] = true
		p := make([]int, a[i]*b[j])
		for k := 0; k < len(p); k++ {
			p[k] = k + 1
		}

		for {
			for k := len(p) - 1; k > 0; k-- {
				idx := randIntn(k + 1)
				p[k], p[idx] = p[idx], p[k]
			}
			for x := 0; x < a[i]; x++ {
				for y := 0; y < b[j]; y++ {
					g[r[i]+x][c[j]+y] = p[x*b[j]+y]
				}
			}
			fg := true
			for x := 0; x < a[i]; x++ {
				for y := 0; y < b[j]; y++ {
					if x != 0 && x != a[i]-1 && y != 0 && y != b[j]-1 {
						continue
					}
					for k := 0; k < 4; k++ {
						nx := r[i] + x + dx[k]
						ny := c[j] + y + dy[k]
						if 0 <= nx && nx < n && 0 <= ny && ny < m && g[nx][ny] == g[r[i]+x][c[j]+y] {
							fg = false
							break
						}
					}
					if !fg {
						break
					}
				}
				if !fg {
					break
				}
			}
			if fg {
				return
			}
		}
	}
	tryFill := func(maxr, maxc int, prer, prec bool) {
		for i := 0; i < h; i++ {
			for j := 0; j < v; j++ {
				if done[i][j] {
					continue
				}
				if a[i] > maxr || b[j] > maxc {
					continue
				}
				if prer && (i == 0 || !done[i-1][j]) {
					continue
				}
				if prec && (j == 0 || !done[i][j-1]) {
					continue
				}
				fill(i, j)
			}
		}
		for i := 0; i < h; i++ {
			for j := v - 1; j >= 0; j-- {
				if done[i][j] {
					continue
				}
				if a[i] > maxr || b[j] > maxc {
					continue
				}
				if prer && (i == 0 || !done[i-1][j]) {
					continue
				}
				if prec && (j == v-1 || !done[i][j+1]) {
					continue
				}
				fill(i, j)
			}
		}
		for i := h - 1; i >= 0; i-- {
			for j := 0; j < v; j++ {
				if done[i][j] {
					continue
				}
				if a[i] > maxr || b[j] > maxc {
					continue
				}
				if prer && (i == h-1 || !done[i+1][j]) {
					continue
				}
				if prec && (j == 0 || !done[i][j-1]) {
					continue
				}
				fill(i, j)
			}
		}
		for i := h - 1; i >= 0; i-- {
			for j := v - 1; j >= 0; j-- {
				if done[i][j] {
					continue
				}
				if a[i] > maxr || b[j] > maxc {
					continue
				}
				if prer && (i == h-1 || !done[i+1][j]) {
					continue
				}
				if prec && (j == v-1 || !done[i][j+1]) {
					continue
				}
				fill(i, j)
			}
		}
	}
	if minSlice(b) > 1 {
		for i := 0; i < h; i++ {
			if a[i] == 1 {
				for j := 0; j < v; j++ {
					done[i][j] = true
					for x := 0; x < b[j]; x++ {
						g[r[i]][c[j]+x] = x + 1
					}
					if r[i] > 0 && g[r[i]-1][c[j]] == 1 {
						for x := 0; x < b[j]; x++ {
							g[r[i]][c[j]+x] = (x+1)%b[j] + 1
						}
					}
				}
			}
		}
		for i := 0; i < h; i++ {
			for j := 0; j < v; j++ {
				if a[i] == 2 && b[j] == 2 {
					done[i][j] = true
					g[r[i]][c[j]] = 1
					g[r[i]][c[j]+1] = 2
					g[r[i]+1][c[j]] = 3
					g[r[i]+1][c[j]+1] = 4
					if r[i] > 0 && g[r[i]-1][c[j]] == 1 {
						g[r[i]][c[j]], g[r[i]][c[j]+1] = g[r[i]][c[j]+1], g[r[i]][c[j]]
					}
				}
			}
		}
	} else if minSlice(a) > 1 {
		for j := 0; j < v; j++ {
			if b[j] == 1 {
				for i := 0; i < h; i++ {
					done[i][j] = true
					for x := 0; x < a[i]; x++ {
						g[r[i]+x][c[j]] = x + 1
					}
					if c[j] > 0 && g[r[i]][c[j]-1] == 1 {
						for x := 0; x < a[i]; x++ {
							g[r[i]+x][c[j]] = (x+1)%a[i] + 1
						}
					}
				}
			}
		}
		for i := 0; i < h; i++ {
			for j := 0; j < v; j++ {
				if a[i] == 2 && b[j] == 2 {
					done[i][j] = true
					g[r[i]][c[j]] = 1
					g[r[i]][c[j]+1] = 3
					g[r[i]+1][c[j]] = 2
					g[r[i]+1][c[j]+1] = 4
					if c[j] > 0 && g[r[i]][c[j]-1] == 1 {
						g[r[i]][c[j]], g[r[i]+1][c[j]] = g[r[i]+1][c[j]], g[r[i]][c[j]]
					}
				}
			}
		}
	} else {
		tryFill(1, 1, false, false)
		tryFill(1, 2, false, true)
		tryFill(1, 2, false, false)
		tryFill(1, m, false, false)
		tryFill(2, 1, true, false)
		tryFill(2, 1, false, false)
		tryFill(n, 1, false, false)
		tryFill(2, 2, true, true)
		tryFill(2, 2, true, false)
		tryFill(2, 2, false, true)
		tryFill(2, 2, false, false)
	}
	tryFill(n, m, false, false)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, g[i][j])
		}
		fmt.Fprintln(out)
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

// Solution from https://codeforces.com/profile/jiangly

/*
This algorithm divides a 2D grid into smaller subgrids defined by horizontal and vertical boundary cuts and attempts to assign
unique integers (1 through a_i × b_j) to each cell within these subgrids such that no two adjacent cells share the same number.
To solve this, it employs a combination of deterministic placement rules for edge-case dimensions (like 1x1, 1x2, 2x1, and 2x2
grids) and a Las Vegas randomized approach with backtracking for arbitrarily sized blocks. The algorithm prioritizes filling
smaller adjacent components safely, sequentially executing predefined patterns where they fit perfectly and falling back to
repeated random shuffling of values for larger dimensions. Because a valid assignment space is heavily unconstrained for blocks
slightly larger than 2x2, the randomized configuration check hits a valid setup in a nearly constant expected number of iterations.
The overall Expected Time Complexity is O(N × M) per test case since the expected number of random placement tries per cell behaves
like O(1). The Space Complexity is also O(N × M) largely determined by the memory needed to store the generated configurations
and boundary state data.
*/
