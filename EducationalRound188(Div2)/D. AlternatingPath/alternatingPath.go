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
	for i := range n {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func readInt64s(n int) []int64 {
	parts := strings.Fields(readLine())
	res := make([]int64, n)
	for i := range n {
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
	temp := readInts(2)
	n, m := temp[0], temp[1]
	adj := make([][]int, n+1)
	for range m {
		temp = readInts(2)
		u, v := temp[0], temp[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	col := make([]int, n+1)
	res := 0
	for i := 1; i <= n; i++ {
		if col[i] == 0 {
			q := []int{i}
			col[i] = 1
			t1, t2 := 1, 0
			val := true
			head := 0
			for head < len(q) {
				cur := q[head]
				head++
				for _, v := range adj[cur] {
					if col[v] == 0 {
						col[v] = 3 - col[cur]
						if col[v] == 1 {
							t1++
						} else {
							t2++
						}
						q = append(q, v)
					} else if col[v] == col[cur] {
						val = false
					}
				}
			}
			if val {
				if t1 > t2 {
					res += t1
				} else {
					res += t2
				}
			}
		}
	}
	fmt.Fprintln(out, res)
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
