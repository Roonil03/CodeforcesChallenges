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
	temp := readInt64s(3)
	n, h, k := int(temp[0]), temp[1], temp[2]
	arr := readInt64s(n)
	pre := make([]int64, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + arr[i]
	}
	sum := pre[n]
	a := make([]int64, n+1)
	a[1] = arr[0]
	for i := 2; i <= n; i++ {
		if arr[i-1] < a[i-1] {
			a[i] = arr[i-1]
		} else {
			a[i] = a[i-1]
		}
	}
	suf := make([]int64, n+2)
	suf[n+1] = 0
	for i := n; i >= 1; i-- {
		if arr[i-1] > suf[i+1] {
			suf[i] = arr[i-1]
		} else {
			suf[i] = suf[i+1]
		}
	}
	res := int64(1<<62) - 1
	for i := 1; i <= n; i++ {
		e := pre[i]
		if i < n {
			if imp := suf[i+1] - a[i]; imp > 0 {
				e += imp
			}
		}
		q := int64(0)
		if e < h {
			// if sum == 0{
			// 	continue
			// }
			q = (h - e + sum - 1) / sum
		}
		t := q*(int64(n)+k) + int64(i)
		if t < res {
			res = t
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
