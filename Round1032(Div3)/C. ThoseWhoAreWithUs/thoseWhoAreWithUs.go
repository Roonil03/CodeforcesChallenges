// Couldn't be solved during contest time, has been solved post the contest hours

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20) // 1 MiB buffer
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n, m int
		fmt.Fscan(in, &n, &m)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
		}
		mx, cntMx := 0, 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Fscan(in, &a[i][j])
				v := a[i][j]
				if v > mx {
					mx, cntMx = v, 1
				} else if v == mx {
					cntMx++
				}
			}
		}
		row := make([]int, n)
		col := make([]int, m)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if a[i][j] == mx {
					row[i]++
					col[j]++
				}
			}
		}
		flag := 0
	outer:
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if row[i]+col[j]-boolToInt(a[i][j] == mx) == cntMx {
					flag = 1
					break outer
				}
			}
		}
		fmt.Fprintln(out, mx-flag)
	}
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

/*
The program scans an n x m grid once to record the global maximum value mx and how many times
it appears (cntMx), then walks the matrix a second time to build two auxiliary integer arrays:
row[i] holds the number of maxima in row i, and col[j] holds the number in column j. Using
these counts, it checks whether there exists at least one cell whose row-plus-column total
(subtracting one if that very cell is also a maximum) equals cntMx; the presence of such a
"covering" cell sets a flag that causes the final answer to be mx 1, otherwise it prints mx.
The only data structures are the input matrix ([] []int), two one-dimensional integer slices
for row and column counts, and a few scalars. Two complete passes over the grid dominate the
work, so the time complexity is O(nm); the extra memory beyond the input is O(n + m), giving
overall space O(nm) if the matrix is stored, or O(n + m) if the values are processed on-the-fly.
*/
