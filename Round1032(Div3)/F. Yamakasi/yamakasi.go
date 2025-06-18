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
		var n, s, x int
		fmt.Fscan(r, &n, &s, &x)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		res := 0
		i := 0
		for i < n {
			if a[i] > x {
				i++
				continue
			}
			j := i
			for j < n && a[j] <= x {
				j++
			}
			cnt := map[int]int{0: 1}
			sum := 0
			total := 0
			for k := i; k < j; k++ {
				sum += a[k]
				total += cnt[sum-s]
				cnt[sum]++
			}
			// start := i
			noX := 0
			cnt = map[int]int{0: 1}
			sum = 0
			for k := i; k < j; k++ {
				if a[k] == x {
					// start = k + 1
					cnt = map[int]int{0: 1}
					sum = 0
				} else {
					sum += a[k]
					noX += cnt[sum-s]
					cnt[sum]++
				}
			}
			res += total - noX
			i = j
		}
		fmt.Fprintln(w, res)
	}
}
