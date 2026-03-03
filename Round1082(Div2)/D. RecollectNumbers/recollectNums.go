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
	n, k := temp[0], temp[1]
	if k < n || k > 2*n-1 {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	m := k - n
	res := make([]int, 0, 2*n)
	if m == 0 {
		for i := 1; i <= n; i++ {
			res = append(res, i)
			res = append(res, i)
		}
	} else {
		l := m + 1
		res = append(res, 1)
		res = append(res, 2)
		for i := 2; i < l; i++ {
			res = append(res, i+1)
			res = append(res, i-1)
		}
		res = append(res, l-1)
		res = append(res, l)
		for i := l + 1; i <= n; i++ {
			res = append(res, i)
			res = append(res, i)
		}
	}
	printSlice(res)
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

/*
Each pair must be flipped at least once, so k ≥ n.
Each pair can contribute at most 2 flips, but not all pairs can reach 2.
The true maximum number of operations is 2n − 1.

Thus valid k satisfies:
    n ≤ k ≤ 2n − 1
Therefore
Let M = k − n.

Case 1: M = 0
We want exactly n operations.
Place cards as:
    1 1 2 2 3 3 ... n n
Each pair is matched immediately when its second occurrence appears.
Total operations = n.

Case 2: M > 0
We create a prefix that maximizes wasted operations.

Let L = M + 1.

Construct:
    1, 2
    For i = 2 to L−1:
        (i+1), (i−1)
    (L−1), L

This prefix creates L pairs contributing 2L−1 operations.

Then for remaining numbers i = L+1 to n:
    i, i

These contribute exactly 1 operation each.

Total operations:
    (2L−1) + (n−L) = n + (L−1) = n + M = k

If k is outside the valid range, output NO.
Otherwise output YES and print the constructed sequence.
*/
