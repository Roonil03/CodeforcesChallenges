package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

// var pow2 [21]int64

func h1(x, g int) int64 {
	n := x - 1
	if g > n {
		return 0
	}
	m := (n - g) / 2
	return int64(1<<bits.OnesCount(uint(g))) * h2(m, g)
}

func h2(m, g int) int64 {
	if m < 0 {
		return 0
	}
	var eq, l int64 = 1, 0
	for bit := 30; bit >= 0; bit-- {
		mb := (m >> bit) & 1
		gb := (g >> bit) & 1
		var neq, nl int64
		if eq > 0 {
			if gb == 1 {
				if mb == 0 {
					neq += eq
				} else {
					nl += eq
				}
			} else {
				if mb == 0 {
					neq += eq
				} else {
					nl += eq
					neq += eq
				}
			}
		}
		if l > 0 {
			if gb == 1 {
				nl += l
			} else {
				nl += l * 2
			}
		}
		eq, l = neq, nl
	}
	return eq + l
}

func solve() {
	temp := readInts(2)
	x1, x2 := temp[0], temp[1]
	if x1 > x2 {
		g := x2
		fmt.Fprintf(out, "%d %d\n", g+1, x1)
		return
	}
	a := 0
	b := h1(x2, 0)
	for i := 1; i < x1; i++ {
		cur := h1(x2, i)
		if cur < b {
			b = cur
			a = i
		}
	}
	fmt.Fprintf(out, "%d %d\n", a+1, x1)
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
