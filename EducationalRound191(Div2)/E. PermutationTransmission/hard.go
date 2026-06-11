package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var N int
		fmt.Fscan(in, &N)

		L := 0
		for 1<<L < N+1 {
			L++
		}

		S := make([]string, L)
		cnt := make([]int, L)
		for i := 0; i < L; i++ {
			fmt.Fscan(in, &S[i])
			for j := 0; j < len(S[i]); j++ {
				if S[i][j] == '1' {
					cnt[i]++
				}
			}
		}

		ord := make([]int, L)
		for i := 0; i < L; i++ {
			ord[i] = i
		}

		slices.SortFunc(ord, func(a, b int) int {
			return cnt[b] - cnt[a]
		})

		vals := make([]int, N)
		for j := 0; j < N; j++ {
			for i := 0; i < L; i++ {
				if S[ord[i]][j] == '1' {
					vals[j] |= 1 << i
				}
			}
		}

		slices.Sort(vals)

		ok := int64(1)
		for i := 0; i < N; i++ {
			if vals[i] != i+1 {
				ok = 0
				break
			}
		}

		if ok == 1 {
			freq := make(map[int]int)
			for _, c := range cnt {
				freq[c]++
			}
			for _, a := range freq {
				for j := 1; j <= a; j++ {
					ok *= int64(j)
				}
			}
		}
		fmt.Fprintln(out, ok)
	}
}

/*
Algorithm and Logic Explanation:

1. Problem Context:
The problem requires finding the number of valid ways to assign L binary strings of length N to bit
positions (0 to L-1) such that reading the j-th character of all assigned strings vertically yields
a sequence of distinct integers from 1 to N. L is the minimum number of bits needed to represent N.

2. Core Observation:
In the binary representations of the numbers from 1 to N, the lower bits flip more frequently than
the higher bits. Consequently, the least significant bit (2^0) will contain the highest number of
'1's, the next bit (2^1) will contain the second highest, and so forth. The most significant bit
(2^(L-1)) will contain the fewest '1's.

3. Algorithm Execution:
- Step 1 (Calculate Bounds): Determine L, the minimum number of bits needed for N.
- Step 2 (Count Frequencies): Read the L strings and count the total number of '1's in each string.
- Step 3 (Sort by Bit Weight): Create an array `ord` representing the indices of the strings. Sort
these indices based on the count of '1's in descending order. This greedily maps the string with
the most '1's to the 0-th bit, the second most to the 1-st bit, etc.
- Step 4 (Reconstruct Values): For each column index j (from 0 to N-1), reconstruct the integer value
using the sorted mapping and bitwise OR shifts.
- Step 5 (Validate): Sort the reconstructed N integers. If they exactly match the sequence 1, 2, ..., N,
the greedy assignment is valid. If any number is out of sequence, the mapping is impossible, returning 0.
- Step 6 (Permutations for Duplicates): If the base configuration is valid, we must account for
indistinguishable string counts. If multiple strings contain the exact same number of '1's, their
assigned bit positions can be freely swapped. We group the strings by their '1' count, calculate the
factorial for each group size, and multiply them together to get the total number of valid permutations.
*/
