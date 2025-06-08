// The Question is solved post the contest times
package main

import (
	"bufio"
	"os"
	"strconv"
)

var scanner *bufio.Scanner
var writer *bufio.Writer

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	writer = bufio.NewWriter(os.Stdout)
}

func nextInt() int {
	scanner.Scan()
	val, _ := strconv.Atoi(scanner.Text())
	return val
}

func solve() {
	n := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		b[i] = nextInt()
	}
	if a[n-1] == b[n-1] {
		writer.WriteString(strconv.Itoa(n))
		writer.WriteByte('\n')
		return
	}
	seen := make([]bool, n+1)
	for i := n - 2; i >= 0; i-- {
		ai, bi := a[i], b[i]
		ai1, bi1 := a[i+1], b[i+1]
		if ai == bi || ai == ai1 || bi == bi1 || seen[ai] || seen[bi] {
			writer.WriteString(strconv.Itoa(i + 1))
			writer.WriteByte('\n')
			return
		}
		seen[ai1] = true
		seen[bi1] = true
	}
	writer.WriteByte('0')
	writer.WriteByte('\n')
}

func main() {
	defer writer.Flush()
	t := nextInt()
	for t > 0 {
		solve()
		t--
	}
}

/*
This algorithm uses a backwards greedy approach that identifies the optimal removal position by checking
when value propagation through operations becomes beneficial. It starts from the end and moves backwards,
tracking which values have been encountered to prevent conflicts. The key insight is that at any position,
matches can be achieved if elements already match, adjacent elements can copy values, or if values haven't
been used in conflicting ways. Time Complexity: O(n) for the single backward pass through arrays. Space
Complexity: O(n) for the seen array and input storage, making this solution highly efficient compared to
the exponential recursive approach.
*/
