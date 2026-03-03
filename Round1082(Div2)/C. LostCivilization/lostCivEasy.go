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
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	arr := readInts(n)
	// res, x, y := 1, arr[0], arr[0]
	// for i := 1; i < n; i++ {
	// 	if res == 0 {
	// 		res++
	// 		x = arr[i]
	// 		y = arr[i]
	// 	} else {
	// 		if arr[i] >= x+1 && arr[i] <= y+1 {
	// 			if arr[i] > y {
	// 				y = arr[i]
	// 			}
	// 		} else {
	// 			res++
	// 			x = arr[i]
	// 			y = arr[i]
	// 		}
	// 	}
	// }
	st := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] == arr[i]+1 {
			st = st[:len(st)-1]
		}
		st = append(st, arr[i])
	}
	fmt.Fprintln(out, len(st))
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
