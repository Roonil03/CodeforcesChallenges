package main

import (
	"bufio"
	"fmt"
	"os"
)

type Seg struct {
	v  []int
	ad []int
}

func NewSeg(size int) *Seg {
	return &Seg{
		v:  make([]int, size*4),
		ad: make([]int, size*4),
	}
}

func (tr *Seg) pushUp(p int) {
	tr.v[p] = tr.v[p<<1]
	if tr.v[p<<1|1] < tr.v[p] {
		tr.v[p] = tr.v[p<<1|1]
	}
}

func (tr *Seg) upd(p, k int) {
	tr.v[p] += k
	tr.ad[p] += k
}

func (tr *Seg) pushDown(p int) {
	if tr.ad[p] != 0 {
		tr.upd(p<<1, tr.ad[p])
		tr.upd(p<<1|1, tr.ad[p])
		tr.ad[p] = 0
	}
}

func (tr *Seg) uadd(p, l, r, tl, trVal, k int) {
	if tl > trVal {
		return
	}
	if tl <= l && r <= trVal {
		tr.upd(p, k)
		return
	}
	mid := (l + r) >> 1
	tr.pushDown(p)
	if tl <= mid {
		tr.uadd(p<<1, l, mid, tl, trVal, k)
	}
	if trVal > mid {
		tr.uadd(p<<1|1, mid+1, r, tl, trVal, k)
	}
	tr.pushUp(p)
}

func (tr *Seg) build(p, l, r int) {
	tr.v[p] = 0
	tr.ad[p] = 0
	if l == r {
		return
	}
	mid := (l + r) >> 1
	tr.build(p<<1, l, mid)
	tr.build(p<<1|1, mid+1, r)
}

func (tr *Seg) qmin(p, l, r, tl, trVal int) int {
	if tl <= l && r <= trVal {
		return tr.v[p]
	}
	mid := (l + r) >> 1
	res := int(1e9)
	tr.pushDown(p)
	if tl <= mid {
		val := tr.qmin(p<<1, l, mid, tl, trVal)
		if val < res {
			res = val
		}
	}
	if trVal > mid {
		val := tr.qmin(p<<1|1, mid+1, r, tl, trVal)
		if val < res {
			res = val
		}
	}
	return res
}

type Node struct {
	l, r, k     int
	priority    int
	left, right *Node
}

type ODTBlock struct {
	l, r, k int
}

type Pii struct {
	fi, se int
}

func newNode(l, r, k int) *Node {
	return &Node{l: l, r: r, k: k, priority: randInt()}
}

var randState = uint32(1337)

func randInt() int {
	randState ^= randState << 13
	randState ^= randState >> 17
	randState ^= randState << 5
	return int(randState & 0x7fffffff)
}

func split(t *Node, l int, lNode, rNode **Node) {
	if t == nil {
		*lNode = nil
		*rNode = nil
		return
	}
	if t.l < l {
		*lNode = t
		split(t.right, l, &t.right, rNode)
	} else {
		*rNode = t
		split(t.left, l, lNode, &t.left)
	}
}

func merge(l, r *Node) *Node {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.priority > r.priority {
		l.right = merge(l.right, r)
		return l
	} else {
		r.left = merge(l, r.left)
		return r
	}
}

func insert(root **Node, n *Node) {
	var l, r *Node
	split(*root, n.l, &l, &r)
	*root = merge(merge(l, n), r)
}

func erase(root **Node, l int) {
	var lt, rt *Node
	split(*root, l, &lt, &rt)
	var eq, gt *Node
	split(rt, l+1, &eq, &gt)
	*root = merge(lt, gt)
}

func lowerBound(root *Node, l int) *Node {
	curr := root
	var ans *Node
	for curr != nil {
		if curr.l >= l {
			ans = curr
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return ans
}

func prevNode(root *Node, l int) *Node {
	curr := root
	var ans *Node
	for curr != nil {
		if curr.l < l {
			ans = curr
			curr = curr.right
		} else {
			curr = curr.left
		}
	}
	return ans
}

var (
	e       [][]int
	s       []int
	gs      []int
	sz      []int
	dfn     []int
	top     []int
	ft      []int
	dep     []int
	dt      int
	rt      int
	n, m, q int
	stRoot  *Node
	segment *Seg
)

func initODT() {
	stRoot = nil
	for i := 0; i <= n+1; i++ {
		insert(&stRoot, newNode(i, i, 0))
	}
}

func findODT(l, r int) []ODTBlock {
	it := prevNode(stRoot, l)
	if it != nil && it.r >= l {
		tmp := *it
		erase(&stRoot, tmp.l)
		if tmp.l < l {
			insert(&stRoot, newNode(tmp.l, l-1, tmp.k))
		}
		insert(&stRoot, newNode(l, tmp.r, tmp.k))
	}
	var res []ODTBlock
	trVal := l
	for {
		it := lowerBound(stRoot, trVal)
		if it == nil || it.l > r {
			break
		}
		tmp := *it
		rBound := tmp.r
		if rBound > r {
			rBound = r
		}
		res = append(res, ODTBlock{l: tmp.l, r: rBound, k: tmp.k})
		if tmp.r > r {
			erase(&stRoot, tmp.l)
			insert(&stRoot, newNode(tmp.l, r, tmp.k))
			insert(&stRoot, newNode(r+1, tmp.r, tmp.k))
		}
		trVal = tmp.r + 1
	}
	return res
}

func dfs(p, fa int) {
	sz[p] = 1
	gs[p] = 0
	ft[p] = fa
	dep[p] = dep[fa] + 1
	for _, i := range e[p] {
		if i == fa {
			continue
		}
		dfs(i, p)
		sz[p] += sz[i]
		if sz[i] > sz[gs[p]] {
			gs[p] = i
		}
	}
}

func dfs2(p, tp int) {
	dt++
	dfn[p] = dt
	top[p] = tp
	if gs[p] != 0 {
		dfs2(gs[p], tp)
	}
	for _, i := range e[p] {
		if i == ft[p] || i == gs[p] {
			continue
		}
		dfs2(i, i)
	}
}

func LCA(x, y int) int {
	for top[x] != top[y] {
		if dfn[x] < dfn[y] {
			x, y = y, x
		}
		x = ft[top[x]]
	}
	if dfn[x] < dfn[y] {
		return x
	}
	return y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func query(p, col int) int {
	P := p
	var tx, ty []Pii
	for top[p] != top[rt] {
		if dfn[p] > dfn[rt] {
			tx = append(tx, Pii{dfn[top[p]], dfn[p]})
			p = ft[top[p]]
		} else {
			ty = append(ty, Pii{dfn[rt], dfn[top[rt]]})
			rt = ft[top[rt]]
		}
	}
	if dfn[p] < dfn[rt] {
		ty = append(ty, Pii{dfn[rt], dfn[p]})
	} else {
		tx = append(tx, Pii{dfn[rt], dfn[p]})
	}
	for i, j := 0, len(tx)-1; i < j; i, j = i+1, j-1 {
		tx[i], tx[j] = tx[j], tx[i]
	}
	var td []ODTBlock
	for _, pair := range ty {
		r := pair.fi
		l := pair.se
		tmp := findODT(l, r)
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
		for _, blk := range tmp {
			if blk.k != col-1 {
				td = append(td, blk)
			}
		}
	}
	for _, pair := range tx {
		l := pair.fi
		r := pair.se
		tmp := findODT(l, r)
		for _, blk := range tmp {
			if blk.k != col-1 {
				td = append(td, blk)
			}
		}
	}
	lenVal := 0
	for _, blk := range td {
		lenVal += blk.r - blk.l + 1
	}
	if lenVal == 0 {
		return -1
	}
	res := lenVal
	var mdy []ODTBlock
	for _, blk := range td {
		erase(&stRoot, blk.l)
		tlen := blk.r - blk.l + 1
		cnt := segment.qmin(1, 1, q+1, blk.k+1, col-1)
		res -= min(tlen, cnt)
		mdy = append(mdy, ODTBlock{blk.l, blk.r, col - 1})
		segment.uadd(1, 1, q+1, blk.k+1, col-1, -min(tlen, cnt))
	}
	for i := 0; i < len(mdy); {
		j := i + 1
		rVal := mdy[i].r
		for j < len(mdy) && mdy[j].l == mdy[j-1].r+1 && mdy[j].k == mdy[j-1].k {
			rVal = mdy[j].r
			j++
		}
		insert(&stRoot, newNode(mdy[i].l, rVal, mdy[i].k))
		i = j
	}
	rt = P
	return res
}

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	fmt.Fscan(reader, &n, &m, &q)
	dt = 0
	for i := 1; i <= n; i++ {
		e[i] = e[i][:0]
	}
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		e[x] = append(e[x], y)
		e[y] = append(e[y], x)
	}
	dfs(1, 0)
	dfs2(1, 1)
	segment.build(1, 1, q+1)
	segment.uadd(1, 1, q+1, 2, q+1, m)
	initODT()
	rt = 0
	for i := 1; i <= m; i++ {
		var x int
		fmt.Fscan(reader, &x)
		erase(&stRoot, dfn[x])
		insert(&stRoot, newNode(dfn[x], dfn[x], 1))
		rt = x
	}
	las := 0
	tt := 2
	for i := 1; i <= q; i++ {
		var p int
		fmt.Fscan(reader, &p)
		p = (p-1+las)%n + 1
		dtVal := query(p, tt)
		if dtVal != -1 {
			tt++
			las += dtVal
		}
		fmt.Fprint(writer, las, " ")
	}
	fmt.Fprint(writer, "\n")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	const MaxN = 100005
	e = make([][]int, MaxN)
	s = make([]int, MaxN)
	gs = make([]int, MaxN)
	sz = make([]int, MaxN)
	dfn = make([]int, MaxN)
	top = make([]int, MaxN)
	ft = make([]int, MaxN)
	dep = make([]int, MaxN)
	segment = NewSeg(MaxN)
	var T int
	if _, err := fmt.Fscan(reader, &T); err != nil {
		return
	}
	for i := 1; i <= T; i++ {
		solve(reader, writer)
	}
}

/*
This algorithm integrates Heavy-Light Decomposition (HLD) with an Old Driver Tree (Chtholly Tree) interval tracking
framework and a lazy-propagated Segment Tree to simulate, update, and query path color spans dynamically over a tree
topology. HLD maps tree paths into flat, continuous segments based on heavy-edge metrics, transforming tree path
updates into simple range manipulations. The Chtholly Tree tracks clusters of identical colors across consecutive node
intervals using a balanced tree layout, supporting quick boundary splitting and neighborhood color updates. When resolving
a query path, intervals with distinct color profiles are extracted, intersected, and used to lazily update a global
minimum-tracking Segment Tree that maintains dynamic allocation caps. The total Time Complexity per query sequence scales
as O((N + Q) log^2 N) due to nested HLD path hops and interval updates, while the overall Space Complexity is O(N + Q) to
hold tree decomposition state attributes, node lookup indexes, and active segment blocks.
*/
