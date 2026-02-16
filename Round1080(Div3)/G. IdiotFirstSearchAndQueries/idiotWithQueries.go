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

func readInt() int {
	s := readLine()
	if s == "" {
		return 0
	}
	v, _ := strconv.Atoi(s)
	return v
}

func readTwoInts() (int, int) {
	s := readLine()
	parts := strings.Fields(s)
	if len(parts) < 2 {
		return 0, 0
	}
	v1, _ := strconv.Atoi(parts[0])
	v2, _ := strconv.Atoi(parts[1])
	return v1, v2
}

const LOG = 20

func solve() {
	line := readLine()
	if line == "" {
		return
	}
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])
	adj := make([][2]int, n+1)
	parent := make([]int, n+1)
	for i := 0; i <= n; i++ {
		l, r := readTwoInts()
		adj[i][0] = l
		adj[i][1] = r
		if l != 0 {
			parent[l] = i
		}
		if r != 0 {
			parent[r] = i
		}
	}
	dp2 := make([]int64, n+1)
	size := make([]int64, n+1)
	var dfsSize func(u int)
	dfsSize = func(u int) {
		if u == 0 {
			return
		}
		if adj[u][0] == 0 {
			size[u] = 1
			dp2[u] = 0
			return
		}
		dfsSize(adj[u][0])
		dfsSize(adj[u][1])
		size[u] = size[adj[u][0]] + size[adj[u][1]] + 1
		dp2[u] = 2*size[u] - 2
	}
	dfsSize(1)
	totalTime := make([]int64, n+1)
	up := make([][LOG]int, n+1)
	var dfsPath func(u int, p int, dist int64)
	dfsPath = func(u int, p int, dist int64) {
		totalTime[u] = dist
		up[u][0] = p
		for i := 1; i < LOG; i++ {
			up[u][i] = up[up[u][i-1]][i-1]
		}
		if adj[u][0] == 0 {
			return
		}
		dfsPath(adj[u][0], u, dist+1)
		dfsPath(adj[u][1], u, dist+2*size[adj[u][0]]+1)
	}
	dfsPath(1, 0, 0)
	tour := make([]int, 0, 2*n)
	firstPos := make([]int, n+1)
	var dfsTour func(u int)
	dfsTour = func(u int) {
		firstPos[u] = len(tour)
		tour = append(tour, u)
		if adj[u][0] == 0 {
			return
		}
		dfsTour(adj[u][0])
		tour = append(tour, u)
		dfsTour(adj[u][1])
		tour = append(tour, u)
	}
	dfsTour(1)
	res := make([]int, q)
	for i := 0; i < q; i++ {
		v, k := readTwoInts()
		kv := int64(k)
		curr := v
		for j := LOG - 1; j >= 0; j-- {
			anc := up[curr][j]
			if anc != 0 && totalTime[v]-totalTime[anc] <= kv {
				curr = anc
			}
		}
		rem := kv - (totalTime[v] - totalTime[curr])
		res[i] = tour[firstPos[curr]+int(rem)]
	}
	for i, v := range res {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	tStr := readLine()
	if tStr == "" {
		return
	}
	t, _ := strconv.Atoi(tStr)
	for t > 0 {
		t--
		solve()
	}
}
