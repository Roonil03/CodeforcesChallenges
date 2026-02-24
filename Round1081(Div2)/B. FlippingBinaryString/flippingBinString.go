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

func h1(arr []int) {
	fmt.Fprintln(out, len(arr))
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	if len(arr) > 0 {
		fmt.Fprintln(out)
	}
}

func solve() {
	// line := readLine()
	// parts := strings.Fields(line)
	// n, _ := strconv.Atoi(parts[0])
	_ = readLine()
	s := readLine()
	o := []int{}
	z := []int{}
	for i, c := range s {
		if c == '1' {
			o = append(o, i+1)
		} else {
			z = append(z, i+1)
		}
	}
	if len(o)%2 == 0 {
		h1(o)
		return
	}
	if len(z)%2 == 1 {
		h1(z)
		return
	}
	fmt.Fprintln(out, -1)
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
