package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	func(reader *bufio.Reader, writer *bufio.Writer) {
		defer writer.Flush()
		func(pN *int) {
			if fmt.Fscan(reader, pN); true {
				*pN++
				func(par, ans *[]int) {
					func(pi *int) {
						for range *pN - 1 {
							fmt.Fscan(reader, &(*par)[*pi])
							(*par)[*pi]--
							*pi++
						}
					}(func(p *int) *int { *p = 1; return p }(new(int)))

					func(pb *int) {
						for {
							if func(ps, pe *int) bool {
								for *ps+1 < *pe {
									func(pm *int) {
										*pm = (*ps + *pe) / 2
										if func(V int) bool {
											return func(ch *[][]int, dp_down, dp_up *[]int, pOk *bool) bool {
												func(pi *int) {
													for range V - 1 {
														(*ch)[(*par)[*pi]] = append((*ch)[(*par)[*pi]], *pi)
														*pi++
													}
												}(func(p *int) *int { *p = 1; return p }(new(int)))

												func(pv *int) {
													for range V {
														func(pm1, pm2 *int) {
															func(pwIdx *int) {
																for range len((*ch)[*pv]) {
																	func(w int) {
																		if (*dp_down)[w] > *pm2 {
																			*pm2 = (*dp_down)[w]
																		}
																		if *pm2 > *pm1 {
																			*pm1, *pm2 = *pm2, *pm1
																		}
																	}((*ch)[*pv][*pwIdx])
																	*pwIdx++
																}
															}(new(int))
															(*dp_down)[*pv] = *pm2 + 1
														}(func(p *int) *int { *p = -1; return p }(new(int)), func(p *int) *int { *p = -1; return p }(new(int)))
														*pv--
													}
												}(func(p *int) *int { *p = V - 1; return p }(new(int)))

												func(pv *int) {
													for range V {
														func(pm1, pm2, pm3 *int) {
															*pm1 = (*dp_up)[*pv]
															func(pwIdx *int) {
																for range len((*ch)[*pv]) {
																	func(w int) {
																		if (*dp_down)[w] > *pm3 {
																			*pm3 = (*dp_down)[w]
																		}
																		if *pm3 > *pm2 {
																			*pm2, *pm3 = *pm3, *pm2
																		}
																		if *pm2 > *pm1 {
																			*pm1, *pm2 = *pm2, *pm1
																		}
																	}((*ch)[*pv][*pwIdx])
																	*pwIdx++
																}
															}(new(int))

															func(pwIdx *int) {
																for range len((*ch)[*pv]) {
																	func(w int) {
																		if (*dp_down)[w] == *pm1 || (*dp_down)[w] == *pm2 {
																			(*dp_up)[w] = *pm3 + 1
																		} else {
																			(*dp_up)[w] = *pm2 + 1
																		}
																	}((*ch)[*pv][*pwIdx])
																	*pwIdx++
																}
															}(new(int))

															if *pm2+1 >= *pb {
																*pOk = true
															}
														}(func(p *int) *int { *p = -1; return p }(new(int)), func(p *int) *int { *p = -1; return p }(new(int)), func(p *int) *int { *p = -1; return p }(new(int)))
														*pv++
													}
												}(new(int))

												func(pv *int) {
													for range V {
														if (*dp_down)[*pv] >= *pb {
															*pOk = true
														}
														*pv++
													}
												}(new(int))

												return *pOk
											}(
												func(p *[][]int) *[][]int { *p = make([][]int, V); return p }(new([][]int)),
												func(p *[]int) *[]int {
													*p = make([]int, V)
													func(pi *int) {
														for range V {
															(*p)[*pi] = -1
															*pi++
														}
													}(new(int))
													return p
												}(new([]int)),
												func(p *[]int) *[]int {
													*p = make([]int, V)
													func(pi *int) {
														for range V {
															(*p)[*pi] = -1
															*pi++
														}
													}(new(int))
													return p
												}(new([]int)),
												new(bool),
											)
										}(*pm) {
											*pe = *pm
										} else {
											*ps = *pm
										}
									}(new(int))
								}
								if *pe == *pN+1 {
									return true
								}
								func(pj *int) {
									for range *pN - *pe + 1 {
										(*ans)[*pj]++
										*pj++
									}
								}(func(p *int) *int { *p = *pe; return p }(new(int)))
								return false
							}(new(int), func(p *int) *int { *p = *pN + 1; return p }(new(int))) {
								break
							}
							*pb++
						}
					}(new(int))

					func(pi *int) {
						for range *pN - 1 {
							if *pi == *pN {
								fmt.Fprint(writer, (*ans)[*pi], "\n")
							} else {
								fmt.Fprint(writer, (*ans)[*pi], " ")
							}
							*pi++
						}
					}(func(p *int) *int { *p = 2; return p }(new(int)))

				}(
					func(p *[]int) *[]int { *p = make([]int, *pN+1); return p }(new([]int)),
					func(p *[]int) *[]int { *p = make([]int, *pN+1); return p }(new([]int)),
				)
			}
		}(new(int))
	}(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
}

/*
This algorithm determines a path configuration property across prefix-induced subgraphs of a tree
by combining tree DP (rerooting) with binary search. For each target path bound 'b', the algorithm
binary searches for the minimum prefix size 'm' (containing nodes 0 to m-1) necessary to satisfy
a path length condition. Inside the binary search, it builds a tree out of the active prefix nodes,
executing a two-pass dynamic programming traversal: the first pass (`dp_down`) computes the
second-longest path descending from each node, while the second pass (`dp_up`) uses a rerooting
strategy to maintain the top three path choices among neighbors to compute the second-longest
path ascending into parents and siblings. If any valid path configuration achieves a length of at
least 'b' within the size limit, the binary search settles on a minimal index 'e', and the algorithm
increments the answer metrics for all subgraphs of size 'e' through 'N' in a suffix-update loop,
resolving the cumulative query properties efficiently.
*/
