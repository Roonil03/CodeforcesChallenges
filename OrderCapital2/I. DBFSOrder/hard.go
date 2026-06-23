package main

import (
	"bufio"
	"fmt"
	"os"
)

func modInvInRange(a, m int) int {
	x, y := a, m
	vx, vy := 1, 0
	swap := false
	for x > 0 {
		k := y / x
		y %= x
		vy += k * vx
		x, y = y, x
		vx, vy = vy, vx
		swap = !swap
	}
	if swap {
		return vy
	}
	return m - vy
}

const MOD = 1000000007

type num int

func (a num) add(b num) num {
	v := a + b
	if v >= MOD {
		v -= MOD
	}
	return v
}

func (a num) sub(b num) num {
	v := a - b
	if v < 0 {
		v += MOD
	}
	return v
}

func (a num) mul(b num) num {
	return num((int64(a) * int64(b)) % MOD)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var T int
	if _, err := fmt.Fscan(reader, &T); err != nil {
		return
	}
	for T > 0 {
		T--
		var N int
		var S string
		fmt.Fscan(reader, &N, &S)
		S = "0" + S
		ch := make([][]int, N)
		for i := 0; i < N; i++ {
			var l int
			fmt.Fscan(reader, &l)
			ch[i] = make([]int, l)
			for j := 0; j < l; j++ {
				fmt.Fscan(reader, &ch[i][j])
				ch[i][j]--
			}
		}
		dp := make([][][6]num, N)
		var dfs func(cur int)
		dfs = func(cur int) {
			dp[cur] = [][6]num{{1, 0, 0, 0, 1, 0}}
			for _, nxt := range ch[cur] {
				dfs(nxt)
				maxLen := len(dp[cur])
				if len(dp[nxt])+1 > maxLen {
					maxLen = len(dp[nxt]) + 1
				}
				ndp := make([][6]num, maxLen)
				for d0 := 0; d0 < len(dp[cur]); d0++ {
					for d1 := 0; d1 < len(dp[nxt]); d1++ {
						for _, curHasActive := range []int{0, 1, 2} {
							for _, curHasNxt := range []bool{false, true} {
								for _, nxtHasActive := range []int{0, 1, 2} {
									for _, nxtHasNxt := range []bool{false, true} {
										curIdx := curHasActive * 2
										if curHasNxt {
											curIdx++
										}
										nxtIdx := nxtHasActive * 2
										if nxtHasNxt {
											nxtIdx++
										}
										vCur := dp[cur][d0][curIdx]
										vNxt := dp[nxt][d1][nxtIdx]
										if vCur == 0 || vNxt == 0 {
											continue
										}
										v := vCur.mul(vNxt)
										for _, e := range []bool{false, true} {
											var charE byte
											if e {
												charE = '1'
											} else {
												charE = '0'
											}
											if S[nxt] != '?' && S[nxt] != charE {
												continue
											}
											var nNxtHasActive int
											var nNxtHasNxt bool
											var nd1 int
											if e && nxtHasActive != 0 {
												nNxtHasActive = 0
												nNxtHasNxt = true
												nd1 = 1
											} else {
												nNxtHasActive = nxtHasActive
												nNxtHasNxt = nxtHasNxt
												nd1 = d1
												if e {
													nd1++
												}
											}
											if curHasActive != 0 && nNxtHasActive != 0 {
												continue
											}
											if curHasNxt && nNxtHasNxt {
												continue
											}
											if curHasActive != 0 && !e {
												continue
											}
											if nNxtHasNxt && d0 > nd1 {
												continue
											}
											if curHasNxt && d0 <= nd1 {
												continue
											}
											nd := d0
											if nd1 > nd {
												nd = nd1
											}
											nHasNxt := curHasNxt || nNxtHasNxt
											if nNxtHasActive != 0 && d0 > 0 {
												continue
											}
											if curHasActive == 2 && nxtHasActive == 2 {
												continue
											}
											var nHasNxtIdx int
											if nHasNxt {
												nHasNxtIdx = 1
											} else {
												nHasNxtIdx = 0
											}
											if curHasActive == 2 || nxtHasActive == 2 {
												if S[nxt] == '0' {
													if nxtHasActive != 2 {
														continue
													}
													ndp[nd][4+nHasNxtIdx] = ndp[nd][4+nHasNxtIdx].add(v)
												} else if S[nxt] == '1' {
													if curHasActive != 2 {
														continue
													}
													if nxtHasActive == 2 {
														continue
													}
													ndp[nd][2+nHasNxtIdx] = ndp[nd][2+nHasNxtIdx].add(v)
												} else if S[nxt] == '?' {
													if curHasActive != 2 {
														continue
													}
													if nxtHasActive == 2 {
														continue
													}
													if !e {
														continue
													}
													ndp[nd][4+nHasNxtIdx] = ndp[nd][4+nHasNxtIdx].add(v)
												}
												continue
											}
											var activeVal int
											if curHasActive != 0 || nNxtHasActive != 0 {
												activeVal = 1
											} else {
												activeVal = 0
											}
											ndp[nd][activeVal*2+nHasNxtIdx] = ndp[nd][activeVal*2+nHasNxtIdx].add(v)
											if curHasActive == 0 && e && d0 == 0 {
												ndp[nd][2+nHasNxtIdx] = ndp[nd][2+nHasNxtIdx].sub(v)
											}
										}
									}
								}
							}
						}
					}
				}
				dp[cur] = ndp
			}
		}
		dfs(0)
		var ans num = 0
		for _, v := range dp[0] {
			ans = ans.add(v[0]).add(v[1]).add(v[2]).add(v[3])
		}
		fmt.Fprintln(writer, ans)
	}
}

/*
This algorithm implements a specialized Tree Dynamic Programming model to count valid structural partition
combinations under dynamic token constraints ('0', '1', and '?'). The DP state tracks combinations across
6 specific layout modes packed within a fixed dimension array per subtree depth configuration. These modes
map pairs of properties tracking whether there is an active partition cut block, whether there is a look-ahead
boundary constraint, and their deep-level intersections within the current node's subtree. During the post-order
traversal phase, the algorithm convolves and updates variable-sized depth records across neighboring subtrees,
filtering out invalid arrangements that violate string token matching rules or threshold inequalities. The Time
Complexity is O(N^2) in the worst case (such as a skewed linear path graph) because it performs an iterative
sequence of nested polynomial multi-product updates matching subtree depth boundaries. The Space Complexity is
O(N^2) due to the dynamic allocation of the variable-length DP multi-state arrays for each vertex across the
tree representation.
*/
