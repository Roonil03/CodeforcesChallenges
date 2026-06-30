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

type Fenwick[T cmp.Ordered] struct {
	n    int
	tree []T
	add  func(T, T) T
}

func NewFenwick[T cmp.Ordered](n int, add func(T, T) T) *Fenwick[T] {
	return &Fenwick[T]{n: n, tree: make([]T, n+1), add: add}
}

func (f *Fenwick[T]) Update(i int, v T) {
	for ; i <= f.n; i += i & -i {
		f.tree[i] = f.add(f.tree[i], v)
	}
}

func (f *Fenwick[T]) Query(i int, zero T) T {
	sum := zero
	for ; i > 0; i -= i & -i {
		sum = f.add(sum, f.tree[i])
	}
	return sum
}

func (f *Fenwick[T]) Kth(k T) int {
	pos := 0
	var res T
	for i := 20; i >= 0; i-- {
		if pos+(1<<i) <= f.n {
			next := f.add(res, f.tree[pos+(1<<i)])
			if next < k {
				pos += 1 << i
				res = next
			}
		}
	}
	return pos + 1
}

func precomp() {
}

func solve() {
	n := in.rInt()
	ch := make([]byte, n+1)
	val := make([]int, n+1)
	p := make([]int, n+1)
	var pos []int
	vis := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		s := in.rString()
		ch[i] = s[0]
		val[i] = in.rInt()
		if ch[i] == 'p' {
			p[i] = val[i]
			vis[p[i]] = true
		} else {
			pos = append(pos, i)
		}
	}

	add := func(a, b int) int { return a + b }
	T := NewFenwick(n, add)
	cur := NewFenwick(n, add)
	blk := NewFenwick(n, add)

	for x := 1; x <= n; x++ {
		T.Update(x, 1)
		if !vis[x] {
			cur.Update(x, 1)
		}
	}

	if len(pos) == 0 {
		printSlice(p[1:])
		return
	}

	for i := n; i > pos[len(pos)-1]; i-- {
		T.Update(p[i], -1)
	}

	for i := len(pos) - 1; i >= 0; i-- {
		L, R := 0, pos[i]
		if i > 0 {
			L = pos[i-1]
		}

		res, cnt := 0, 0
		var vec []int

		for j := R - 1; j >= L+1; j-- {
			res += j - T.Query(p[j], 0) + blk.Query(p[j], 0)
			blk.Update(p[j], 1)
			vec = append(vec, p[j])
			cnt++
		}

		calc := func(w int) int {
			return res + (R - T.Query(w, 0)) + (cnt - blk.Query(w-1, 0))
		}

		posl := n + 1
		l, r := 1, n
		for l <= r {
			mid := (l + r) >> 1
			if calc(mid) <= val[R]-val[L] {
				posl = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		w := cur.Kth(cur.Query(posl-1, 0) + 1)
		p[R] = w

		cur.Update(w, -1)
		T.Update(w, -1)
		for _, x := range vec {
			T.Update(x, -1)
			blk.Update(x, -1)
		}
	}

	printSlice(p[1:])
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

/*
The algorithm reconstructs a permutation by determining the values of missing elements dynamically from right to
left. It employs three Binary Indexed Trees (Fenwick Trees): one to track all active elements, one to track the
currently available (unassigned) numbers, and one to track elements within the local block being processed. For
each missing element, the algorithm counts the inversions contributed by known elements in the current block, then
performs a binary search over the possible values. The binary search maps values to the required inversion
differences to deduce the original element. Once the logical position is found, the exact available value is extracted
using a k-th element query on the available elements' BIT.

Time Complexity: O(N log^2 N) per testcase. For each missing element, a binary search evaluates a condition using
point-queries on the Fenwick Trees, resulting in logarithmic tree query times strictly nested inside the logarithmic
binary search span.

Space Complexity: O(N) due to the use of slices tracking character states, values, permutations, visited properties,
and the internal array allocations required by the Fenwick Tree structs.
*/
