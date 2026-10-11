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
	k := in.rInt()
	a := readInts(n)
	var qsort func(l, r int)
	qsort = func(l, r int) {
		if l >= r {
			return
		}
		p := a[l+(r-l)/2]
		i, j := l, r
		for i <= j {
			for a[i] < p {
				i++
			}
			for a[j] > p {
				j--
			}
			if i <= j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
		}
		qsort(l, j)
		qsort(i, r)
	}
	qsort(0, n-1)
	pref := make([]int, n+1)
	for i := range n {
		pref[a[i]]++
	}
	for i := k; i <= n; i++ {
		pref[i]--
	}
	for i := range n {
		pref[i+1] += pref[i]
	}
	low, high := -1, n
	for low+1 < high {
		mid := (low + high) / 2
		b := make([]int, n+1)
		copy(b, pref)
		for i := 0; i <= n; i++ {
			b[i] += 2 * mid
		}
		f := make([]int, n+1)
		for i := range n {
			f[a[i]]++
		}
		c := make([]int, n+1)
		var st []int
		ok := true
		used := 0
		for i := n; i >= 0; i-- {
			for b[i]+c[i] < 0 && used <= mid {
				if len(st) == 0 {
					ok = false
					break
				}
				x := st[len(st)-1]
				st = st[:len(st)-1]
				c[i]++
				used++
				nx := (x + 1) / 2
				idx := min(nx-1, i)
				c[idx] -= 2
				if nx > i {
					st = append(st, nx, nx)
				} else {
					f[nx] += 2
				}
			}
			if !ok || used > mid {
				break
			}
			if i > 0 {
				c[i-1] += c[i]
			}
			for j := 0; j < f[i]; j++ {
				st = append(st, i)
			}
		}
		if ok && used <= mid {
			high = mid
		} else {
			low = mid
		}
	}
	fmt.Fprintln(out, n+2*high)
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

// Solution by https://codeforces.com/profile/tourist

/*
This algorithm finds the minimum number of division-by-two operations required to transform a multiset of numbers to satisfy
specific bounds via a combination of greedy prefix matching and binary search. First, it computes a generalized prefix sum
array (`pref`) capturing element frequencies relative to the constraint thresholds. The search dynamically tests if a candidate
number of operations (`mid`) is sufficient. During the evaluation of `mid`, it iterates backwards from the largest possible
value `n` down to `0`, eagerly utilizing a stack (`st`) to simulate subdividing larger available elements (`(x+1)/2`) whenever
the accumulated prefix constraint `b[i] + c[i] < 0` demands reinforcement. The arrays track shifting frequency pools (`f`) and
cascading prefix demands (`c`), terminating early if the constraints become impossible to satisfy or operations exceed `mid`.
The overall time complexity is bounded by O(n log n) per test case, dominated by sorting the input array and the binary search
iteratively validating state configurations in amortized O(n) time. The space complexity is O(n) due to allocating and maintaining
various prefix structures, frequency arrays, and the operational stack dynamically at execution time.
*/
