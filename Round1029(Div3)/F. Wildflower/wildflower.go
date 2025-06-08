// The Question is solved post the contest times

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N   = 200010
	MOD = 1000000007
)

var (
	adj  [N][]int
	lens []int
	pw   [N]int
	lca  int
)

func dfs(u, par, length int) {
	if len(adj[u]) > 2 {
		lca = length
	}
	leaf := true
	for _, v := range adj[u] {
		if v != par {
			dfs(v, u, length+1)
			leaf = false
		}
	}
	if leaf {
		lens = append(lens, length)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(scanner *bufio.Scanner) {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 1; i <= n; i++ {
		adj[i] = adj[i][:0]
	}
	lens = lens[:0]
	lca = -1
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		parts := strings.Split(scanner.Text(), " ")
		u, _ := strconv.Atoi(parts[0])
		v, _ := strconv.Atoi(parts[1])
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	adj[1] = append(adj[1], 0)
	dfs(1, 0, 1)
	if len(lens) > 2 {
		fmt.Println(0)
	} else if len(lens) == 1 {
		fmt.Println(pw[n])
	} else {
		diff := abs(lens[0] - lens[1])
		x := diff + lca
		if diff != 0 {
			fmt.Println((pw[x] + pw[x-1]) % MOD)
		} else {
			fmt.Println((2 * pw[x]) % MOD)
		}
	}
}

func main() {
	pw[0] = 1
	for i := 1; i < N; i++ {
		pw[i] = (pw[i-1] * 2) % MOD
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for t > 0 {
		solve(scanner)
		t--
	}
}

/*
This algorithm solves a tree-based counting problem with the following approach:

Core Logic:
1. Tree Analysis: The algorithm builds a tree from the input and adds a dummy node to the root
(node 1) to ensure it's never considered a leaf

2. Leaf Detection: It performs DFS to find all true leaf nodes and their distances from the root,
while tracking the deepest branching point (node with degree > 2)

3. Case-based Counting: Based on the number of leaves found:
	- More than 2 leaves: Returns 0 (impossible/invalid case)
	- Exactly 1 leaf: Returns 2^n (likely counting all possible subsets)
	- Exactly 2 leaves: Uses a complex formula involving the distance difference between leaves and
	  the branching point depth

Time Complexity: O(n) per test case for the DFS traversal
Space Complexity: O(n) for the adjacency list representation and auxiliary arrays

Algorithm Nature: This appears to be solving a combinatorial problem on trees, possibly counting valid
paths, subsequences, or configurations between leaf nodes. The use of powers of 2 suggests it's counting
binary choices or subsets, while the leaf-based logic indicates the solution depends on the tree's
"endpoints" and their relative positions.

The optimized Go version uses a larger buffer for bufio.Scanner to handle large inputs efficiently and
prevent TLE issues.
*/
