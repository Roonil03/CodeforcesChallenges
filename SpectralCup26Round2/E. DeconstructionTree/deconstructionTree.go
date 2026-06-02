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
	b := make([]byte, 1024*1024)
	s.Buffer(b, 1024*1024*10)
	nx := func() int { s.Scan(); v, _ := strconv.Atoi(s.Text()); return v }
	if !s.Scan() {
		return
	}
	t, _ := strconv.Atoi(s.Text())
	for tc := 0; tc < t; tc++ {
		n := nx()
		a := make([][]int, n+1)
		for i := 0; i < n-1; i++ {
			u, v := nx(), nx()
			a[u] = append(a[u], v)
			a[v] = append(a[v], u)
		}
		if len(a[n]) <= 1 {
			fmt.Println(1)
			continue
		}
		m, r := make([]int, n+1), make([]int, n+1)
		l := make([]bool, n+1)
		var d func(u, p int)
		d = func(u, p int) {
			m[u] = u
			c := 0
			for _, v := range a[u] {
				if v != p {
					c++
					d(v, u)
					if m[v] > m[u] {
						m[u] = m[v]
					}
					if m[v] > r[u] {
						r[u] = m[v]
					}
				}
			}
			if c == 0 {
				l[u] = true
			}
		}
		d(n, 0)
		M1, M2 := 0, 0
		for _, v := range a[n] {
			x := m[v]
			if x > M1 {
				M2 = M1
				M1 = x
			} else if x > M2 {
				M2 = x
			}
		}
		I := 0
		for i := 1; i < n; i++ {
			if l[i] && i > I {
				I = i
			}
		}
		f, p := make([]int, n+1), make([]int, n+1)
		if I > 0 {
			f[I] = 1
		}
		for i := 1; i <= I; i++ {
			p[i] = (p[i-1] + f[i]) % 998244353
		}
		for v := I + 1; v < n; v++ {
			if r[v] < v-1 {
				f[v] = (p[v-1] - p[r[v]] + 998244353) % 998244353
			}
			p[v] = (p[v-1] + f[v]) % 998244353
		}
		A, S := 0, M2+1
		if S < 1 {
			S = 1
		}
		for w := S; w < n; w++ {
			A = (A + f[w]) % 998244353
		}
		fmt.Println(A)
	}
}

/*
Explanation of the Bug & Algorithm:

1. The Hidden Edge Case (The Bug):
The DP formulation is extremely robust, but it fundamentally assumes that the initial max leaf is `< n`.
If `n` is a leaf in the initial unrooted tree, then `n` is inherently the absolute maximum index. Because `n` is a leaf, it will immediately be designated as `x` (the max leaf). Since the rules dictate we must remove a leaf *other* than `x`, `n` is never removed. It will remain a leaf, and thus remain `x` for every single operation until the tree is gone.
Therefore, if `n` is a leaf, no other node can ever become the max leaf, meaning the set `S` will ALWAYS be exactly `{n}`. If `n` has only 1 neighbor (degree 1), we can instantly output 1 and skip the DP.

2. Identifying Shields & the Requirement Array:
If `n` is not a leaf, we root the tree at `n`. The problem transforms into finding how many valid strictly increasing sequences of "maximum leaves" we can form.
To expose an internal node `v` and make it a leaf, all of its children in the rooted tree must be completely deleted. To delete these subtrees safely without their elements becoming the new "maximum leaf", we need our current max leaf `u` to act as a strictly larger "shield".
Let `m[c]` be the maximum node value in the subtree of child `c`.
For any node `v`, define `r[v]` = max(m[c]) for all children `c` of `v`. This `r[v]` is the absolute minimum leaf value required to safely clear all of `v`'s children.

3. DP Transitions:
Any valid set `S` corresponds to an increasing sequence of leaves: s_0 < s_1 < ... < s_k.
- `s_0` MUST be `I`, which is the maximum among the initial leaves of the rooted tree (excluding `n`).
- We can transition from a max leaf `u` to expose `v` if and only if `r[v] < u < v`.

4. Termination Condition:
The sequence terminates and lets `n` take over when the current max leaf `w` is large enough to safely clear all the *other* main branches of the root `n`. Let `M1` and `M2` be the largest and second-largest `m[v]` values among the direct branches of the root `n`. To clear the rest of `n`'s branches leaving only the branch with `w`, `w` must act as a shield strictly greater than all other branches (`w > M2`).

5. Complexity:
Let `f[v]` be the number of valid sequences ending exactly at node `v`. We compute `f[v] = sum(f[u])` for all `u` in `[r[v] + 1, v - 1]` using a prefix sum array `p`. Time complexity per testcase resolves efficiently in O(N).
*/
