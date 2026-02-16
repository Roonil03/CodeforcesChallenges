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

func reading(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = readInts(3)
	}
	return res
}

type f struct {
	a, b, c, id int
}

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	arr := reading(n)
	a := make([]f, n)
	for i := range a {
		a[i] = f{a: arr[i][0], b: arr[i][1], c: arr[i][2], id: i}
	}
	adj, deg := make([][]int, n), make([]int, n)
	for i := range deg {
		for j := range deg {
			if i == j {
				continue
			}
			if h1(a[i], a[j]) {
				adj[i] = append(adj[i], j)
				deg[j]++
			}
		}
	}
	t, q := make([]int, 0, n), make([]int, 0, n)
	for i := range deg {
		if deg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		t = append(t, u)
		for _, v := range adj[u] {
			deg[v]--
			if deg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	dp1, dp2 := make([]int, n), make([]int, n)
	for i := range dp1 {
		u := t[i]
		if dp1[u] == 0 {
			dp1[u] = 1
		}
		for _, v := range adj[u] {
			if dp1[u]+1 > dp1[v] {
				dp1[v] = dp1[u] + 1
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		u := t[i]
		if dp2[u] == 0 {
			dp2[u] = 1
		}
		for _, v := range adj[u] {
			if dp2[v]+1 > dp2[u] {
				dp2[u] = dp2[v] + 1
			}
		}
	}
	res := make([]int, n)
	for i := range res {
		res[a[i].id] = dp1[i] + dp2[i] - 1
	}
	printSlice(res)
}

func h1(a, b f) bool {
	da := int64(a.a - b.a)
	db := int64(a.b - b.b)
	dc := int64(a.c - b.c)
	if da > 0 {
		return db*db < 4*da*dc
	}
	return da == 0 && db == 0 && dc > 0
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
