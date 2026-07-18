package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/bits"
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

func clearFenwick2D(f *Fenwick2D[int]) {
	for i := range f.tree {
		for j := range f.tree[i] {
			f.tree[i][j] = 0
		}
	}
}

func addFenwick(f *Fenwick2D[int], x, y int) {
	f.Update(x+1, y+1, 1)
}

func getFenwick(f *Fenwick2D[int], x0, x1, y0, y1 int) int {
	q := func(x, y int) int {
		return f.Query(x+1, y+1, 0)
	}
	return q(x1, y1) - q(x0-1, y1) - q(x1, y0-1) + q(x0-1, y0-1)
}

const small = 32

type pt struct {
	y, v    int
	x, u, d uint8
}

type ev struct {
	x       int
	a, b, d uint8
	w       int8
}

type rows [100][2]uint64

type rmq struct {
	n, m, nb    int
	v, pre, suf []int
	t           [][]int
}

func newRMQ(nn, mm int, a []int) *rmq {
	nb := (nn + 8 - 1) / 8
	res := &rmq{
		n: nn, m: mm, nb: nb,
		v: a, pre: make([]int, nn*mm), suf: make([]int, nn*mm),
		t: make([][]int, 1),
	}
	res.t[0] = make([]int, nb*mm)
	for b := range nb {
		l := b * 8
		r := min(nn, l+8)
		copy(res.pre[l*mm:(l+1)*mm], a[l*mm:(l+1)*mm])
		for i := l + 1; i < r; i++ {
			for j := range mm {
				res.pre[i*mm+j] = max(res.pre[(i-1)*mm+j], a[i*mm+j])
			}
		}
		copy(res.suf[(r-1)*mm:r*mm], a[(r-1)*mm:r*mm])
		for i := r - 2; i >= l; i-- {
			for j := range mm {
				res.suf[i*mm+j] = max(res.suf[(i+1)*mm+j], a[i*mm+j])
			}
		}
		copy(res.t[0][b*mm:(b+1)*mm], res.pre[(r-1)*mm:r*mm])
	}
	for k := 1; (1 << k) <= nb; k++ {
		h := 1 << (k - 1)
		w := h << 1
		res.t = append(res.t, make([]int, (nb-w+1)*mm))
		for i := 0; i+w <= nb; i++ {
			for j := range mm {
				res.t[k][i*mm+j] = max(res.t[k-1][i*mm+j], res.t[k-1][(i+h)*mm+j])
			}
		}
	}
	return res
}

func (q *rmq) next(i, j int) int {
	return q.v[i*q.m+j]
}

func (q *rmq) get(j, l, r int) int {
	x := l / 8
	y := r / 8
	if x == y {
		z := q.v[l*q.m+j]
		for i := l + 1; i <= r; i++ {
			if q.v[i*q.m+j] > z {
				z = q.v[i*q.m+j]
			}
		}
		return z
	}
	z := q.suf[l*q.m+j]
	if preVal := q.pre[r*q.m+j]; preVal > z {
		z = preVal
	}
	x++
	y--
	if x <= y {
		k := bits.Len(uint(y-x+1)) - 1
		if val1 := q.t[k][x*q.m+j]; val1 > z {
			z = val1
		}
		if val2 := q.t[k][(y-(1<<k)+1)*q.m+j]; val2 > z {
			z = val2
		}
	}
	return z
}

func vpath(a, b pt, q *rmq) bool {
	p, r := a, b
	if p.y > r.y {
		p, r = r, p
	}
	l := max(r.u, p.u)
	h := p.d
	if r.d < h {
		h = r.d
	}
	return l <= h && q.get(int(p.y), int(l), int(h)) >= r.y
}

func orderPoints(p []pt, z int) []int {
	c := make([]int, z+1)
	o := make([]int, len(p))
	for _, x := range p {
		c[x.v]++
	}
	for i := 1; i <= z; i++ {
		c[i] += c[i-1]
	}
	for i := len(p) - 1; i >= 0; i-- {
		c[p[i].v]--
		o[c[p[i].v]] = i
	}
	return o
}

func mark(p []pt, o []int, s *rows) bool {
	any := false
	for l, r := 0, 0; l < len(o); l = r {
		for r = l + 1; r < len(o) && p[o[r]].v == p[o[l]].v; r++ {
		}
		if r-l < 2 {
			continue
		}
		any = true
		var q [2]uint64
		for i := l; i < r; i++ {
			x := p[o[i]].x
			q[x>>6] |= 1 << (x & 63)
		}
		for i := l; i < r; i++ {
			x := p[o[i]].x
			s[x][0] |= q[0]
			s[x][1] |= q[1]
		}
	}
	return any
}

func vcount(n, m int, p []pt, o []int, q *rmq) int64 {
	var ans int64
	a := make([]int, 100)
	pl := make([]int, 100)
	nr := make([]int, 100)
	st := make([]int, 100)
	f := NewFenwick2D[int](102, 102, func(u, v int) int {
		return u + v
	})
	for l, r := 0, 0; l < len(o); l = r {
		for r = l + 1; r < len(o) && p[o[r]].v == p[o[l]].v; r++ {
		}
		if r-l < 2 {
			continue
		}
		if r-l <= small {
			for i := l; i < r; i++ {
				for j := i + 1; j < r; j++ {
					if p[o[i]].y != p[o[j]].y && vpath(p[o[i]], p[o[j]], q) {
						ans++
					}
				}
			}
			continue
		}
		var ne int
		for i := l; i < r; i++ {
			ne += int(p[o[i]].d - p[o[i]].u + 1)
		}
		e := make([]ev, 0, 2*ne)
		for ii := l; ii < r; ii++ {
			x := p[o[ii]]
			h := int(x.d - x.u + 1)
			for i := range h {
				a[i] = q.next(int(x.u)+i, x.y)
			}
			z := 0
			for i := range h {
				for z > 0 && a[st[z-1]] < a[i] {
					z--
				}
				if z > 0 {
					pl[i] = st[z-1]
				} else {
					pl[i] = -1
				}
				st[z] = i
				z++
			}
			z = 0
			for i := h - 1; i >= 0; i-- {
				for z > 0 && a[st[z-1]] <= a[i] {
					z--
				}
				if z > 0 {
					nr[i] = st[z-1]
				} else {
					nr[i] = h
				}
				st[z] = i
				z++
			}
			for i := range h {
				y := min(a[i], m-1)
				if y <= x.y {
					continue
				}
				k := int(x.u) + i
				u := 0
				if pl[i] >= 0 {
					u = int(x.u) + pl[i] + 1
				}
				d := n - 1
				if nr[i] != h {
					d = int(x.u) + nr[i] - 1
				}
				e = append(e, ev{x: y, a: uint8(u), b: uint8(k), d: uint8(d), w: 1})
				e = append(e, ev{x: x.y, a: uint8(u), b: uint8(k), d: uint8(d), w: -1})
			}
		}
		subO := o[l:r]
		sort.Slice(subO, func(i, j int) bool {
			return p[subO[i]].y < p[subO[j]].y
		})
		sort.Slice(e, func(i, j int) bool {
			return e[i].x < e[j].x
		})
		clearFenwick2D(f)
		j := l
		for _, x := range e {
			for j < r && p[o[j]].y <= x.x {
				addFenwick(f, int(p[o[j]].u), int(p[o[j]].d))
				j++
			}
			ans += int64(x.w) * int64(getFenwick(f, int(x.a), int(x.b), int(x.b), int(x.d)))
		}
	}
	return ans
}

func hcount(n, m int, p []pt, rb []int, dn []uint8, q *rmq, same *rows) int64 {
	var ans int64
	s := make([]int, m+1)
	for x := range n {
		for y := x + 1; y < n; y++ {
			if (same[x][y>>6]>>(y&63))&1 == 0 {
				continue
			}
			d := dn[x*m : (x+1)*m]
			for j := 0; j < m; j++ {
				s[j+1] = s[j]
				if int(d[j]) >= y {
					s[j+1]++
				}
			}
			j := rb[y]
			for i := rb[x]; i < rb[x+1]; i++ {
				al := 0
				if i != rb[x] {
					al = p[i-1].y + 1
				}
				ar := m - 1
				if i+1 != rb[x+1] {
					ar = p[i+1].y - 1
				}
				for j < rb[y+1] {
					br := m - 1
					if j+1 != rb[y+1] {
						br = p[j+1].y - 1
					}
					if br >= al {
						break
					}
					j++
				}
				for k := j; k < rb[y+1]; k++ {
					bl := 0
					if k != rb[y] {
						bl = p[k-1].y + 1
					}
					if bl > ar {
						break
					}
					br := m - 1
					if k+1 != rb[y+1] {
						br = p[k+1].y - 1
					}
					l := max(al, bl)
					r := min(ar, br)
					if p[i].v != p[k].v || s[r+1] == s[l] {
						continue
					}
					old := false
					if p[i].y == p[k].y {
						old = int(d[p[i].y]) == y
					} else {
						old = vpath(p[i], p[k], q)
					}
					if !old {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func precomp() {

}

func solve() {
	n := in.rInt()
	m := in.rInt()
	z := n * m
	var ans int64
	rb := make([]int, n+1)
	last := make([]int, m)
	for i := range last {
		last[i] = -1
	}
	p := make([]pt, 0, z)
	for i := range n {
		rb[i] = len(p)
		for j := range m {
			v := in.rInt()
			if v == 0 {
				continue
			}
			h := last[j]
			k := len(p)
			var u uint8
			if h >= 0 {
				u = uint8(p[h].x + 1)
			}
			p = append(p, pt{y: j, v: v, x: uint8(i), u: u, d: 0})
			if h >= 0 {
				p[h].d = uint8(i - 1)
				if p[h].v == v {
					ans++
				}
			}
			last[j] = k
		}
	}
	rb[n] = len(p)
	for j := range m {
		if last[j] >= 0 {
			p[last[j]].d = uint8(n - 1)
		}
	}
	o := orderPoints(p, z)
	var same rows
	if !mark(p, o, &same) {
		fmt.Fprintln(out, 0)
		return
	}
	dn := make([]uint8, z)
	rt := make([]int, z)
	for j := range last {
		last[j] = n
	}
	for i := n - 1; i >= 0; i-- {
		k := rb[i+1] - 1
		r := m
		for j := m - 1; j >= 0; j-- {
			w := i*m + j
			dn[w] = uint8(last[j])
			rt[w] = r
			if k >= rb[i] && p[k].y == j {
				last[j] = i
				r = j
				k--
			}
		}
	}
	q := newRMQ(n, m, rt)
	ans += vcount(n, m, p, o, q)
	ans += hcount(n, m, p, rb, dn, q, &same)
	fmt.Fprintln(out, ans)
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

// Solution referenced from https://codeforces.com/profile/vangoat

/*
The solution counts paths mapped across identical values connecting locations on a grid matrix. It effectively utilizes a
customized sparse table RMQ built to accelerate horizontal boundary path checks dynamically over adjacent bounds. Vertical
sweeping is employed to systematically scan grouped matching values over row spans. An integrated 2-Dimensional Fenwick
tree accurately monitors active nodes during the sweep to effectively determine overlaps utilizing a 4-point inclusivity
approach (`FenwickRange2D`). Concurrently, horizontal paths between identical values are validated and resolved by bypassing
unnecessary states referencing pre-populated optimized 64-bit row-level matching bitmasks. The algorithm efficiently balances
both memory scale and execution overhead. The time complexity equates broadly to $O(N^2 M + N M^2)$ typically accelerated
substantially in sparse test cases due to overlapping masking pruning and localized checks scaling optimally over overlapping
event queries. The space complexity is confined to roughly $O(N M \log(N))$ constrained intrinsically by dynamically reserving
RMQ blocks overlaying grid footprints strictly bounding memory thresholds across constraints.
*/
