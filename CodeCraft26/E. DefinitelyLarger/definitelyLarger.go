package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	p := readInts(n)
	d := readInts(n)
	type E struct {
		id int
		p  int
		d  int
	}
	e := make([]E, n)
	for i := range n {
		e[i] = E{id: i + 1, p: p[i], d: d[i]}
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].p > e[j].p
	})
	l := make([]int, 0, n)
	for _, v := range e {
		i := v.id
		req := v.d
		if req < 0 {
			fmt.Fprintln(out, "-1")
			return
		}
		if req == 0 {
			l = append(l, i)
			continue
		}
		c := 0
		fg := false
		for k := len(l) - 1; k >= 0; k-- {
			if l[k] > i {
				c++
			}
			if c == req {
				l = append(l, 0)
				copy(l[k+1:], l[k:])
				l[k] = i
				fg = true
				break
			}
		}
		if !fg {
			fmt.Fprintln(out, "-1")
			return
		}
	}
	q := make([]int, n)
	for rank, id := range l {
		q[id-1] = rank + 1
	}
	printSlice(q)
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
