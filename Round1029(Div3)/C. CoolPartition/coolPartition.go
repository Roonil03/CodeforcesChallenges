// The Question is solved post the contest times
package main

import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	v := make([]int, n)
	distinct := make([]int, n)
	freq := make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
		freq[v[i]]++
		if freq[v[i]] == 1 {
			distinct[i] = 1
		}
		if i > 0 {
			distinct[i] += distinct[i-1]
		}
	}
	for i := 0; i <= n; i++ {
		freq[i] = 0
	}
	ans := 0
	end := n - 1
	total := 0
	for i := n - 1; i >= 0; i-- {
		freq[v[i]]++
		if freq[v[i]] == 1 {
			total++
		}
		if total == distinct[end] {
			ans++
			for j := i; j <= end; j++ {
				freq[v[j]] = 0
			}
			end = i - 1
			total = 0
		}
	}
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for t > 0 {
		solve()
		t--
	}
}

/*
This algorithm uses a two-pass approach with a key mathematical insight. First, it computes a prefix
array distinct[i] representing the number of unique elements from index 0 to i. Then, it iterates
backwards maintaining a count of distinct elements in the current suffix. When the suffix's distinct
count equals distinct[end] (where end marks the boundary of the current segment), it means we can make
a valid cut because all elements in this segment will appear in the remaining array. The time complexity
is O(n) per test case since each element is processed at most twice, and the space complexity is O(n)
for the arrays storing distinct counts and frequencies.
*/
