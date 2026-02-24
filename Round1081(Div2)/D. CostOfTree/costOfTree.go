package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, err := in.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

func readInts(n int) []int {
	line := readLine()
	if line == "" {
		return nil
	}
	parts := strings.Fields(line)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func printSlice(arr []int64) {
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

var (
	adj  [][]int
	a    []int
	size []int64
	dist []int64
	h    []int
	f    []int64
	res  []int64
)

func dfs(u, p int) {
	size[u] = int64(a[u-1])
	dist[u] = 0
	h[u] = 0
	var children []int
	for _, v := range adj[u] {
		if v == p {
			continue
		}
		dfs(v, u)
		size[u] += size[v]
		dist[u] += dist[v] + size[v]
		if h[v]+1 > h[u] {
			h[u] = h[v] + 1
		}
		children = append(children, v)
	}
	f[u] = dist[u]
	for _, v := range children {
		current := dist[u] - (dist[v] + size[v]) + f[v] + size[v]
		if current > f[u] {
			f[u] = current
		}
	}
	if len(children) >= 2 {
		maxH1, maxH2 := -1, -1
		for _, v := range children {
			height := h[v]
			if height > maxH1 {
				maxH2 = maxH1
				maxH1 = height
			} else if height > maxH2 {
				maxH2 = height
			}
		}
		for _, v := range children {
			targetH := maxH1
			if h[v] == maxH1 {
				targetH = maxH2
			}
			gain := int64(targetH+1) * size[v]
			current := dist[u] + gain
			if current > f[u] {
				f[u] = current
			}
		}
	}
	res[u-1] = f[u]
}

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	a = readInts(n)
	adj = make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		edge := readInts(2)
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	size = make([]int64, n+1)
	dist = make([]int64, n+1)
	h = make([]int, n+1)
	f = make([]int64, n+1)
	res = make([]int64, n)
	dfs(1, 0)
	printSlice(res)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	line := readLine()
	if line == "" {
		return
	}
	t, _ := strconv.Atoi(line)
	for range t {
		solve()
	}
}

/*
Cost Definition: For a subtree rooted at r, the base cost is ∑au​⋅d(r,u).

Recursive Relation: The cost of a subtree r is the sum of (cost of child c + sum of values in child c's subtree).

The Operation: To maximize the cost by moving a subtree u, we should attach it to the current deepest leaf in the remaining tree. The gain is sum(u) * (max_depth_available - depth(u) + 1).

DFS Strategy: We perform a single DFS to calculate:

    size[u]: Sum of ai​ in subtree u.

    dist[u]: Base cost of subtree u.

    h[u]: Maximum height (distance to deepest leaf) in subtree u.

    f[u]: Max cost of subtree u after performing at most one operation within it.

Optimization: For each node u, we track the two largest heights among its children to calculate the potential gain when moving one child's subtree to the deepest leaf of another child's branch.
*/
