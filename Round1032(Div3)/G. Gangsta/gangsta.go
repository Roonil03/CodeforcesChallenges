package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(r, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(r, &n)
		var s string
		fmt.Fscan(r, &s)
		m := 2*n + 5
		c := make([]int, m)
		u := make([]int64, m)
		add := func(f []int, i int, v int) {
			for ; i < m; i += i & -i {
				f[i] += v
			}
		}
		adds := func(f []int64, i int, v int) {
			for ; i < m; i += i & -i {
				f[i] += int64(v)
			}
		}
		qry := func(f []int, i int) int {
			rt := 0
			for ; i > 0; i -= i & -i {
				rt += f[i]
			}
			return rt
		}
		qrys := func(f []int64, i int) int64 {
			sm := int64(0)
			for ; i > 0; i -= i & -i {
				sm += f[i]
			}
			return sm
		}
		b := n + 2
		add(c, b, 1)
		adds(u, b, 0)
		d, a := 0, int64(0)
		for _, ch := range s {
			if ch == '0' {
				d++
			} else {
				d--
			}
			ix := d + b
			lc := qry(c, ix)
			ls := qrys(u, ix)
			tc := qry(c, m-1)
			ts := qrys(u, m-1)
			gc := tc - lc
			gs := ts - ls
			a += int64(d)*int64(lc) - ls + gs - int64(d)*int64(gc)
			add(c, ix, 1)
			adds(u, ix, d)
		}
		tl := int64(n) * (int64(n) + 1) * (int64(n) + 2) / 6
		fmt.Println((tl + a) / 2)
	}
}
