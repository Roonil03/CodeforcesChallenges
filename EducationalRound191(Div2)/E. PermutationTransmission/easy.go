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
					vals[j] ^= 1 << i
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
You are given L binary strings of length N, where L is the minimum number of bits required to represent N
(i.e., 2^L >= N+1). The goal is to determine the number of valid ways to map each of these L strings to a
unique bit position (0 to L-1) such that reading the j-th character of all assigned strings vertically
yields a distinct integer from 1 to N.

2. Core Observation:
In the binary representations of the sequence of numbers from 1 to N, the lower bits flip more frequently
than the higher bits. Consequently, the least significant bit (2^0) will contain the highest number of '1's,
the next bit (2^1) will contain the second highest, and the most significant bit (2^(L-1)) will contain the
fewest '1's.

3. Algorithm Execution:
- Step 1 (Calculate Bounds): Determine L, the number of bits needed.
- Step 2 (Count Frequencies): Iterate through the L input strings and count the total number of '1's in
each string.
- Step 3 (Sort by Bit Weight): Create an array `ord` representing the indices of the strings. Sort these
indices based on the count of '1's in descending order. This greedily maps the string with the most '1's
to the 0-th bit, the second most to the 1-st bit, etc.
- Step 4 (Reconstruct Values): For each column index j (from 0 to N-1), reconstruct the integer value by
placing the j-th character of the sorted strings into their respective bit positions.
- Step 5 (Validate): Sort the reconstructed N integers. If they exactly match the sequence 1, 2, ..., N,
then our greedy bit-assignment is valid. If any number is missing or out of place, no valid configuration
exists, and the answer is 0.
- Step 6 (Permutations for Duplicates): If the configuration is valid, we must account for interchangeable
strings. If multiple strings have the exact same count of '1's, they can be swapped without invalidating
the sequence. Thus, we count the frequency of each '1's count, calculate the factorial of each frequency,
and multiply them together to get the total number of valid permutations.
*/
