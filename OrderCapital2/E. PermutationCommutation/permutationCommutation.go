package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tt int
	if _, err := fmt.Fscan(reader, &tt); err != nil {
		return
	}
	for tt > 0 {
		tt--
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
			a[i]--
		}
		b := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
			if b[i] > 0 {
				b[i]--
			}
		}
		was := make([]bool, n)
		used := make([]bool, n)
		fail := false
		var toFill [][]int
		for i := 0; i < n; i++ {
			if was[i] {
				continue
			}
			var cyc []int
			x := i
			for !was[x] {
				cyc = append(cyc, x)
				was[x] = true
				x = a[x]
			}
			sz := len(cyc)
			found := false
			for j := 0; j < sz; j++ {
				if b[cyc[j]] != -1 {
					found = true
					seq := []int{b[cyc[j]]}
					for k := 0; k < sz; k++ {
						seq = append(seq, a[seq[len(seq)-1]])
						if used[seq[len(seq)-1]] {
							fail = true
							break
						}
						used[seq[len(seq)-1]] = true
					}
					if fail {
						break
					}
					if seq[0] != seq[len(seq)-1] {
						fail = true
						break
					}
					for k := 0; k < sz; k++ {
						idx := cyc[(j+k)%sz]
						if b[idx] != -1 && b[idx] != seq[k] {
							fail = true
							break
						}
						b[idx] = seq[k]
					}
					break
				}
			}
			if fail {
				break
			}
			if !found {
				toFill = append(toFill, cyc)
			}
		}
		if fail {
			fmt.Fprint(writer, "NO\n")
			continue
		}
		fmt.Fprint(writer, "YES\n")
		fillWith := make([][][]int, n+1)
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			var cyc []int
			x := i
			for !used[x] {
				cyc = append(cyc, x)
				used[x] = true
				x = a[x]
			}
			fillWith[len(cyc)] = append(fillWith[len(cyc)], cyc)
		}
		ptr := make([]int, n+1)
		for _, cyc := range toFill {
			sz := len(cyc)
			seq := fillWith[sz][ptr[sz]]
			ptr[sz]++
			for j := 0; j < sz; j++ {
				b[cyc[j]] = seq[j]
			}
		}
		for i := 0; i < n; i++ {
			if i == n-1 {
				fmt.Fprint(writer, b[i]+1, "\n")
			} else {
				fmt.Fprint(writer, b[i]+1, " ")
			}
		}
	}
}

/*
This algorithm solves a permutation reconstruction and validation problem by analyzing the functional cycle
structure of a given permutation graph A. It decomposes the graph into disjoint functional cycles and processes
each one independently; if a cycle contains any element whose corresponding mapping in B is already pre-determined,
the mapping rule for the entire cycle becomes structurally fixed and is uniquely propagated along A's cycle path.
The algorithm thoroughly validates these pre-filled assignments for consistency and tracking contradictions, marking
matching elements as used. Any completely unassigned cycles are gathered, categorized into buckets by their lengths,
and sequentially matched with remaining unused cycle components of identical sizes to fully determine the missing
elements of permutation B. The time complexity is O(N) per testcase because every element is processed a constant
number of times during cycle extraction, validation, and matching phases. The space complexity is O(N) due to the
linear allocation of state tracking slices and dynamic buckets used to temporarily hold the permutation cycles.
*/
