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
	for i := 0; i < q; i++ {
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
R[r] = number of black cells in row r
C[c] = number of black cells in column c
rS[k] = number of rows having exactly k black cells

sumRs = sum over all row pairs with equal sizes
sumC  = sum over squares of column sizes

Matrix is monotone if and only if sumRs == sumC
For each query (r, c):
1. Let val = R[r]
   Increase R[r] by 1
   The row moves from group size val to val+1.
   If currently rS[val+1] rows already have val+1 elements,
   then the new row forms:  2 * rS[val+1] + 1
   new equal-row contributions.
   So:
        sumRs += 2 * rS[val+1] + 1
        rS[val+1]++

2. Let oldC = C[c]
   Increase C[c] by 1
   Contribution to sumC increases by:
        2 * oldC + 1
   So:
        sumC += 2 * oldC + 1

3. If sumRs == sumC:
        print YES
   else:
        print NO
*/
