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
	if err != nil && len(line) == 0 {
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
	n, k := temp[0], temp[1]
	arr := readInts(n)
	p := readInts(k)
	x := arr[p[0]-1]
	f := make([]int, k+1)
	fg := false
	for i := 0; i < p[0]-1; i++ {
		if (arr[i] ^ x) == 1 {
			if !fg {
				f[0]++
				fg = true
			}
		} else {
			fg = false
		}
	}
	for i := 1; i < k; i++ {
		fg = false
		for j := p[i-1]; j < p[i]-1; j++ {
			if (arr[j] ^ x) == 1 {
				if !fg {
					f[i]++
					fg = true
				}
			} else {
				fg = false
			}
		}
	}
	fg = false
	for i := p[k-1]; i < n; i++ {
		if (arr[i] ^ x) == 1 {
			if !fg {
				f[k]++
				fg = true
			}
		} else {
			fg = false
		}
	}
	tot, m := 0, 0
	for _, v := range f {
		tot += v
		m = max(v, m)
	}
	res := max(tot, 2*m)
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
