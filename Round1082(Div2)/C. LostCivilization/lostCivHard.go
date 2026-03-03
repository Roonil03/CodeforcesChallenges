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

type Elm struct {
	id   int
	val  int64
	minB int
	newB int
}

func solve() {
	line := readLine()
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	arr := readInt64s(n)
	A := make([]int64, n+1)
	for i := 0; i < n; i++ {
		A[i+1] = arr[i]
	}
	st := make([]Elm, 0)
	var res int64
	for i := 1; i <= n; i++ {
		pp := int(1e9)
		for len(st) > 0 && st[len(st)-1].val >= A[i] {
			if st[len(st)-1].newB < pp {
				pp = st[len(st)-1].newB
			}
			st = st[:len(st)-1]
		}
		j := 0
		if len(st) > 0 {
			j = st[len(st)-1].id
		}
		bi := 0
		if i == 1 {
			bi = 0
		} else if A[i] > A[i-1]+1 {
			bi = 0
		} else {
			if pp == int(1e9) {
				bi = j
			} else {
				if j < pp {
					bi = j
				} else {
					bi = pp
				}
			}
		}
		res += int64(i-bi) * int64(n-i+1)
		newB := bi
		if pp != int(1e9) {
			if pp < newB {
				newB = pp
			}
		}
		st = append(st, Elm{i, A[i], bi, newB})
	}
	fmt.Fprintln(out, res)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t, _ := strconv.Atoi(readLine())
	for t > 0 {
		t--
		solve()
	}
}

/*
While processing from left to right, we maintain a monotonic stack.
Each element in the stack represents a candidate start position
for a minimal initial sequence.

We record:
- idx: current position
- val: value at position
- minB: earliest valid boundary
- newB: propagated minimal boundary

For each position i:

1. Maintain a monotonic stack by popping elements whose value
   is greater than or equal to A[i].
   While popping, track the minimum boundary encountered.

2. Let j be the index of the new stack top (or 0 if empty).

3. Compute bi:
   - If i == 1, bi = 0.
   - If A[i] > A[i-1] + 1, bi = 0.
   - Otherwise:
       If nothing was popped, bi = j.
       Else bi = min(j, smallest boundary from popped elements).

4. The contribution of position i is:
       (i - bi) * (n - i + 1)

   Add this to the answer.

5. Push the new element with its propagated boundary.

Since each element is pushed and popped at most once,
the time complexity is O(n) per test case.
*/
