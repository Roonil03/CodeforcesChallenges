package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var T int
	fmt.Fscan(r, &T)
	for T > 0 {
		T--
		var n, m, q int
		fmt.Fscan(r, &n, &m, &q)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		gh := make([]int, 0, 16)
		for d := 1; d*d <= m; d++ {
			if m%d == 0 {
				gh = append(gh, d)
				if d*d != m {
					gh = append(gh, m/d)
				}
			}
		}
		we := len(gh)
		v := make(map[int]int, we)
		zx := make([]int, we)
		for i, g := range gh {
			v[g] = i
			zx[i] = m/g - 1
		}
		dec := make([]int, we)
		for i := 1; i < n; i++ {
			x, y := a[i-1], a[i]
			for j, g := range gh {
				if y%g < x%g {
					dec[j]++
				}
			}
		}
		for i := 0; i < q; i++ {
			var t int
			fmt.Fscan(r, &t)
			if t == 1 {
				var p, x int
				fmt.Fscan(r, &p, &x)
				p--
				old := a[p]
				a[p] = x
				for j, g := range gh {
					if p > 0 {
						_, r := old%g, a[p-1]%g
						od := 0
						if old%g < r {
							od = 1
						}
						nd := 0
						if a[p]%g < r {
							nd = 1
						}
						dec[j] += nd - od
					}
					if p+1 < n {
						r1, r2 := a[p+1]%g, old%g
						od := 0
						if r1 < r2 {
							od = 1
						}
						nd := 0
						if r1 < a[p]%g {
							nd = 1
						}
						dec[j] += nd - od
					}
				}
			} else {
				var k int
				fmt.Fscan(r, &k)
				g := gcd(k, m)
				j := v[g]
				if dec[j] <= zx[j] {
					fmt.Fprintln(w, "YES")
				} else {
					fmt.Fprintln(w, "NO")
				}
			}
		}
	}
}
