package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var t int
	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(r, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		c := make([]int, n+2)
		for _, v := range a {
			c[v]++
		}
		pre := make([]bool, n+2)
		pre[0] = true
		for i := 1; i <= n; i++ {
			pre[i] = pre[i-1] && c[i-1] > 0
		}
		x := make([]int, n+3)
		for m := 0; m <= n; m++ {
			if !pre[m] {
				continue
			}
			l := c[m]
			h := n - m
			if l <= h {
				x[l]++
				x[h+1]--
			}
		}
		s := 0
		for k := 0; k <= n; k++ {
			s += x[k]
			if k > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprint(w, s)
		}
		fmt.Fprint(w, "\n")
	}
}
