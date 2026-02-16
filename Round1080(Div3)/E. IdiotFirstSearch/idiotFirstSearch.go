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

const MOD = int64(1e9) + 7

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	// arr := readInts(n)
	l, r := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		lr := readInts(2)
		l[i] = lr[0]
		r[i] = lr[1]
	}
	ord, q := make([]int, 0, n), make([]int, 0, n)
	q = append(q, 1)
	top := 0
	for top < len(q) {
		u := q[top]
		top++
		ord = append(ord, u)
		if l[u] != 0 {
			q = append(q, l[u], r[u])
		}
	}
	e := make([]int64, n+1)
	for i := len(ord) - 1; i >= 0; i-- {
		u := ord[i]
		if l[u] == 0 {
			e[u] = int64(1)
		} else {
			e[u] = (e[l[u]] + e[r[u]] + 3) % MOD
		}
	}
	res := make([]int64, n+1)
	res[1] = e[1]
	for i := range ord {
		u := ord[i]
		if l[u] != 0 {
			a, b := l[u], r[u]
			// val := (res[u] + e[u]) % MOD
			res[a], res[b] = (res[u]+e[a])%MOD, (res[u]+e[b])%MOD
		}
	}
	for i := 1; i <= n; i++ {
		if i > 1 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, res[i])
	}
	fmt.Fprintln(out)
	// printSlice(arr)
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
