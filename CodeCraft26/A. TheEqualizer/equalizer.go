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
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%2 != 0 || (n*k)%2 == 0 {
		fmt.Fprintln(out, "yes")
	} else {
		fmt.Fprintln(out, "no")
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
