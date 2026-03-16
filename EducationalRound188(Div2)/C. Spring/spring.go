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
	temp := readInt64s(4)
	a, b, c, m := temp[0], temp[1], temp[2], temp[3]
	a1, b1, c1 := m/a, m/b, m/c
	lab, lbc, lac := lcm(a, b, m), lcm(b, c, m), lcm(a, c, m)
	labc := lcm(lab, c, m)
	ab, bc, ac, abc := m/lab, m/lbc, m/lac, m/labc
	r1 := 6*a1 - 3*ab - 3*ac + 2*abc
	r2 := 6*b1 - 3*ab - 3*bc + 2*abc
	r3 := 6*c1 - 3*ac - 3*bc + 2*abc
	fmt.Fprintf(out, "%d %d %d\n", r1, r2, r3)
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a %= b
		a, b = b, a
	}
	return a
}

func lcm(a, b, m int64) int64 {
	g := gcd(a, b)
	t := a / g
	if b > 0 && t > m/b {
		return m + 1
	}
	return t * b
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
