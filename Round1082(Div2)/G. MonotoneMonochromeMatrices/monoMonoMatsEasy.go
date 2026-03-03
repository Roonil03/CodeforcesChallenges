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
	q, _ := strconv.Atoi(parts[1])
	R := make([]int64, n+1)
	C := make([]int64, n+1)
	rS := make([]int64, q+2)
	var sumC int64
	var sumRs int64
	for range q {
		line = readLine()
		parts = strings.Fields(line)
		r, _ := strconv.Atoi(parts[0])
		c, _ := strconv.Atoi(parts[1])
		val := R[r]
		R[r]++
		sumRs += 2*rS[val+1] + 1
		rS[val+1]++
		oldC := C[c]
		C[c]++
		sumC += 2*oldC + 1
		if sumRs == sumC {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
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

/*
R[r] be the number of black cells in row r.
C[c] be the number of black cells in column c.

sumC  = sum of squares of column counts
sumRs = sum over all pairs of rows that have equal counts

The matrix is monotone if and only if:
    sumRs == sumC

For each query (r, c):
1. Let val = R[r].
   Increase R[r] by 1.
   Before increment, there were rS[val+1] rows having val+1 black cells.
   After increment, the new row joins that group.
   So update:
       sumRs += 2 * rS[val+1] + 1
       rS[val+1]++

2. Let oldC = C[c].
   Increase C[c] by 1.
   Update:
       sumC += 2 * oldC + 1

3. If sumRs == sumC:
       Matrix is monotone → print YES
   Else:
       print NO
*/
