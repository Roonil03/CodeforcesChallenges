package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const inf int64 = 4e18

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type info struct {
	b1, k1, t1 int64
	b2, k2, t2 int64
}

func empty() info { return info{-inf, 0, inf, -inf, 0, inf} }

func merge(l, r info) info {
	x := info{t1: min(l.t1, r.t1), t2: min(l.t2, r.t2)}
	if l.b1 < r.b1 || (l.b1 == r.b1 && l.k1 < r.k1) {
		l.b1, r.b1, l.k1, r.k1 = r.b1, l.b1, r.k1, l.k1
	}
	if l.b2 < r.b2 || (l.b2 == r.b2 && l.k2 < r.k2) {
		l.b2, r.b2, l.k2, r.k2 = r.b2, l.b2, r.k2, l.k2
	}
	x.b1, x.k1, x.b2, x.k2 = l.b1, l.k1, l.b2, l.k2

	if r.k1 > l.k1 {
		x.t1 = min(x.t1, (l.b1-r.b1+r.k1-l.k1-1)/(r.k1-l.k1))
	}
	if r.k2 > l.k2 {
		x.t2 = min(x.t2, (l.b2-r.b2+r.k2-l.k2-1)/(r.k2-l.k2))
	}
	return x
}

type seg struct {
	n    int
	tree []info
	tag  []int64
}

func newSeg(n int) *seg {
	s := &seg{n: n, tree: make([]info, 4*n), tag: make([]int64, 4*n)}
	for i := range s.tree {
		s.tree[i] = empty()
	}
	return s
}

func (s *seg) apply(p int, v int64) {
	s.tree[p].b1 += s.tree[p].k1 * v
	s.tree[p].b2 += s.tree[p].k2 * v
	s.tree[p].t1 -= v
	s.tree[p].t2 -= v
	s.tag[p] += v
}

func (s *seg) push(p int) {
	if s.tag[p] != 0 {
		s.apply(2*p, s.tag[p])
		s.apply(2*p+1, s.tag[p])
		s.tag[p] = 0
	}
}

func (s *seg) modify(p, l, r, x int, v info) {
	if r-l == 1 {
		s.tree[p] = v
		return
	}
	m := (l + r) / 2
	s.push(p)
	if x < m {
		s.modify(2*p, l, m, x, v)
	} else {
		s.modify(2*p+1, m, r, x, v)
	}
	s.tree[p] = merge(s.tree[2*p], s.tree[2*p+1])
}

func (s *seg) query(p, l, r, x, y int) info {
	if l >= y || r <= x {
		return empty()
	}
	if l >= x && r <= y {
		return s.tree[p]
	}
	m := (l + r) / 2
	s.push(p)
	return merge(s.query(2*p, l, m, x, y), s.query(2*p+1, m, r, x, y))
}

func (s *seg) advance(p, l, r int, v int64) {
	if v < s.tree[p].t1 && v < s.tree[p].t2 {
		s.apply(p, v)
		return
	}
	m := (l + r) / 2
	s.push(p)
	s.advance(2*p, l, m, v)
	s.advance(2*p+1, m, r, v)
	s.tree[p] = merge(s.tree[2*p], s.tree[2*p+1])
}

type pair struct{ tm, i int }

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024*1024), 10*1024*1024)

	nextInt := func() int {
		sc.Scan()
		res, sign := 0, 1
		for i, b := range sc.Bytes() {
			if i == 0 && b == '-' {
				sign = -1
				continue
			}
			res = res*10 + int(b-'0')
		}
		return res * sign
	}

	if !sc.Scan() {
		return
	}
	res, _ := 0, 1
	for _, b := range sc.Bytes() {
		res = res*10 + int(b-'0')
	}
	tCases := res

	for tc := 0; tc < tCases; tc++ {
		n := nextInt()
		k := int64(nextInt())
		x := nextInt() - 1

		h := make([]int, n)
		for i := 0; i < n; i++ {
			h[i] = nextInt()
		}

		d := make([]int, n-1)
		for i := 0; i < n-1; i++ {
			d[i] = nextInt()
		}

		tArr := make([]int, n)
		for i := x - 1; i >= 0; i-- {
			tArr[i] = int(max(int64(tArr[i+1]+1), int64(d[i])))
		}
		for i := x + 1; i < n; i++ {
			tArr[i] = int(max(int64(tArr[i-1]+1), int64(d[i-1])))
		}

		dp := make([]int64, n)
		for i := range dp {
			dp[i] = -inf
		}
		dp[x] = 0

		pre := make([]int64, n+1)
		for i := 0; i < n; i++ {
			pre[i+1] = pre[i] + int64(h[i])
		}

		st := newSeg(n)
		var p []pair
		for i := 0; i < n; i++ {
			if i != x {
				p = append(p, pair{tArr[i] + (i - x), i})
				p = append(p, pair{tArr[i] - (i - x), i})
			}
		}

		sort.Slice(p, func(i, j int) bool {
			a, b := p[i], p[j]
			if a.tm != b.tm {
				return a.tm < b.tm
			}
			fa := (a.i < x) == (a.tm == tArr[a.i]-(a.i-x))
			fb := (b.i < x) == (b.tm == tArr[b.i]-(b.i-x))
			if fa != fb {
				return fa
			}
			return abs(a.i-x) < abs(b.i-x)
		})

		st.modify(1, 0, n, x, info{b1: pre[x], k1: int64(h[x]), t1: inf, b2: -pre[x+1], k2: int64(h[x]), t2: inf})

		lst := 0
		for _, ev := range p {
			tm, i := ev.tm, ev.i
			st.advance(1, 0, n, int64(tm-lst))
			lst = tm

			if i < x {
				if tm == tArr[i]+(i-x) {
					dp[i] = max(dp[i], st.query(1, 0, n, i+1, n).b1-pre[i])
					res := st.query(1, 0, n, i, i+1)
					res.b1 = dp[i] + pre[i]
					res.k1 = int64(h[i])
					st.modify(1, 0, n, i, res)
				} else {
					res := st.query(1, 0, n, i, i+1)
					res.b2 = dp[i] - pre[i+1]
					res.k2 = int64(h[i])
					st.modify(1, 0, n, i, res)
				}
			} else {
				if tm == tArr[i]-(i-x) {
					dp[i] = max(dp[i], st.query(1, 0, n, 0, i).b2+pre[i+1])
					res := st.query(1, 0, n, i, i+1)
					res.b2 = dp[i] - pre[i+1]
					res.k2 = int64(h[i])
					st.modify(1, 0, n, i, res)
				} else {
					res := st.query(1, 0, n, i, i+1)
					res.b1 = dp[i] + pre[i]
					res.k1 = int64(h[i])
					st.modify(1, 0, n, i, res)
				}
			}
		}

		ans := int64(0)
		for i := 0; i < n; i++ {
			if int64(tArr[i]) <= k {
				ans = max(ans, dp[i]+int64(h[i])*(k-int64(tArr[i])))
			}
		}
		fmt.Println(ans)
	}
}

/*
Explanation of the Logic & Algorithm:

1. Core Strategy:
The problem is about maximizing satisfaction by finding an optimal sequence of moving between houses and waiting at specific houses. Because the satisfaction grows linearly over time spent waiting (score = base_score + time_spent * h[i]), the relationships between scores form linear functions of the form `y = kt + b`. To compute the optimal strategy fast, the solution employs a Kinetic Segment Tree (a Li Chao Tree variant).

2. Reachability Array (`tArr`):
We first calculate the absolute earliest day `tArr[i]` we can reach any house `i` if we continuously move outward from the starting house `x`.
- For nodes to the left: `tArr[i] = max(tArr[i+1] + 1, d[i])`
- For nodes to the right: `tArr[i] = max(tArr[i-1] + 1, d[i-1])`

3. The Kinetic Segment Tree:
The segment tree maintains the upper envelope of multiple intersecting linear functions over time.
- `Info` stores two lines `(b1, k1)` and `(b2, k2)` representing paths extending rightwards and leftwards respectively.
- `t1` and `t2` track the earliest time (intersection point) when a sub-optimal line in a child node will overtake the current maximum line.
- When `advance` is called, we globally increment the time variable. If the elapsed time pushes past an intersection `t1` or `t2`, the tree is lazily pushed down to recalculate the exact upper envelope. Otherwise, it updates in `O(1)`.

4. Event Sweep:
We can only change directions efficiently at specific pivotal moments (events). For every house `i != x`, we generate two events: when we reach it going outwards, and when we could theoretically bounce back towards it. We sort these events strictly by time.
- As we iterate through the sorted events, we `advance` the segment tree by the time difference `tm - lst`.
- We query the segment tree for the best possible linear function coming from the previous direction, evaluate its value at the current time to update `dp[i]`, and insert the new line state for node `i` into the segment tree.

5. Final Calculation:
Once all events are processed and `dp` stores the optimal base score for terminating a journey at node `i`, we simply project all valid nodes forward to day `k` using `dp[i] + h[i] * (k - tArr[i])` and find the global maximum. The heavily optimized data structure runs effortlessly inside O(N log N) per test case.
*/
