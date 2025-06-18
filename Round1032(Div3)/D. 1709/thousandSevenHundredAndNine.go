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
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &b[i])
		}
		var o [][]int
		sa := append([]int(nil), a...)
		sb := append([]int(nil), b...)
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if sa[j] > sa[j+1] {
					sa[j], sa[j+1] = sa[j+1], sa[j]
					o = append(o, []int{1, j + 1})
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if sb[j] > sb[j+1] {
					sb[j], sb[j+1] = sb[j+1], sb[j]
					o = append(o, []int{2, j + 1})
				}
			}
		}
		for i := 0; i < n; i++ {
			if sa[i] > sb[i] {
				sa[i], sb[i] = sb[i], sa[i]
				o = append(o, []int{3, i + 1})
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if sa[j] > sa[j+1] {
					sa[j], sa[j+1] = sa[j+1], sa[j]
					o = append(o, []int{1, j + 1})
				}
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n-1; j++ {
				if sb[j] > sb[j+1] {
					sb[j], sb[j+1] = sb[j+1], sb[j]
					o = append(o, []int{2, j + 1})
				}
			}
		}
		fmt.Fprintln(w, len(o))
		for _, v := range o {
			fmt.Fprintln(w, v[0], v[1])
		}
	}
}
