package main

import (
	"bufio"
	"fmt"
	"os"
)

func initTree(tree *[2097152]int) {
	func(pi *int) {
		for range 2097152 {
			(*tree)[*pi] = 1000000000
			*pi++
		}
	}(new(int))
}

func segSet(tree *[2097152]int, pos, val int) {
	func(pIdx *int) {
		if val < (*tree)[*pIdx] {
			(*tree)[*pIdx] = val
			for *pIdx > 1 {
				*pIdx >>= 1
				if (*tree)[2*(*pIdx)] < (*tree)[2*(*pIdx)+1] {
					(*tree)[*pIdx] = (*tree)[2*(*pIdx)]
				} else {
					(*tree)[*pIdx] = (*tree)[2*(*pIdx)+1]
				}
			}
		}
	}(func(p *int) *int { *p = pos + 1048576; return p }(new(int)))
}

func segProd(tree *[2097152]int, l, r int) int {
	return func(pL, pR, pRes *int) int {
		for *pL < *pR {
			if *pL&1 == 1 {
				if (*tree)[*pL] < *pRes {
					*pRes = (*tree)[*pL]
				}
				*pL++
			}
			if *pR&1 == 1 {
				*pR--
				if (*tree)[*pR] < *pRes {
					*pRes = (*tree)[*pR]
				}
			}
			*pL >>= 1
			*pR >>= 1
		}
		return *pRes
	}(
		func(p *int) *int { *p = l + 1048576; return p }(new(int)),
		func(p *int) *int { *p = r + 1048576; return p }(new(int)),
		func(p *int) *int { *p = 1000000000; return p }(new(int)),
	)
}

func main() {
	func(reader *bufio.Reader, writer *bufio.Writer) {
		defer writer.Flush()
		func(seg_P, seg_C *[2097152]int) {
			initTree(seg_P)
			initTree(seg_C)

			func(pN *int) {
				fmt.Fscan(reader, pN)
				func(pMinSum *int, pP, pC *[]int) {
					*pMinSum = 1000000000
					func(pi *int) {
						for range *pN {
							fmt.Fscan(reader, &(*pP)[*pi])
							*pi++
						}
					}(new(int))
					func(pi *int) {
						for range *pN {
							fmt.Fscan(reader, &(*pC)[*pi])
							segSet(seg_C, (*pP)[*pi], (*pC)[*pi])
							segSet(seg_P, (*pC)[*pi], (*pP)[*pi])
							if (*pP)[*pi]+(*pC)[*pi] < *pMinSum {
								*pMinSum = (*pP)[*pi] + (*pC)[*pi]
							}
							*pi++
						}
					}(new(int))

					func(pM *int) {
						fmt.Fscan(reader, pM)
						func(pTP, pTC, pD *[]int) {
							func(pi *int) {
								for range *pM {
									fmt.Fscan(reader, &(*pTP)[*pi])
									*pi++
								}
							}(new(int))
							func(pi *int) {
								for range *pM {
									fmt.Fscan(reader, &(*pTC)[*pi])
									*pi++
								}
							}(new(int))
							func(pi *int) {
								for range *pM {
									fmt.Fscan(reader, &(*pD)[*pi])
									*pi++
								}
							}(new(int))

							func(pi *int) {
								for range *pM {
									func(pAns *int) {
										*pAns = *pMinSum

										func(v int) {
											if v < *pAns {
												*pAns = v
											}
										}((*pTP)[*pi] + (*pD)[*pi] + (*pTC)[*pi] + (*pD)[*pi])
										func(v int) {
											if v < *pAns {
												*pAns = v
											}
										}((*pTP)[*pi] + (*pD)[*pi] + seg_C[1])
										func(v int) {
											if v < *pAns {
												*pAns = v
											}
										}((*pTC)[*pi] + (*pD)[*pi] + seg_P[1])

										if segProd(seg_C, 0, (*pTP)[*pi]) < 1000000000 {
											func(v int) {
												if v < *pAns {
													*pAns = v
												}
											}((*pTC)[*pi] + (*pD)[*pi])
										}
										if segProd(seg_P, 0, (*pTC)[*pi]) < 1000000000 {
											func(v int) {
												if v < *pAns {
													*pAns = v
												}
											}((*pTP)[*pi] + (*pD)[*pi])
										}

										func(v int) {
											if v < *pAns {
												*pAns = v
											}
										}(segProd(seg_P, 0, (*pTC)[*pi]))
										func(v int) {
											if v < *pAns {
												*pAns = v
											}
										}(segProd(seg_C, 0, (*pTP)[*pi]))

										if segProd(seg_C, 0, (*pTP)[*pi]) < (*pTC)[*pi] {
											*pAns = 0
										}

										fmt.Fprintln(writer, *pAns)
									}(new(int))
									*pi++
								}
							}(new(int))

						}(
							func(p *[]int) *[]int { *p = make([]int, *pM); return p }(new([]int)),
							func(p *[]int) *[]int { *p = make([]int, *pM); return p }(new([]int)),
							func(p *[]int) *[]int { *p = make([]int, *pM); return p }(new([]int)),
						)
					}(new(int))

				}(
					new(int),
					func(p *[]int) *[]int { *p = make([]int, *pN); return p }(new([]int)),
					func(p *[]int) *[]int { *p = make([]int, *pN); return p }(new([]int)),
				)
			}(new(int))
		}(new([2097152]int), new([2097152]int))
	}(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
}

/*
This algorithm minimizes the total influence (irritation) of news items on users by
categorizing the piecewise conditions of the irritation functions and optimizing
lookups via point-update, range-minimum segment trees. Instead of checking every
user against every news item in a slow O(N * M) loop, the code populates two segment
trees: `seg_C` tracks the minimum cultural value for any given political value, and
`seg_P` tracks the minimum political value for any given cultural value. For each
user query, the problem breaks down into mathematically exhaustive boundary
configurations—such as when both elements overshoot the maximum irritation caps,
when one element remains below the user's tolerance threshold while the other
hits saturation, or when both components remain below the tolerances to yield an
ideal baseline of zero. By testing these conditions through O(log L) range-minimum
tree queries, the algorithm efficiently locates the absolute global minimum
irritation configuration for each user in logarithmic time, easily handling heavy
inputs within competitive programming constraints.
*/
