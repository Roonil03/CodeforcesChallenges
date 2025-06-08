package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, x int
		fmt.Scan(&n, &x)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
		}
		p, u, r := 0, false, 0
		for p < n {
			if a[p] == 0 || r > 0 || !u {
				if a[p] == 1 && r == 0 && !u {
					u, r = true, x
				}
				p++
				if r > 0 {
					r--
				}
			} else {
				break
			}
		}
		if p >= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
