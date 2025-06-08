package main

import "fmt"

func solve(n int, a []int) string {
	d := a[1] - a[0]
	for i := 2; i < n; i++ {
		if a[i]-a[i-1] != d {
			return "NO"
		}
	}
	c := a[0]
	if (c-d)%(n+1) == 0 && c-d >= 0 && d*n+c >= 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		t--
		var n int
		fmt.Scan(&n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		fmt.Println(solve(n, a))
	}
}
