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
		var n, k int
		fmt.Fscan(r, &n, &k)
		var s string
		fmt.Fscan(r, &s)
		o := 0
		for _, c := range s {
			if c == '1' {
				o++
			}
		}
		if o <= k || 2*k > n {
			fmt.Fprintln(w, "Alice")
		} else {
			fmt.Fprintln(w, "Bob")
		}
	}
}
