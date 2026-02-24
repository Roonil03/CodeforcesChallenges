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
		panic(err)
	}
	return strings.TrimSpace(line)
}

func readInts(n int) []int {
	parts := strings.Fields(readLine())
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func readInt64s(n int) []int64 {
	parts := strings.Fields(readLine())
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i], _ = strconv.ParseInt(parts[i], 10, 64)
	}
	return res
}

func printSlice(arr []int) {
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	a := readInts(n)
	b := readInts(n)
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		cnt[a[i]]++
		cnt[b[i]]++
	}
	for i := 1; i <= n; i++ {
		if cnt[i]%2 != 0 {
			fmt.Fprintln(out, -1)
			return
		}
	}
	type Edge struct {
		to, id int
		fwd    bool
	}
	adj := make([][]Edge, n+1)
	for i := 0; i < n; i++ {
		u, v := a[i], b[i]
		if u == v {
			continue
		}
		adj[u] = append(adj[u], Edge{v, i, true})
		adj[v] = append(adj[v], Edge{u, i, false})
	}
	u := make([]bool, n)
	ptr := make([]int, n+1)
	type PE struct {
		node, edge int
		fwd        bool
	}
	var swp []int
	for i := 1; i <= n; i++ {
		for ptr[i] < len(adj[i]) && u[adj[i][ptr[i]].id] {
			ptr[i]++
		}
		if ptr[i] >= len(adj[i]) {
			continue
		}
		st := []PE{{i, -1, false}}
		for len(st) > 0 {
			v := st[len(st)-1].node
			for ptr[v] < len(adj[v]) && u[adj[v][ptr[v]].id] {
				ptr[v]++
			}
			if ptr[v] < len(adj[v]) {
				e := adj[v][ptr[v]]
				ptr[v]++
				u[e.id] = true
				st = append(st, PE{e.to, e.id, e.fwd})
			} else {
				pe := st[len(st)-1]
				st = st[:len(st)-1]
				if pe.edge >= 0 && !pe.fwd {
					swp = append(swp, pe.edge+1)
				}
			}
		}
	}
	fmt.Fprintln(out, len(swp))
	for i, v := range swp {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	if len(swp) > 0 {
		fmt.Fprintln(out)
	}
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t, _ := strconv.Atoi(readLine())
	for t > 0 {
		t--
		solve()
	}
}
