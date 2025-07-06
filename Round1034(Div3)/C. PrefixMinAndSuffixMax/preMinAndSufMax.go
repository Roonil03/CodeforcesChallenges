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
	for t > 0 {
		t--
		var n int
		fmt.Fscan(r, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		mn, mx := int(1e9+7), 0
		pre := make([]bool, n)
		for i, v := range a {
			if v < mn {
				mn = v
				pre[i] = true
			}
		}
		su := make([]bool, n)
		for i := n - 1; i >= 0; i-- {
			if a[i] > mx {
				mx = a[i]
				su[i] = true
			}
		}
		b := make([]byte, n)
		for i := 0; i < n; i++ {
			if pre[i] || su[i] {
				b[i] = '1'
			} else {
				b[i] = '0'
			}
		}
		fmt.Fprint(w, string(b), "\n")
	}
}
