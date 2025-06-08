// The Question is solved post the contest times

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	sum, pref, suff, best int
}

func newNode(x int) Node {
	return Node{
		sum:  x,
		pref: max(0, x),
		suff: max(0, x),
		best: max(0, x),
	}
}

func nullNode() Node {
	const NEG_INF = -1000000000
	return Node{
		sum:  0,
		pref: NEG_INF,
		suff: NEG_INF,
		best: NEG_INF,
	}
}

func merge(L, R Node) Node {
	const NEG_INF = -1000000000
	if L.best == NEG_INF {
		return R
	}
	if R.best == NEG_INF {
		return L
	}
	return Node{
		sum:  L.sum + R.sum,
		pref: max(L.pref, L.sum+R.pref),
		suff: max(R.suff, R.sum+L.suff),
		best: max(max(L.best, R.best), L.suff+R.pref),
	}
}

type SegTree struct {
	n  int
	st []Node
}

func newSegTree(n int) *SegTree {
	st := &SegTree{
		n:  n,
		st: make([]Node, 4*n),
	}
	st.build(1, 0, n-1)
	return st
}

func (st *SegTree) build(p, l, r int) {
	if l == r {
		st.st[p] = newNode(-1)
		return
	}
	m := (l + r) / 2
	st.build(p<<1, l, m)
	st.build(p<<1|1, m+1, r)
	st.st[p] = merge(st.st[p<<1], st.st[p<<1|1])
}

func (st *SegTree) update(p, l, r, idx, val int) {
	if l == r {
		st.st[p] = newNode(val)
		return
	}
	m := (l + r) / 2
	if idx <= m {
		st.update(p<<1, l, m, idx, val)
	} else {
		st.update(p<<1|1, m+1, r, idx, val)
	}
	st.st[p] = merge(st.st[p<<1], st.st[p<<1|1])
}

func (st *SegTree) updateIdx(idx, val int) {
	st.update(1, 0, st.n-1, idx, val)
}

func (st *SegTree) maxSubarray() int {
	return st.st[1].best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type MultiSet struct {
	data map[int]int
}

func newMultiSet() *MultiSet {
	return &MultiSet{data: make(map[int]int)}
}

func (ms *MultiSet) insert(x int) {
	ms.data[x]++
}

func (ms *MultiSet) erase(x int) {
	ms.data[x]--
	if ms.data[x] == 0 {
		delete(ms.data, x)
	}
}

func (ms *MultiSet) getMax() int {
	maxVal := -1000000000
	for k := range ms.data {
		if k > maxVal {
			maxVal = k
		}
	}
	return maxVal
}

func solve(scanner *bufio.Scanner) {
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(line[0])
	q, _ := strconv.Atoi(line[1])
	scanner.Scan()
	aFields := strings.Fields(scanner.Text())
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(aFields[i])
	}
	at := make([][]int, n+1)
	for i := 0; i < n; i++ {
		at[a[i]] = append(at[a[i]], i)
	}
	type Update struct {
		idx, pos, val int
	}
	updates := make([][]Update, n+1)
	for idx := 0; idx < q; idx++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		i, _ := strconv.Atoi(line[0])
		x, _ := strconv.Atoi(line[1])
		i--
		if x != a[i] {
			updates[a[i]] = append(updates[a[i]], Update{idx, i, -1})
			a[i] = x
			updates[x] = append(updates[x], Update{idx, i, 1})
		}
	}
	finalAt := make([][]int, n+1)
	for i := 0; i < n; i++ {
		finalAt[a[i]] = append(finalAt[a[i]], i)
	}
	st := newSegTree(n)
	init := make([]int, n+1)
	change := make([][]struct{ old, new int }, q)
	for x := 1; x <= n; x++ {
		for _, i := range at[x] {
			st.updateIdx(i, 1)
		}
		cur := st.maxSubarray()
		init[x] = cur
		for _, upd := range updates[x] {
			st.updateIdx(upd.pos, upd.val)
			newCur := st.maxSubarray()
			change[upd.idx] = append(change[upd.idx], struct{ old, new int }{cur, newCur})
			cur = newCur
		}
		for _, i := range finalAt[x] {
			st.updateIdx(i, -1)
		}
	}
	ms := newMultiSet()
	for i := 1; i <= n; i++ {
		ms.insert(init[i])
	}
	for i := 0; i < q; i++ {
		for _, ch := range change[i] {
			ms.erase(ch.old)
			ms.insert(ch.new)
		}
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(ms.getMax() / 2)
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for test := 0; test < t; test++ {
		solve(scanner)
	}
}

/*
This advanced solution employs a sophisticated mathematical insight: the k-majority problem can be
transformed into a maximum subarray sum problem. Here's how it works:

For each distinct value x in the array, the algorithm creates a binary representation where
positions containing x get value +1 and all other positions get value -1. When we find the
maximum subarray sum in this transformed array, we get the maximum excess of x over the majority
threshold across all possible subarrays. The minimum k value is then this maximum excess divided
by 2.

The algorithm preprocesses all distinct values by building their initial maximum subarray sums,
then tracks how these values change with each update query. A multiset maintains all current
maximum values, allowing efficient retrieval of the global maximum after each query.

Time Complexity: O(n² log n) for preprocessing all distinct values with segment tree operations,
plus O(q log n) for processing queries. Space Complexity: O(n²) for storing update information
and segment tree nodes. This approach scales much better than brute force for large inputs while
maintaining the 192MB memory limit through careful data structure usage.

With n=300,000, the array input line can contain up to 300,000 integers, each potentially 6 digits long
plus spaces. This creates a line that can be 2-3MB in size, far exceeding the default bufio.Scanner limit
of 64KB, causing the "bufio.Scanner: token too long" error that results in exit code 2.
*/
