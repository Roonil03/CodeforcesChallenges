package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q(n int) []int {
	p := make([]int, 0, n)
	if n >= 2 {
		p = append(p, 2)
	}
	if n >= 3 {
		p = append(p, 3)
	}
	if n >= 4 {
		p = append(p, n)
		for i := 4; i <= n-1; i++ {
			p = append(p, i)
		}
	}
	if n >= 2 {
		p = append(p, 1)
	}
	return p
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())
	for ; t > 0; t-- {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())
		p := q(n)
		for i, v := range p {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}
