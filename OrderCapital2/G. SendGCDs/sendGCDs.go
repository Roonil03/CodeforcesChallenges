package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	const N = 1000001
	dv := make([]int, N)
	for i := range dv {
		dv[i] = i
	}
	for i := 2; i < N; i++ {
		if dv[i] == i {
			for j := i + i; j < N; j += i {
				if dv[j] == j {
					dv[j] = i
				}
			}
		}
	}
	const LIM = 150
	var ps []int
	var pps []int
	for x := 2; ; x++ {
		if dv[x] == x {
			ps = append(ps, x)
			y := x
			for int64(y)*int64(x) < int64(N) {
				y *= x
			}
			pps = append(pps, y)
			if len(ps) == LIM {
				break
			}
		}
	}
	var goods []int
	for i := 1; i < N; i++ {
		good := true
		x := i
		for x > 1 {
			if dv[x] > ps[len(ps)-1] {
				good = false
				break
			}
			x /= dv[x]
		}
		if good {
			goods = append(goods, i)
		}
	}
	var foo string
	if _, err := fmt.Fscan(in, &foo); err != nil {
		return
	}
	var tt int
	fmt.Fscan(in, &tt)
	for tt > 0 {
		tt--
		if foo == "first" {
			var n int
			fmt.Fscan(in, &n)
			a := make([]int, n)
			for i := 0; i < n; i++ {
				fmt.Fscan(in, &a[i])
			}
			var seq []byte
			for i := 0; i < n; i++ {
				for j := 0; j < 20; j++ {
					b := (a[i] >> j) & 1
					seq = append(seq, byte('0'+b))
				}
			}
			for len(seq)%18 != 0 {
				seq = append(seq, '0')
			}
			var send []int
			for i := 0; i < len(seq); i += 18 {
				id := 0
				for j := 0; j < 18; j++ {
					if seq[i+j] == '1' {
						id |= 1 << j
					}
				}
				num := goods[id]
				send = append(send, num)
			}
			for _, y := range pps {
				send = append(send, y)
			}
			fmt.Fprintln(out, len(send))
			for i := 0; i < len(send); i++ {
				if i == len(send)-1 {
					fmt.Fprint(out, send[i], "\n")
				} else {
					fmt.Fprint(out, send[i], " ")
				}
			}
			out.Flush()
		} else {
			var n, k int
			fmt.Fscan(in, &n, &k)
			var seq []byte
			for i := 0; i < k-LIM; i++ {
				num := 1
				for j := k - LIM; j < k; j++ {
					out.Flush()
					fmt.Printf("? %d %d\n", i+1, j+1)
					var got int
					fmt.Scan(&got)
					num *= got
				}
				idx := sort.SearchInts(goods, num)
				for j := 0; j < 18; j++ {
					b := (idx >> j) & 1
					seq = append(seq, byte('0'+b))
				}
			}
			var a []int
			for i := 0; i+19 < len(seq); i += 20 {
				num := 0
				for j := 0; j < 20; j++ {
					if seq[i+j] == '1' {
						num |= 1 << j
					}
				}
				a = append(a, num)
			}
			fmt.Print("!")
			for i := 0; i < n; i++ {
				fmt.Printf(" %d", a[i])
			}
			fmt.Println()
		}
	}
}

/*
This algorithm implements an interactive communication and encoding scheme to accurately reconstruct
state values over two distinct runs. During an extensive precomputation step, it initializes a prime
sieve up to 1,000,000, chooses the first 150 prime numbers to derive high-power modular bases ("anchors"),
and tracks a subset of integers ("goods") completely smooth over these chosen primes. In the "first"
run phase, it behaves as a sender that compresses an array of numbers into a single continuous 20-bit
binary representation, chunks this output into 18-bit block codes, maps them directly into index positions
of the "goods" collection, and safely appends the prime-power base elements. In the "second" run phase,
it acts as an interactive decoder that discovers the value of each chunk by repeatedly querying the greatest
common divisor (GCD) against the pre-appended base anchors. Multiplying these shared factors seamlessly
recovers the encoded value, which is decoded back to the index via a binary search step and unrolled to
fetch the initial integer values. The overall Time Complexity is O(N log log N) for the initial sieve
operations, combined with O(N_a) operations for the compression code path, and exactly 150 * (K - 150)
queries for execution on the decoder side. Its Space Complexity evaluates to O(N) memory due to the
large arrays allocated for tracking prime values, exponents, and valid coordinate buckets up to the
designated absolute bound.
*/
