package main

import (
	"fmt"
)

var comp []bool
var primes []int

func sieve() {
	comp = make([]bool, 1e5+1)
	for i := 2; i*i <= 1e5; i++ {
		if !comp[i] {
			for j := i * i; j <= 1e5; j += i {
				comp[j] = true
			}
		}
	}
	for i := 2; i <= 1e5; i++ {
		if !comp[i] {
			primes = append(primes, i)
		}
	}
}

func solve() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n+1)

	for i := len(primes) - 1; i >= 0; i-- {
		prime := primes[i]
		var cycle []int
		for j := prime; j <= n; j += prime {
			if p[j] == 0 {
				cycle = append(cycle, j)
			}
		}
		for k := 0; k < len(cycle); k++ {
			p[cycle[k]] = cycle[(k+1)%len(cycle)]
		}
	}

	for i := 1; i <= n; i++ {
		if p[i] == 0 {
			p[i] = i
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Print(p[i])
		if i != n {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}

func main() {
	sieve()
	var t int
	fmt.Scan(&t)
	for t > 0 {
		solve()
		t--
	}
}
