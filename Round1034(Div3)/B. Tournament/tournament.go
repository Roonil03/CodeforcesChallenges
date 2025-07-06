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
	for i := 0; i < t; i++ {
		var n, j, k int
		fmt.Fscan(r, &n, &j, &k)
		a := make([]int, n)
		for x := 0; x < n; x++ {
			fmt.Fscan(r, &a[x])
		}
		s := a[j-1]
		if k > 1 {
			fmt.Fprint(w, "YES\n")
		} else {
			c := 0
			for _, v := range a {
				if v > s {
					c++
				}
			}
			if c == 0 {
				fmt.Fprint(w, "YES\n")
			} else {
				fmt.Fprint(w, "NO\n")
			}
		}
	}
}
