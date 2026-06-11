package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func factor(x int64) []int64 {
	var f []int64
	for d := int64(1); d*d <= x; d++ {
		if x%d == 0 {
			f = append(f, d)
			if x/d != d {
				f = append(f, x/d)
			}
		}
	}
	slices.Sort(f)
	return f
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var N, A, B int64
	fmt.Fscan(in, &N, &A, &B)

	G := gcd(A, B)
	A /= G
	B /= G

	fA := factor(A)
	fB := factor(B)
	lA, lB := len(fA), len(fB)

	eA := make([][]int, lA)
	eB := make([][]int, lB)
	pA := make([][]int, lA)
	pB := make([][]int, lB)

	prime := make(map[int64]bool)

	for i := 0; i < lA; i++ {
		nd := 0
		for j := 0; j < i; j++ {
			if fA[i]%fA[j] == 0 {
				nd++
			}
		}
		if nd <= 1 {
			prime[fA[i]] = true
		}
	}
	for i := 0; i < lB; i++ {
		nd := 0
		for j := 0; j < i; j++ {
			if fB[i]%fB[j] == 0 {
				nd++
			}
		}
		if nd <= 1 {
			prime[fB[i]] = true
		}
	}

	for i := 0; i < lA; i++ {
		for j := i; j < lA; j++ {
			if fA[j]%fA[i] == 0 {
				eA[i] = append(eA[i], j)
				if prime[fA[j]/fA[i]] {
					pA[i] = append(pA[i], j)
				}
			}
		}
	}
	for i := 0; i < lB; i++ {
		for j := i; j < lB; j++ {
			if fB[j]%fB[i] == 0 {
				eB[i] = append(eB[i], j)
				if prime[fB[j]/fB[i]] {
					pB[i] = append(pB[i], j)
				}
			}
		}
	}

	dp := make([][]int64, lA)
	for i := range dp {
		dp[i] = make([]int64, lB)
		for j := range dp[i] {
			dp[i][j] = N + 1
		}
	}
	dp[0][0] = 0

	for i := 0; i < lA; i++ {
		for j := 0; j < lB; j++ {
			if dp[i][j] == N+1 {
				continue
			}
			for _, ei := range eA[i] {
				for _, ej := range pB[j] {
					c := fA[ei] / fA[i]
					d := fB[ej] / fB[j]
					cost := dp[i][j] + max(c, d)
					if cost < dp[ei][ej] {
						dp[ei][ej] = cost
					}
				}
			}
			for _, ei := range pA[i] {
				for _, ej := range eB[j] {
					c := fA[ei] / fA[i]
					d := fB[ej] / fB[j]
					cost := dp[i][j] + max(c, d)
					if cost < dp[ei][ej] {
						dp[ei][ej] = cost
					}
				}
			}
		}
	}
	fmt.Println(dp[lA-1][lB-1])
}

/*
Algorithm and Logic Explanation:

1. Base Reduction:
The algorithm starts by simplifying the targets A and B. By dividing both by their Greatest Common Divisor
(GCD), the starting state effectively shifts so we are tracking the multiplicative paths to reach the
non-shared factors of A and B.

2. Factorization:
It computes all the divisors of the newly simplified A and B, maintaining them in sorted slices
(`fA` and `fB`).

3. Identifying "Prime" Multipliers:
By iterating through the list of divisors, it identifies the subset of divisors that act as prime numbers
(or 1) in this context. A divisor is flagged as "prime" if it has at most 1 proper divisor (meaning it's
only divisible by 1).

4. Precomputing Valid Transitions:
It precomputes adjacency lists to represent valid jumps (multiplications) from one divisor to another.
- `eA` and `eB` store generic edges: any multiplier to reach a larger divisor.
- `pA` and `pB` store prime edges: specific jumps where the multiplier used is strictly one of our identified "primes".

5. Dynamic Programming (DP):
A 2D DP array (`dp[i][j]`) is initialized with a high placeholder value, tracking the minimum cost to reach
divisor `fA[i]` and `fB[j]`. The base state `dp[0][0]` (representing reaching factor 1 for both) costs 0.

6. State Transitions:
We iterate through all possible states of `(i, j)`. For a state to move to a new state `(ei, ej)`, the problem
imposes a rule: the jump on at least one dimension must be an exact prime multiplier, while the jump on the
other can be any valid multiple.
- The inner loops simulate these two cases:
  a) `A` multiplies by any valid factor, `B` multiplies by a prime factor.
  b) `A` multiplies by a prime factor, `B` multiplies by any valid factor.

7. Cost Calculation:
The cost added to the current DP state is the maximum of the two applied multipliers (`max(c, d)`). We greedily
check if this path offers a cheaper way to reach the state `dp[ei][ej]` and update it. Finally, `dp[lA-1][lB-1]`
outputs the minimum cost to achieve the full target state.
*/
