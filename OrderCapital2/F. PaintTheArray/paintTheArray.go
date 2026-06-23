package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var tt int
	if _, err := fmt.Fscan(in, &tt); err != nil {
		return
	}
	for tt > 0 {
		tt--
		var n, m int
		fmt.Fscan(in, &n, &m)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			a[i]--
		}
		const inf = int(1e9)
		seq := make([]int, n+m)
		for i := range seq {
			seq[i] = -inf
		}
		bound := -inf
		mx := 0
		start := n
		for i := 0; i < n; i++ {
			bound = max(bound, seq[start+m-1])
			start -= 1
			seq[start] = mx
			seq[start+a[i]] = max(seq[start+a[i]], bound) + 1
			mx = max(mx, seq[start+a[i]])
		}
		ans := max(bound, seq[m-1])
		fmt.Fprintln(out, n-ans)
	}
}

/*
This algorithm employs a highly optimized Dynamic Programming approach designed to find the maximum length of a
valid subsequence restricted by a sliding value constraint or window boundary of size m. To avoid an explicit
O(m) array-shifting operation at each step, the algorithm cleverly decrements a 'start' pointer, effectively
simulating a rightward shift of the state table in O(1) time. The 'bound' variable tracks the historical maximum
of values shifting out of the active window frame, while the 'seq' slice maintains the maximum subsequence length
obtainable for specific value offsets. By calculating the maximum valid subsequence length 'ans', the minimum
operations required is resolved via 'n - ans'. The Time Complexity is O(n) per test case since the array is processed
in a single linear pass with constant-time state transitions, and the Space Complexity is O(n + m) to allocate the
linear memory space needed for the shifting sequence buffer and the input element slice.
*/
