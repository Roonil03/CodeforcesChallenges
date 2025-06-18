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
		var s1, s2 string
		fmt.Fscan(r, &s1, &s2)
		m := len(s1)
		inf := 1 << 30
		dp := make([][2][2]int, m+1)
		for i := 0; i <= m; i++ {
			for a := 0; a < 2; a++ {
				for b := 0; b < 2; b++ {
					dp[i][a][b] = inf
				}
			}
		}
		dp[m][0][0], dp[m][0][1], dp[m][1][0], dp[m][1][1] = 0, 0, 0, 0
		for i := m - 1; i >= 0; i-- {
			li := int(s1[i] - '0')
			ri := int(s2[i] - '0')
			for a := 0; a < 2; a++ {
				for b := 0; b < 2; b++ {
					low, high := 0, 9
					if a == 1 {
						low = li
					}
					if b == 1 {
						high = ri
					}
					for d := low; d <= high; d++ {
						if a == 1 && d < li {
							continue
						}
						if b == 1 && d > ri {
							continue
						}
						na := 0
						if a == 1 && d == li {
							na = 1
						}
						nb := 0
						if b == 1 && d == ri {
							nb = 1
						}
						c := dp[i+1][na][nb]
						if d == li {
							c++
						}
						if d == ri {
							c++
						}
						if c < dp[i][a][b] {
							dp[i][a][b] = c
						}
					}
				}
			}
		}
		fmt.Fprintln(w, dp[0][1][1])
	}
}
