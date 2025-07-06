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
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(r, &n)
		if n%4 == 0 {
			fmt.Fprintln(w, "Bob")
		} else {
			fmt.Fprintln(w, "Alice")
		}
	}
}
