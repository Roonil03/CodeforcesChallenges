package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	b, _ := io.ReadAll(os.Stdin)
	p := 0
	I := func() int {
		for p < len(b) && (b[p] < '0' || b[p] > '9') {
			p++
		}
		if p >= len(b) {
			return 0
		}
		r := 0
		for p < len(b) && b[p] >= '0' && b[p] <= '9' {
			r = r*10 + int(b[p]-'0')
			p++
		}
		return r
	}
	t := I()
	for i := 0; i < t; i++ {
		n, k := I(), I()
		A := make([]int64, n)
		M, S := int64(0), int64(0)
		idx := -1
		for j := 0; j < n; j++ {
			A[j] = int64(I())
			S += A[j]
			if A[j] > M {
				M = A[j]
				idx = j
			}
		}
		if k == 1 {
			fmt.Println(S)
			continue
		}
		B := make([]int64, 0, n-1)
		for j := 0; j < n; j++ {
			if j != idx {
				B = append(B, A[j])
			}
		}
		m := n - 1
		K := int64(k)
		C := func(W int64) bool {
			if W <= 0 {
				return true
			}
			d := make([]int64, 1<<m)
			for j := range d {
				d[j] = -1
			}
			d[0] = 0
			for Msk := 0; Msk < 1<<m; Msk++ {
				v := d[Msk]
				if v == -1 {
					continue
				}
				c, r := v>>40, v&((1<<40)-1)
				for j := 0; j < m; j++ {
					if Msk&(1<<j) == 0 {
						nc, nr := c, r+B[j]
						if nr >= W {
							nc++
							nr = 0
							if nc >= K {
								return true
							}
						}
						nv := (nc << 40) | nr
						if nv > d[Msk|(1<<j)] {
							d[Msk|(1<<j)] = nv
						}
					}
				}
			}
			return false
		}
		L, H := M, S
		ans := L
		for L <= H {
			mid := L + (H-L)/2
			if C(mid - M) {
				ans = mid
				L = mid + 1
			} else {
				H = mid - 1
			}
		}
		fmt.Println(ans)
	}
}

/*
Explanation of the Algorithm:

1. Core Insight:
The greedy process "add to the minimum bin" allows us to build target bins as long as we carefully reorder elements. Specifically, a target bin can achieve a maximum sum V ending with some element x if and only if the array's remaining elements can form k subsets, each with a sum >= V - x.

2. Max Element is Always Optimal:
To maximize the target sum V, the mathematically optimal choice for the final element x is the absolute maximum element of the array. Let x = max(A). Swapping x with any smaller element y in the theoretical subsets maintains or strictly lowers the achievable max sum. Thus, we ONLY ever need to consider x = max(A), reducing the problem from checking all x to just a single predefined x.

3. Binary Search:
Since the achievable maximum sum is monotonic, we binary search the answer V within the range [max(A), sum(A)]. For a fixed guess V, the target sum for every subset becomes W = V - x.

4. SOS DP (Sum Over Subsets):
For a fixed W, we check if the remaining array B = A \ {x} can be partitioned into k subsets each >= W. This is checked using an O(N * 2^N) Dynamic Programming over bitmasks.
- dp[mask] packs two values into an int64: the maximum number of completed subsets (shifted by 40 bits) and the maximum sum of the current incomplete subset (lower 40 bits).
- We iterate over unvisited elements and greedily "seal" a subset the moment its sum reaches >= W. This greedy immediate sealing over all permutations perfectly yields the optimal maximum subsets.
- If we ever reach k completed subsets, we immediately return true (early exit).
Because the problem guarantees sum(2^N) <= 2^18, the total DP state transitions across all binary searches for all testcases are heavily bounded, executing efficiently within the strict 1.5s limit.
*/
