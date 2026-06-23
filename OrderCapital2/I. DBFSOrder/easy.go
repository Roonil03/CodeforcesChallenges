package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

var (
	line    [3000][]int
	s       string
	dp      [4][3000][3001]int64
	dpc     [4][3000][3001]int64
	tdp     [4][3001]int64
	tdpc    [4][3001]int64
	dep     [3000]int
	vis1    [3000]int
	vislink [3000]int
)

func merge(u, v, pu, pv, pt int) {
	for i := 0; i <= dep[u]; i++ {
		for j := 0; j <= dep[v]; j++ {
			mx := max(i, j)
			tdp[pt][mx] = (tdp[pt][mx] + dp[pu][u][i]*dp[pv][v][j]) % mod
			if i <= j {
				tdpc[pt][j] = (tdpc[pt][j] + dp[pu][u][i]*dpc[pv][v][j]) % mod
			}
			if i > j {
				tdpc[pt][i] = (tdpc[pt][i] + dpc[pu][u][i]*dp[pv][v][j]) % mod
			}
		}
	}
}

func dfs(t int) {
	dp[0][t][0] = 1
	dep[t] = 0
	vis1[t] = 0
	vislink[t] = 0
	r0 := 0
	for _, child := range line[t] {
		if s[child] == '0' {
			r0++
		}
	}
	for _, child := range line[t] {
		dfs(child)
		if s[child] == '0' {
			r0--
		}
		merge(t, child, 0, 0, 0)
		if s[child] == '1' {
			merge(t, child, 3, 1, 2)
			merge(t, child, 3, 1, 3)
		} else {
			merge(t, child, 2, 1, 2)
			merge(t, child, 3, 1, 3)
		}
		if vis1[t] == 0 && r0 == 0 && vislink[child] != 0 {
			if vislink[t] == 0 {
				vislink[t] = 1
				for i := 0; i <= dep[child]; i++ {
					tdp[2][i] = (tdp[2][i] + dp[2][child][i]) % mod
					tdpc[2][i] = (tdpc[2][i] + dpc[2][child][i]) % mod
					tdp[3][i] = (tdp[3][i] + dp[3][child][i]) % mod
					tdpc[3][i] = (tdpc[3][i] + dpc[3][child][i]) % mod
				}
			} else {
				for i := 0; i <= dep[child]; i++ {
					tdp[2][i] = (tdp[2][i] + dp[2][child][i]) % mod
					tdpc[2][i] = (tdpc[2][i] + dpc[2][child][i]) % mod
					tdp[3][i] = (tdp[3][i] + dp[2][child][i]) % mod
					tdpc[3][i] = (tdpc[3][i] + dpc[2][child][i]) % mod
					if s[child] == '1' {
						tdp[2][i] = (tdp[2][i] + dp[1][child][i]) % mod
						tdpc[2][i] = (tdpc[2][i] + dpc[1][child][i]) % mod
						tdp[3][i] = (tdp[3][i] + dp[1][child][i]) % mod
						tdpc[3][i] = (tdpc[3][i] + dpc[1][child][i]) % mod
					}
				}
			}
		}
		vis1[t] |= vis1[child]
		dep[t] = max(dep[t], dep[child])
		for p := 0; p < 4; p++ {
			for i := 0; i <= dep[t]; i++ {
				dp[p][t][i] = tdp[p][i]
				dpc[p][t][i] = tdpc[p][i]
				tdp[p][i] = 0
				tdpc[p][i] = 0
			}
		}
	}
	if s[t] == '0' {
		dep[t]++
		return
	}
	for i := 0; i <= dep[t]; i++ {
		tdp[0][i+1] = (tdp[0][i+1] + dp[0][t][i]) % mod
		tdpc[0][i+1] = (tdpc[0][i+1] + dpc[0][t][i]) % mod
	}
	var vl int64 = 0
	for i := 0; i <= dep[t]; i++ {
		vl = (vl + dp[2][t][i]) % mod
		vl = (vl + mod - dpc[2][t][i]) % mod
	}
	tdpc[0][1] = (tdpc[0][1] + vl) % mod
	dep[t]++
	vislink[t] = 1
	if s[t] == '1' {
		vis1[t] = 1
		for i := 0; i <= dep[t]; i++ {
			dp[0][t][i] = tdp[0][i]
			dpc[0][t][i] = tdpc[0][i]
			dp[1][t][i] = tdp[0][i]
			dpc[1][t][i] = tdpc[0][i]
			dp[2][t][i] = 0
			dpc[2][t][i] = 0
			dp[3][t][i] = 0
			dpc[3][t][i] = 0
		}
	} else {
		for i := 0; i <= dep[t]; i++ {
			dp[0][t][i] = (dp[0][t][i] + tdp[0][i]) % mod
			dpc[0][t][i] = (dpc[0][t][i] + tdpc[0][i]) % mod
			dp[1][t][i] = tdp[0][i]
			dpc[1][t][i] = tdpc[0][i]
			dp[3][t][i] = dp[2][t][i]
			dpc[3][t][i] = dpc[2][t][i]
			dp[2][t][i] = (dp[2][t][i] + tdp[0][i]) % mod
			dpc[2][t][i] = (dpc[2][t][i] + tdpc[0][i]) % mod
		}
	}
	for i := 0; i <= dep[t]; i++ {
		tdp[0][i] = 0
		tdpc[0][i] = 0
	}
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
		var n int
		var str string
		fmt.Fscan(reader, &n, &str)
		s = "1" + str
		for i := 0; i < n; i++ {
			var c int
			fmt.Fscan(reader, &c)
			line[i] = make([]int, c)
			for j := 0; j < c; j++ {
				fmt.Fscan(reader, &line[i][j])
				line[i][j]--
			}
		}
		for p := 0; p < 4; p++ {
			for i := 0; i < n; i++ {
				for j := 0; j <= n; j++ {
					dp[p][i][j] = 0
					dpc[p][i][j] = 0
				}
			}
		}
		dfs(0)
		var ans int64 = 0
		for i := 0; i <= dep[0]; i++ {
			ans = (ans + dp[0][0][i]) % mod
			ans = (ans + mod - dpc[0][0][i]) % mod
		}
		fmt.Fprintln(writer, ans)
		for i := 0; i < n; i++ {
			line[i] = nil
		}
	}
}

/*
This algorithm evaluates the number of valid distinct tree traversal configurations (DFS/BFS combinations matching
sequence colors) via a sophisticated Tree Dynamic Programming approach with state-convolution style merging. The DP
tracks combinations across 4 separate context states ('p') where the state table 'dp[p][t][i]' captures the number
of valid sequences for a node 't' bounded by an effective configuration depth 'i', while 'dpc' processes identical
sub-combinations to implement the inclusion-exclusion rule for over-counted overlapping states. During the post-order
traversal, a node's table iteratively absorbs the tables of its ordered children using bounded depth metrics, adjusting
parameters whenever a node with a fixed token color restriction is encountered. The Time Complexity is O(N^2) per
testcase because the combinatoric nested loop merges are bounded by tree height/depth pairs which sums cleanly to
O(N^2) on skew trees and behaves linearly on well-balanced inputs. The Space Complexity is O(N^2) due to the fixed
allocation arrays containing maximum coordinates corresponding to both node limits and sequence depth combinations
up to the maximum constraints.
*/
