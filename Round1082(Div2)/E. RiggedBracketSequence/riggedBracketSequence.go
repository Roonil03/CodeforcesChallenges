package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

const MOD int64 = 998244353

var pow2 []int64

func readLine() string {
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(line)
}

func readInts(n int) []int {
	parts := strings.Fields(readLine())
	res := make([]int, n)
	for i := range n {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func readInt64s(n int) []int64 {
	parts := strings.Fields(readLine())
	res := make([]int64, n)
	for i := range n {
		res[i], _ = strconv.ParseInt(parts[i], 10, 64)
	}
	return res
}

func printSlice(arr []int) {
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	s := readLine()
	P := make([]int, n+1)
	for i := range n {
		if s[i] == '(' {
			P[i+1] = P[i] + 1
		} else {
			P[i+1] = P[i] - 1
		}
	}
	B := make([]int, n+1)
	nxt := n
	for i := n; i >= 1; i-- {
		if P[i] <= 1 {
			nxt = i
		}
		B[i] = nxt
	}
	var SumR int64
	var sumL1 int64
	sumL2 := make([]int64, n+3)
	var res int64
	for i := 1; i <= n; i++ {
		sumL1 = (sumL1 - sumL2[i] + MOD) % MOD
		dpi := (1 + SumR + sumL1) % MOD
		if s[i-1] == ')' {
			SumR = (SumR + dpi) % MOD
			res = (res + dpi) % MOD
		} else {
			res = (res + pow2[i-1]) % MOD
			if B[i] >= i+1 {
				sumL1 = (sumL1 + dpi) % MOD
				sumL2[B[i]+1] = (sumL2[B[i]+1] + dpi) % MOD
			}
		}
	}
	fmt.Fprintln(out, res)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t, _ := strconv.Atoi(readLine())
	maxN := 300000
	pow2 = make([]int64, maxN+1)
	pow2[0] = 1
	for i := 1; i <= maxN; i++ {
		pow2[i] = (pow2[i-1] * 2) % MOD
	}
	for t > 0 {
		t--
		solve()
	}
}

/*
Let '(' = +1 and ')' = -1.
Compute prefix sums P[i].

When shifting a subsequence, prefix sums change only if
the prefix partially contains the subsequence.
The dangerous case is when the last chosen character is ')',
which may decrease prefix sums by 2.

DP[i] = number of non-trivial subsequences ending at index i
        where the last character is ')'.

Precomp:
Compute prefix sums P[i].
Define B[i] as the next index ≥ i where prefix sum ≤ 1.
This helps determine valid ranges where shifting remains regular.

SumR  = sum of DP values for ')'
sumL1 = active DP contributions for '('
sumL2 = difference array to remove expired contributions

For each index i:

1. Remove expired contributions:
   sumL1 -= sumL2[i]

2. Compute:
   DP[i] = 1 + SumR + sumL1

3. If S[i] == ')':
       Add DP[i] to SumR
       Add DP[i] to answer
   Else:
       Add 2^(i-1) to answer
       If B[i] >= i+1:
           Activate DP[i] in range [i+1, B[i]]
           Use difference array technique

*/
