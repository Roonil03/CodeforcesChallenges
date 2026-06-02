package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Buffer(make([]byte, 1024*1024), 1024*1024*10)
	nx := func() int {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		return v
	}
	if !s.Scan() {
		return
	}
	t, _ := strconv.Atoi(s.Text())
	for tc := 0; tc < t; tc++ {
		n, k := nx(), nx()-1
		w := make([]int64, n)
		for i := 0; i < n; i++ {
			w[i] = int64(nx())
		}
		tr := make([][]int, n)
		for i := 0; i < n-1; i++ {
			u, v := nx()-1, nx()-1
			tr[u] = append(tr[u], v)
			tr[v] = append(tr[v], u)
		}
		ch := make([][]int, n)
		sub := make([]int, n)
		var b func(int, int)
		b = func(v, p int) {
			sub[v] = 1
			for _, c := range tr[v] {
				if c != p {
					ch[v] = append(ch[v], c)
					b(c, v)
					sub[v] += sub[c]
				}
			}
		}
		b(0, -1)
		dp := make([][]int64, n)
		pt := make([][]int64, n)
		bfs := make([][][]int64, n)
		sm := func(a, b []int64) []int64 {
			c := make([]int64, len(a)+len(b)-1)
			for i := range c {
				c[i] = -1e18
			}
			for i, va := range a {
				if va == -1e18 {
					continue
				}
				for j, vb := range b {
					if vb == -1e18 {
						continue
					}
					if va+vb > c[i+j] {
						c[i+j] = va + vb
					}
				}
			}
			return c
		}
		adv := func(a []int64, val int64) []int64 {
			b := make([]int64, len(a)+1)
			for i := range b {
				b[i] = -1e18
			}
			for i, va := range a {
				if va == -1e18 {
					continue
				}
				if va > b[i] {
					b[i] = va
				}
				if va+val*int64(i+1) > b[i+1] {
					b[i+1] = va + val*int64(i+1)
				}
			}
			return b
		}
		var d1 func(int)
		d1 = func(v int) {
			for _, c := range ch[v] {
				d1(c)
			}
			bfs[v] = make([][]int64, len(ch[v])+1)
			bfs[v][0] = []int64{0}
			for i, c := range ch[v] {
				bfs[v][i+1] = sm(bfs[v][i], dp[c])
			}
			pt[v] = bfs[v][len(ch[v])]
			dp[v] = adv(pt[v], w[v])
		}
		d1(0)
		ans := make([]int64, n)
		for i := range ans {
			ans[i] = -1e18
		}
		var d2 func(int, []int64)
		d2 = func(v int, z []int64) {
			for i, vz := range z {
				if vz == -1e18 || pt[v][i] == -1e18 {
					continue
				}
				if pt[v][i]+vz > ans[v] {
					ans[v] = pt[v][i] + vz
				}
			}
			for id := len(ch[v]) - 1; id >= 0; id-- {
				x, c := bfs[v][id], ch[v][id]
				p, q := len(x)-1, sub[c]
				y := make([]int64, q+1)
				for i := range y {
					y[i] = -1e18
				}
				for i := 0; i <= q; i++ {
					for j := 0; j <= p; j++ {
						if x[j] == -1e18 || z[i+j] == -1e18 {
							continue
						}
						if x[j]+z[i+j] > y[i] {
							y[i] = x[j] + z[i+j]
						}
					}
				}
				for i := 0; i < q; i++ {
					if y[i+1] != -1e18 {
						if val := y[i+1] + int64(k-i)*w[v]; val > y[i] {
							y[i] = val
						}
					}
				}
				y = y[:q]
				d2(c, y)
				nz := make([]int64, p+1)
				for i := range nz {
					nz[i] = -1e18
				}
				for i := 0; i <= p; i++ {
					mv := int64(-1e18)
					for j := 0; j <= q; j++ {
						if dp[c][j] == -1e18 || z[i+j] == -1e18 {
							continue
						}
						if dp[c][j]+z[i+j] > mv {
							mv = dp[c][j] + z[i+j]
						}
					}
					nz[i] = mv
				}
				z = nz
			}
		}
		ini := make([]int64, n)
		for i := range ini {
			ini[i] = -1e18
		}
		if k < n {
			ini[k] = 0
		}
		d2(0, ini)
		for i, v := range ans {
			fmt.Print(v + w[i]*int64(k+1))
			if i == n-1 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}

/*
Explanation of the Logic:

1. Algorithm Family:
The algorithm uses the "Rerooting Tree DP" (or All-Roots DP) combined with a Knapsack-like convolution. The time complexity stays strictly bounded at O(N^2) instead of O(N^3) by leveraging the typical property of merging subtree sizes—each pair of nodes is matched against each other exactly once across the entire tree.

2. Bottom-Up Pass (dfs1):
The tree is rooted arbitrarily at node 0.
We compute a state `dp[v][i]`, which tracks the maximum configuration value obtainable within the subtree of `v`, given `i` remaining operational capacities.
The convolution step `simple(a, b)` merges the states of two branches, acting exactly like polynomial multiplication but substituting addition for multiplication and the `max` function for addition.
Importantly, to facilitate the rerooting later, we store the prefix-aggregated DP states in `bfs[v]`.

3. Top-Down Rerooting Pass (dfs2):
This pass moves downwards, propagating the DP state representing everything "outside" of the current subtree `v` (passed as the array `z`).
- At node `v`, the local `ans[v]` is definitively solved by combining the precalculated aggregated subtree DP `pt[v]` and the outside context `z`.
- To branch out into a child `c`, we must exclude `c`'s contribution from the full context. Because we iterate the children in reverse, we can build the missing suffix on-the-fly (`nz`), effectively merging it with the precalculated prefix `bfs[v][id]` and external context `z` to form the true state `y` representing the whole tree excluding the subtree of `c`. This state `y` is then recursively passed down into `c`.
*/
