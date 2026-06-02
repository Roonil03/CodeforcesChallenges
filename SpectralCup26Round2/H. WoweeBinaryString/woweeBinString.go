package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	b := make([]byte, 1024*1024)
	s.Buffer(b, 1024*1024*10)
	nx := func() []byte { s.Scan(); return s.Bytes() }
	if !s.Scan() {
		return
	}
	t := 0
	for _, c := range s.Bytes() {
		t = t*10 + int(c-'0')
	}
	for tc := 0; tc < t; tc++ {
		nx()
		str := nx()
		n := len(str)
		nxt := make([]int, n+1)
		nxt[n] = n
		for i := n - 1; i >= 0; i-- {
			if str[i] == '?' {
				nxt[i] = i
			} else {
				nxt[i] = nxt[i+1]
			}
		}
		pre := make([]int, n+1)
		for i := 0; i < n; i++ {
			v := 0
			if str[i] == '1' {
				v = 1
			}
			pre[i+1] = pre[i] ^ v
		}
		end := make([]bool, n+1)
		for i := 0; i <= n; i++ {
			end[i] = nxt[i] < n || pre[i] == pre[n]
		}
		tr := make([][2][2]int, n+1)
		for i := 0; i <= n; i++ {
			tr[i][0][0], tr[i][0][1] = -1, -1
			tr[i][1][0], tr[i][1][1] = -1, -1
			for c := 0; c < 2; c++ {
				for j := i; j < n; j++ {
					if str[j] == '?' || str[j] == byte('0'+c) {
						if nxt[i] < j || pre[i] == pre[j] {
							if tr[i][c][0] == -1 {
								tr[i][c][0] = j + 1
							} else {
								k := tr[i][c][0]
								if nxt[k] >= j+1 && pre[j+1] != pre[k] {
									tr[i][c][1] = j + 1
									break
								}
							}
						}
					}
				}
			}
		}
		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		dp[0][0] = 1
		ans := 0
		for i := 0; i <= n; i++ {
			for j := i; j <= n; j++ {
				v := dp[i][j]
				if v == 0 {
					continue
				}
				if end[i] || end[j] {
					ans = (ans + v) % 998244353
				}
				for c := 0; c < 2; c++ {
					a := []int{tr[i][c][0], tr[i][c][1], tr[j][c][0], tr[j][c][1]}
					slices.Sort(a)
					x, y := -1, -1
					for _, z := range a {
						if z == -1 {
							continue
						}
						if x == -1 {
							x = z
						} else if nxt[x] >= z && pre[x] != pre[z] {
							y = z
							break
						}
					}
					if y != -1 {
						dp[x][y] = (dp[x][y] + v) % 998244353
					} else if x != -1 {
						dp[x][x] = (dp[x][x] + v) % 998244353
					}
				}
			}
		}
		fmt.Println(ans)
	}
}

/*
Explanation of the Logic:

1. Problem Translation:
The algorithm counts valid configurations of a string with wildcards ('?') where we append elements to match conditions, avoiding overcounting via a Deterministic Finite Automaton (DFA) over prefixes. It systematically identifies equivalence classes of string segments to track valid "frontiers" (how far we have successfully matched the string conditions).

2. Precomputation Arrays:
- `nxt`: Tracks the immediate next index containing a '?' wildcard. This allows quick checks to see if any flexible character insertions exist in a block.
- `pre`: A prefix XOR sum array of the '1' characters in the string. This tracks parity/alternating logic constraints dictated by the problem's criteria.
- `end`: A boolean sequence denoting whether state `i` is a valid terminal/accepting state (either it has '?' wildcard support ahead, or its parity matches the final parity).

3. Transition Logic (`tr`):
For any state `i` and appending character `c` (0 or 1), the problem calculates the earliest valid index mappings. Because strings can match in overlapping ways, we must keep track of up to two minimal non-redundant forward extensions (`tr[i][c][0]` and `tr[i][c][1]`) to maintain the exact bounds of our search space without over-stepping.

4. 2D State Dynamic Programming:
Since overlapping matches can occur from ambiguous wildcards, we simulate the DFA transitions over a pair of pointers `(i, j)`.
- `dp[i][j]` stores the modulo configurations where the two furthest valid tracked prefix lengths are `i` and `j`.
- From `(i, j)`, branching with character `c`, we look at all four possible state projections.
- We sort the valid reachable projections and merge them greedily. If they overlap or represent the same equivalence state (checked by the `nxt` and `pre` XOR parity rule), we consolidate them into `(x, x)`. If they split, they form a new bounding box `(x, y)`.
- If either frontier pointer is marked as an `end` state, the permutation is valid, and its DP value is aggregated into the answer.

This approach compresses exponential wildcard matching into polynomial DP by strictly tracking only the maximal minimal matching prefixes.
*/
