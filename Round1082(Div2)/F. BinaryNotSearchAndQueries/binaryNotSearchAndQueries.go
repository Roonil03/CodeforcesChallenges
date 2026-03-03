package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, _ := in.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInts(n int) []int {
	parts := strings.Fields(readLine())
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])
	a := readInts(n)
	S := make([]*Treap, n+1)
	st := NewSegTree(n)
	removeId := func(id, v int) {
		if S[v] == nil {
			return
		}
		first := S[v].Min()
		st.Update(first, -1)
		S[v] = S[v].Delete(id)
		if S[v] != nil {
			newF := S[v].Min()
			newL := S[v].Max()
			st.Update(newF, newL-newF)
		}
	}
	addId := func(id, v int) {
		if S[v] != nil {
			oldF := S[v].Min()
			st.Update(oldF, -1)
			S[v] = S[v].Insert(id)
		} else {
			S[v] = NewTreap(id)
		}
		newF := S[v].Min()
		newL := S[v].Max()
		st.Update(newF, newL-newF)
	}
	for i := range n {
		addId(i+1, a[i])
	}
	for ; q > 0; q-- {
		line := readLine()
		parts := strings.Fields(line)
		id, _ := strconv.Atoi(parts[0])
		x, _ := strconv.Atoi(parts[1])
		u := a[id-1]
		if u != x {
			removeId(id, u)
			addId(id, x)
			a[id-1] = x
		}
		currentMax := st.GetMax()
		if currentMax <= 0 {
			fmt.Fprintln(out, 0, 0)
		} else {
			fmt.Fprintln(out, currentMax, st.GetAns())
		}
	}
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t, _ := strconv.Atoi(readLine())
	for t > 0 {
		t--
		solve()
	}
}

type Treap struct {
	key, pri, sz int
	left, right  *Treap
}

func NewTreap(key int) *Treap {
	return &Treap{key: key, pri: rand.Int(), sz: 1}
}

func (t *Treap) recalc() {
	t.sz = 1
	if t.left != nil {
		t.sz += t.left.sz
	}
	if t.right != nil {
		t.sz += t.right.sz
	}
}

func split(t *Treap, key int) (*Treap, *Treap) {
	if t == nil {
		return nil, nil
	}
	if key < t.key {
		l, r := split(t.left, key)
		t.left = r
		t.recalc()
		return l, t
	}
	l, r := split(t.right, key)
	t.right = l
	t.recalc()
	return t, r
}

func merge(a, b *Treap) *Treap {
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

func (t *Treap) Insert(key int) *Treap {
	n := NewTreap(key)
	a, b := split(t, key)
	return merge(merge(a, n), b)
}

func (t *Treap) Delete(key int) *Treap {
	a, b := split(t, key-1)
	b, c := split(b, key)
	return merge(a, c)
}

func (t *Treap) Min() int {
	cur := t
	for cur.left != nil {
		cur = cur.left
	}
	return cur.key
}

func (t *Treap) Max() int {
	cur := t
	for cur.right != nil {
		cur = cur.right
	}
	return cur.key
}

type Node struct {
	val  int
	ans  int64
	pref int
	suff int
	len  int
}

type SegTree struct {
	n    int
	tree []Node
}

func comb(L int64) int64 {
	return L * (L + 1) / 2
}

func mergeNode(A, B Node) Node {
	var C Node
	C.len = A.len + B.len
	if A.val > B.val {
		C.val = A.val
		C.ans = A.ans
		C.pref = A.pref
		C.suff = 0
	} else if B.val > A.val {
		C.val = B.val
		C.ans = B.ans
		C.pref = 0
		C.suff = B.suff
	} else {
		C.val = A.val
		C.ans = A.ans + B.ans + comb(int64(A.suff+B.pref)) - comb(int64(A.suff)) - comb(int64(B.pref))
		C.pref = A.pref
		if A.pref == A.len {
			C.pref = A.len + B.pref
		}
		C.suff = B.suff
		if B.suff == B.len {
			C.suff = B.len + A.suff
		}
	}
	return C
}

func NewSegTree(n int) *SegTree {
	st := &SegTree{n: n}
	st.tree = make([]Node, 4*n+5)
	st.build(1, 1, n)
	return st
}

func (st *SegTree) build(node, l, r int) {
	if l == r {
		st.tree[node] = Node{-1, 0, 0, 0, 1}
		return
	}
	mid := (l + r) >> 1
	st.build(node<<1, l, mid)
	st.build(node<<1|1, mid+1, r)
	st.tree[node] = mergeNode(st.tree[node<<1], st.tree[node<<1|1])
}

func (st *SegTree) update(node, l, r, id, val int) {
	if l == r {
		st.tree[node].val = val
		if val == -1 {
			st.tree[node].ans = 0
			st.tree[node].pref = 0
			st.tree[node].suff = 0
		} else {
			st.tree[node].ans = 1
			st.tree[node].pref = 1
			st.tree[node].suff = 1
		}
		return
	}
	mid := (l + r) >> 1
	if id <= mid {
		st.update(node<<1, l, mid, id, val)
	} else {
		st.update(node<<1|1, mid+1, r, id, val)
	}
	st.tree[node] = mergeNode(st.tree[node<<1], st.tree[node<<1|1])
}

func (st *SegTree) Update(id, val int) {
	st.update(1, 1, st.n, id, val)
}

func (st *SegTree) GetMax() int {
	return st.tree[1].val
}

func (st *SegTree) GetAns() int64 {
	return st.tree[1].ans
}

/*
Uses my versions of Treap and a slightly modified version of my SegmentTree

kmax(a) is equal to the maximum |i − j| such that a[i] = a[j].
For every value v, only the first and last occurrence matter.

For each value v:
Let first = smallest index of v
Let last  = largest index of v
Candidate k = last − first

For each value v → a set of indices.
For each value → at most one candidate k.

Store all candidates in a segment tree indexed by first position.

At position first[v], store value:
    last[v] − first[v]
Otherwise store −1.

Segment tree maintains:
- Maximum k value.
- Number of consecutive positions achieving maximum,
  using prefix/suffix merging.
- For each chain of equal k,
  contribution = L*(L+1)/2.

For each update:
1. Remove old index from its value set.
2. Update candidate for that value.
3. Insert index into new value set.
4. Update candidate for that value.

Each operation costs O(log n).

Final answer:
kmax = segmentTree.getMax()
f(a) = segmentTree.getAns()

Time Complexity: O((n + q) log n)

*/
