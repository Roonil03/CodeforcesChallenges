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

const MOD = 1000000007

func readLine() string {
	line, err := in.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

func readInts() []int {
	line := readLine()
	for line == "" {
		line = readLine()
	}
	parts := strings.Fields(line)
	res := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func bexpo(b, e int64) int64 {
	var ans int64 = 1
	b %= MOD
	for e > 0 {
		if e&1 == 1 {
			ans = (ans * b) % MOD
		}
		e >>= 1
		b = (b * b) % MOD
	}
	return ans
}

func inv(n int64) int64 {
	return bexpo(n, MOD-2)
}

func solve() {
	line := readLine()
	for line == "" {
		line = readLine()
	}
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}
	n, _ := strconv.Atoi(parts[0])
	line2 := readLine()
	parts2 := strings.Fields(line2)
	x, _ := strconv.Atoi(parts2[0])
	y, _ := strconv.Atoi(parts2[1])
	x--
	y--
	a := readInts()
	b := readInts()
	if x > y {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
			b[i], b[j] = b[j], b[i]
		}
		x = n - x - 1
		y = n - y - 1
	}
	b[x] = 0
	b[y] = 0
	pb := make([]int64, n+1)
	for i := range n {
		pb[i+1] = pb[i] + int64(b[i])
	}
	inv2 := inv(2)
	adp := make([][]int64, n)
	for i := range adp {
		adp[i] = make([]int64, n)
	}
	adp[x][x] = 1
	for d := 1; d < n; d++ {
		for r := d; r < n; r++ {
			l := r - d
			if x < l || x > r || (y >= l && y <= r) {
				continue
			}
			sizeWithoutR := int64(a[x]) + pb[r] - pb[l]
			if sizeWithoutR >= int64(a[r]) {
				val := adp[l][r-1]
				if l != 0 && sizeWithoutR >= int64(a[l-1]) {
					adp[l][r] = (adp[l][r] + val*inv2) % MOD
				} else {
					adp[l][r] = (adp[l][r] + val) % MOD
				}
			}
			sizeWithoutL := int64(a[x]) + pb[r+1] - pb[l+1]
			if sizeWithoutL >= int64(a[l]) {
				val := adp[l+1][r]
				if r != n-1 && sizeWithoutL >= int64(a[r+1]) {
					adp[l][r] = (adp[l][r] + val*inv2) % MOD
				} else {
					adp[l][r] = (adp[l][r] + val) % MOD
				}
			}
		}
	}
	bdp := make([][]int64, n)
	for i := range bdp {
		bdp[i] = make([]int64, n)
	}
	bdp[y][y] = 1
	for d := 1; d < n; d++ {
		for r := d; r < n; r++ {
			l := r - d
			if y < l || y > r || (x >= l && x <= r) {
				continue
			}
			sizeWithoutR := int64(a[y]) + pb[r] - pb[l]
			if sizeWithoutR >= int64(a[r]) {
				val := bdp[l][r-1]
				if l != 0 && sizeWithoutR >= int64(a[l-1]) {
					bdp[l][r] = (bdp[l][r] + val*inv2) % MOD
				} else {
					bdp[l][r] = (bdp[l][r] + val) % MOD
				}
			}
			sizeWithoutL := int64(a[y]) + pb[r+1] - pb[l+1]
			if sizeWithoutL >= int64(a[l]) {
				val := bdp[l+1][r]
				if r != n-1 && sizeWithoutL >= int64(a[r+1]) {
					bdp[l][r] = (bdp[l][r] + val*inv2) % MOD
				} else {
					bdp[l][r] = (bdp[l][r] + val) % MOD
				}
			}
		}
	}
	asum := make([][]int64, n)
	for i := range asum {
		asum[i] = make([]int64, n)
	}
	for l := range n {
		for r := l; r < n; r++ {
			asum[r-l][r] = adp[l][r]
		}
	}
	for diff := range n {
		for mxr := 1; mxr < n; mxr++ {
			asum[diff][mxr] = (asum[diff][mxr] + asum[diff][mxr-1]) % MOD
		}
	}
	var ans int64 = 0
	for l := 2; l < n; l++ {
		for r := l; r < n; r++ {
			bsz := int64(a[y]) + pb[r+1] - pb[l]
			if l != 0 && bsz >= int64(a[l-1]) {
				continue
			}
			if r != n-1 && bsz >= int64(a[r+1]) {
				continue
			}
			ans = (ans + asum[r-l+1][l-2]*bdp[l][r]) % MOD
		}
	}
	dp := make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, n)
	}
	if y == x+1 {
		dp[x][y] = 1
	}
	for diff := y - x; diff < n; diff++ {
		for r := y; r < n; r++ {
			l := r - diff
			if l < 0 || l > x || r < y {
				continue
			}
			aend := (l + r) / 2
			if l%2 == r%2 {
				asz := int64(a[x]) + pb[aend+1] - pb[l+1]
				bsz := int64(a[y]) + pb[r+1] - pb[aend+1]
				if asz >= int64(a[l]) {
					if asz >= bsz {
						dp[l][r] = (dp[l][r] + dp[l+1][r]*inv2) % MOD
					} else {
						dp[l][r] = (dp[l][r] + dp[l+1][r]) % MOD
					}
				}
				if aend != 0 {
					chance := (adp[l][aend-1] * bdp[aend+1][r]) % MOD
					asz_mid := int64(a[x]) + pb[aend] - pb[l]
					if asz_mid >= int64(a[aend]) {
						if l != 0 && asz_mid >= int64(a[l-1]) {
							dp[l][r] = (dp[l][r] + chance*inv2) % MOD
						} else {
							dp[l][r] = (dp[l][r] + chance) % MOD
						}
					}
				}
			} else {
				asz := int64(a[x]) + pb[aend+1] - pb[l]
				bsz := int64(a[y]) + pb[r] - pb[aend+1]
				if bsz >= int64(a[r]) {
					if bsz >= asz {
						dp[l][r] = (dp[l][r] + dp[l][r-1]*inv2) % MOD
					} else {
						dp[l][r] = (dp[l][r] + dp[l][r-1]) % MOD
					}
				}
				if aend+2 < n {
					chance := (adp[l][aend] * bdp[aend+2][r]) % MOD
					bsz_mid := int64(a[y]) + pb[r+1] - pb[aend+2]
					if bsz_mid >= int64(a[aend+1]) {
						if r != n-1 && bsz_mid >= int64(a[r+1]) {
							dp[l][r] = (dp[l][r] + chance*inv2) % MOD
						} else {
							dp[l][r] = (dp[l][r] + chance) % MOD
						}
					}
				}
			}
			asz_fin := int64(a[x]) + pb[aend+1] - pb[l]
			bsz_fin := int64(a[y]) + pb[r+1] - pb[aend+1]
			if l%2 == r%2 {
				if bsz_fin < asz_fin {
					if r == n-1 || int64(a[r+1]) > bsz_fin {
						ans = (ans + dp[l][r]) % MOD
					}
				}
			} else {
				if asz_fin >= bsz_fin {
					if l != 0 && asz_fin >= int64(a[l-1]) {
						ans = (ans + inv2*dp[l][r]) % MOD
					} else {
						ans = (ans + dp[l][r]) % MOD
					}
				}
			}
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	tStr := readLine()
	if tStr != "" {
		t, _ := strconv.Atoi(tStr)
		for t > 0 {
			t--
			solve()
		}
	}
}

/*The problem asks for the probability that Alice's fish ($x$) survives and Bob's fish ($y$) is eliminated, given they grow by eating neighbors and can be eaten by their environment. The solution uses a **Two-Phase Dynamic Programming** approach in $O(N^2)$.

---

## Phase 1: Independent Growth
In this phase, Alice and Bob are separated by other fish. They grow independently by eating fish to their left or right.

* **State**: `adp[l][r]` is the probability that Alice has successfully eaten all fish in the range $[l, r]$.
* **Expansion**: A fish can expand from $[l, r]$ to $[l-1, r]$ or $[l, r+1]$ if its current size (initial $a_i +$ sum of $b_j$ in its range) is $\ge$ the size of the target neighbor.
* **Probabilities**: If a fish can eat both neighbors, it chooses one with $1/2$ probability (`inv2`). If only one is edible, it eats that one with $1/1$ probability.
* **Constraint**: This phase stops for a fish if its next move would involve interacting with the other player's fish.



---

## Phase 2: Environmental Elimination (Pre-Meeting)
It is possible for Bob to lose before he even encounters Alice. This happens if Bob becomes "stuck."

* **The "Stuck" Condition**: A fish is eliminated by the environment if it cannot eat its left neighbor AND cannot eat its right neighbor.
* **Calculation**: We use `asum` (a prefix sum of Alice's probabilities) to check: "What is the probability Bob is stuck at range $[l, r]$ while Alice is still safely growing somewhere else?"
* **Winning**: If Bob is stuck and Alice is still alive, this probability is added directly to Alice's win total.

---

## Phase 3: The Interaction (Fish Fight)
When Alice's range and Bob's range become adjacent, they enter a unified DP state `dp[l][r]`, where the interval $[l, r]$ represents the total territory consumed by both fish combined.

* **Turn Parity**: Turns are determined by the number of fish eaten. Alice moves on turns $0, 2, 4 \dots$ and Bob moves on turns $1, 3, 5 \dots$.
* **The Conflict**:
    * On **Alice's turn**, she tries to eat a neighbor or Bob. If her size is $\ge$ Bob's size, she has a chance to win.
    * On **Bob's turn**, he tries to eat a neighbor or Alice. If Alice is larger or Bob is stuck, Alice moves closer to winning.
* **Midpoint Logic**: The variable `aend` helps determine the boundary between Alice's territory and Bob's territory within the combined range $[l, r]$.



---

## Complexity & Constraints
* **Time Complexity**: $O(N^2)$ because we iterate through all possible subarrays $[l, r]$. With $N=3000$, $N^2 = 9 \times 10^6$, which fits within the 2.0s time limit.
* **Space Complexity**: $O(N^2)$ for the DP tables.
* **Modular Arithmetic**: All additions and multiplications are performed modulo $10^9+7$. Division is handled via the Modular Multiplicative Inverse (Fermat's Little Theorem).
*/
