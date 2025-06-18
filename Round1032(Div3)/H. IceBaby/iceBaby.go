// Couldn't be solved during contest time, has been solved post the contest hours

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var t int
	fmt.Fscan(reader, &t)
	for tc := 0; tc < t; tc++ {
		var n int
		fmt.Fscan(reader, &n)
		dp := make([]int, 0, n)
		for i := 0; i < n; i++ {
			var l, r int
			fmt.Fscan(reader, &l, &r)
			idx := sort.Search(len(dp), func(j int) bool {
				return dp[j] > r
			})
			if idx < len(dp) {
				dp = append(dp[:idx], dp[idx+1:]...)
			}
			ins := sort.Search(len(dp), func(j int) bool {
				return dp[j] >= l
			})
			dp = append(dp, 0)
			copy(dp[ins+1:], dp[ins:])
			dp[ins] = l
			fmt.Fprint(writer, len(dp), " ")
		}
		fmt.Fprint(writer, "\n")
	}
}

/*
The code implements an interval processing algorithm where for each pair (l, r), it maintains a multiset of values
by first finding and removing the smallest element greater than r (using upper_bound(r)), then inserting l. The
algorithm uses a multiset data structure (typically implemented as a balanced binary search tree like Red-Black
tree in C++) to maintain sorted order and support efficient search, insertion, and deletion operations. The logic
appears to solve an interval scheduling or activity selection variant where intervals are processed sequentially,
potentially representing some form of greedy optimization where conflicting elements (those > r) are removed before
adding new candidates (l). The time complexity is O(n log n) per test case, since each of the n pairs requires
O(log n) operations for upper_bound, erase, and insert on the multiset, resulting in O(t × n log n) overall where
t is the number of test cases. The space complexity is O(n) for storing up to n elements in the multiset at any
given time.
*/
