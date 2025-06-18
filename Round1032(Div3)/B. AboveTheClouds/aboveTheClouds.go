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
		var s string
		fmt.Fscan(r, &n, &s)
		if n < 3 {
			fmt.Println("No")
			continue
		}
		r := make([][26]bool, n)
		c := [26]bool{}
		for i := n - 1; i >= 0; i-- {
			c[s[i]-'a'] = true
			r[i] = c
			c = [26]bool{}
			for k := 0; k < 26; k++ {
				c[k] = r[i][k]
			}
		}
		f := false
		l := [26]bool{}
		for a := 1; a <= n-2; a++ {
			l[s[a-1]-'a'] = true
			d := s[a]
			if l[d-'a'] || (a+1 < n && r[a+1][d-'a']) {
				f = true
				break
			}
		}
		if f {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
