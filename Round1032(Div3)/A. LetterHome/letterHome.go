package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		var n, s int
		fmt.Scan(&n, &s)
		x := make([]int, n)
		for i := range x {
			fmt.Scan(&x[i])
		}
		l, r := x[0], x[n-1]
		var a int
		if s <= l {
			a = r - s
		} else if s >= r {
			a = s - l
		} else {
			d := s - l
			if r-s < d {
				d = r - s
			}
			a = r - l + d
		}
		fmt.Println(a)
	}
}
