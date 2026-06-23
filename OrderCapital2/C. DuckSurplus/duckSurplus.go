package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math/rand"
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

type Treap[K cmp.Ordered] struct {
	key         K
	pri, sz     int
	left, right *Treap[K]
}

func NewTreap[K cmp.Ordered](key K) *Treap[K] {
	return &Treap[K]{key: key, pri: rand.Int(), sz: 1}
}

func (t *Treap[K]) recalc() {
	t.sz = 1
	if t.left != nil {
		t.sz += t.left.sz
	}
	if t.right != nil {
		t.sz += t.right.sz
	}
}

func splitStrict[K cmp.Ordered](t *Treap[K], key K) (*Treap[K], *Treap[K]) {
	if t == nil {
		return nil, nil
	}
	if t.key < key {
		l, r := splitStrict(t.right, key)
		t.right = l
		t.recalc()
		return t, r
	}
	l, r := splitStrict(t.left, key)
	t.left = r
	t.recalc()
	return l, t
}

func splitEq[K cmp.Ordered](t *Treap[K], key K) (*Treap[K], *Treap[K]) {
	if t == nil {
		return nil, nil
	}
	if key < t.key {
		l, r := splitEq(t.left, key)
		t.left = r
		t.recalc()
		return l, t
	}
	l, r := splitEq(t.right, key)
	t.right = l
	t.recalc()
	return t, r
}

func merge[K cmp.Ordered](a, b *Treap[K]) *Treap[K] {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.pri < b.pri {
		a.right = merge(a.right, b)
		a.recalc()
		return a
	}
	b.left = merge(a, b.left)
	b.recalc()
	return b
}

func (t *Treap[K]) Insert(key K) *Treap[K] {
	n := NewTreap(key)
	a, b := splitEq(t, key)
	return merge(merge(a, n), b)
}

func (t *Treap[K]) Delete(key K) *Treap[K] {
	a, b := splitStrict(t, key)
	b1, c := splitEq(b, key)
	_ = b1
	return merge(a, c)
}

var mp map[*Treap[int64]]int64

func push(t *Treap[int64]) {
	if t == nil {
		return
	}
	if lz, ok := mp[t]; ok && lz != 0 {
		t.key += lz
		if t.left != nil {
			mp[t.left] += lz
		}
		if t.right != nil {
			mp[t.right] += lz
		}
		delete(mp, t)
	}
}

func lSplit(t *Treap[int64], key int64) (*Treap[int64], *Treap[int64]) {
	if t == nil {
		return nil, nil
	}
	push(t)
	if t.key <= key {
		l, r := lSplit(t.right, key)
		t.right = l
		t.recalc()
		return t, r
	}
	l, r := lSplit(t.left, key)
	t.left = r
	t.recalc()
	return l, t
}

func lMerge(a, b *Treap[int64]) *Treap[int64] {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.pri < b.pri {
		push(a)
		a.right = lMerge(a.right, b)
		a.recalc()
		return a
	}
	push(b)
	b.left = lMerge(a, b.left)
	b.recalc()
	return b
}

func h1(t *Treap[int64]) int64 {
	if t == nil {
		return 0
	}
	push(t)
	if t.right == nil {
		return t.key
	}
	return h1(t.right)
}

func precomp() {
	// mp = make(map[*Treap[int64]]int64)
}

func solve() {
	mp = make(map[*Treap[int64]]int64)
	n := in.rInt()
	arr := readInt64s(n)
	var root *Treap[int64]
	for i := range n {
		v := arr[i]
		l, r := lSplit(root, v)
		if r != nil {
			mp[r] += v
		}
		root = lMerge(lMerge(l, NewTreap(v)), r)
	}
	fmt.Fprintln(out, h1(root))
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
