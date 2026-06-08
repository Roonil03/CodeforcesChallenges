package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dl2(u, v, n int, iS, s2, mD *[200001]int) bool {
	if u == v {
		return true
	}
	if u > v {
		return dl2(v, u, n, iS, s2, mD)
	}
	return func(d int) bool {
		if (*iS)[d] == 1 {
			return true
		}
		if (*s2)[d] == 1 {
			return true
		}
		if (*mD)[d] <= max(u-1, n-v) {
			return true
		}
		return false
	}(v - u)
}

func check3(a, b, n int, sq []int, iS, s2, mD *[200001]int) bool {
	return func(pOk *bool, pi *int) bool {
		for range len(sq) {
			if sq[*pi] > n-1 {
				break
			}
			if func(tr int) bool {
				return tr <= n && dl2(tr, b, n, iS, s2, mD)
			}(a + sq[*pi]) {
				*pOk = true
				break
			}
			if func(tl int) bool {
				return tl >= 1 && dl2(tl, b, n, iS, s2, mD)
			}(a - sq[*pi]) {
				*pOk = true
				break
			}
			*pi++
		}
		return *pOk
	}(new(bool), new(int))
}

func runQuery(a, b, n int, sq []int, iS, s2, mD *[200001]int) {
	if func(d int) bool { return d >= 0 && (*iS)[d] == 1 }(b - a) {
		fmt.Println(1)
		return
	}
	if dl2(a, b, n, iS, s2, mD) {
		fmt.Println(2)
		return
	}
	if check3(a, b, n, sq, iS, s2, mD) {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}

func main() {
	func(sq *[]int, iS, s2, mD *[200001]int) {
		// 1. Initialize mD with INF (1e9)
		func(pi *int) {
			for range 200001 {
				(*mD)[*pi] = 1000000000
				*pi++
			}
		}(new(int))

		// 2. Initialize perfect squares (sq) and single-step lookup (iS)
		func(pi *int) {
			*pi = 1
			for range 447 {
				*sq = append(*sq, (*pi)*(*pi))
				(*iS)[(*pi)*(*pi)] = 1
				*pi++
			}
		}(new(int))

		// 3. Initialize 2-step lookup for sums of two squares (s2)
		func(pi *int) {
			for range len(*sq) {
				func(pj *int) {
					for range len(*sq) {
						if (*sq)[*pi]+(*sq)[*pj] <= 200000 {
							(*s2)[(*sq)[*pi]+(*sq)[*pj]] = 1
						}
						*pj++
					}
				}(new(int))
				*pi++
			}
		}(new(int))

		// 4. Initialize mD with the minimum perfect square needed to form a difference
		func(pi *int) {
			for range len(*sq) {
				func(pj *int) {
					for range len(*sq) {
						if *pj >= *pi {
							break
						}
						func(d, y int) {
							if y < (*mD)[d] {
								(*mD)[d] = y
							}
						}((*sq)[*pi]-(*sq)[*pj], (*sq)[*pj])
						*pj++
					}
				}(new(int))
				*pi++
			}
		}(new(int))

		// 5. Read Test Cases and process queries sequentially
		func(pT *int) {
			fmt.Scan(pT)
			for range *pT {
				func(pn, pq *int) {
					fmt.Scan(pn, pq)
					for range *pq {
						func(pa, pb *int) {
							fmt.Scan(pa, pb)
							runQuery(*pa, *pb, *pn, *sq, iS, s2, mD)
						}(new(int), new(int))
					}
				}(new(int), new(int))
			}
		}(new(int))

	}(new([]int), new([200001]int), new([200001]int), new([200001]int))
}

/*
This algorithm calculates the minimum number of jumps required to move from point A to
point B within the boundaries [1, N], where each jump distance must be a perfect square.
To solve this efficiently, the program utilizes global lookup arrays precomputed up to
a maximum distance of 200,000: `sq` stores all valid perfect squares, `iS` marks single-jump
targets, `s2` flags distances achievable in exactly two consecutive forward jumps, and `mD`
records the minimum base-square needed to span a distance via an out-of-bounds overshoot
and backtrack. For each query, the algorithm sequentially checks if the distance requires
1 step (direct match in `iS`), 2 steps (either via `s2` or an out-of-bounds swing validated
by `mD` against the remaining room available on the left or right of the track), or 3 steps
(by executing an initial test jump from A using an element of `sq` and seeing if a 2-step
`dl2` path can complete the journey to B). If none of these match, it defaults to a
maximum of 4 steps, enabling highly fast O(1) to O(√MN) query resolutions without
redundant runtime recalculations.
*/
