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

	inv []int64
	cat []int64

	mod int64 = int64(1e9 + 7)
)

func readLine() string {
	line, err := in.ReadString('\n')
	if err != nil && len(line) == 0 {
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

type Query struct {
	n, k int
}

func precompute(maxN int) {
	inv = make([]int64, maxN+3)
	cat = make([]int64, maxN+1)
	inv[1] = 1
	for i := 2; i <= maxN+2; i++ {
		inv[i] = mod - (mod/int64(i))*inv[mod%int64(i)]%mod
	}
	cat[0] = 1
	for i := 0; i < maxN; i++ {
		cat[i+1] = cat[i] * int64(2*(2*i+1)) % mod * inv[i+2] % mod
	}
}

func buildPrefixSums(m, r int) []int64 {
	res := make([]int64, m+1)
	if r < 0 {
		return res
	}
	p := int64(1)
	res[0] = 1
	var c int64
	if r == 0 {
		c = 1
	}
	for a := 0; a < m; a++ {
		p = (2*p - c) % mod
		if p < 0 {
			p += mod
		}
		res[a+1] = p
		ap1 := a + 1
		if ap1 < r {
			c = 0
		} else if ap1 == r {
			c = 1
		} else {
			c = c * int64(ap1) % mod * inv[ap1-r] % mod
		}
	}
	return res
}

func solve(n, k int) {
	m := n - 1
	pk := buildPrefixSums(m, k)
	pk2 := buildPrefixSums(m, k-2)
	var ans int64
	for a := 0; a <= m; a++ {
		b := m - a
		shapeWays := cat[a] * cat[b] % mod
		labelWays := (pk[a]*pk[b] - pk2[a]*pk2[b]) % mod
		if labelWays < 0 {
			labelWays += mod
		}
		ans += shapeWays * labelWays % mod
		if ans >= mod {
			ans -= mod
		}
	}
	fmt.Fprintln(out, ans)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t, _ := strconv.Atoi(readLine())
	queries := make([]Query, t)
	maxN := 0
	for i := 0; i < t; i++ {
		v := readInts(2)
		queries[i] = Query{n: v[0], k: v[1]}
		if v[0] > maxN {
			maxN = v[0]
		}
	}
	precompute(maxN)
	for _, q := range queries {
		solve(q.n, q.k)
	}
}

/*
This works by converting each subtree labeling into a unique set of downward root-to-node prefix flips:
if y_u = x_u xor x_left xor x_right, then y_u = 1 means “take the prefix ending at u”, so for any fixed
subtree shape the number of labelings with exactly p such prefixes is just C(size, p). For the whole
tree, if the left subtree needs p prefixes and the right subtree needs q, then one root-passing operation
can pair at most one left prefix with at most one right prefix, so the minimum number of operations is
the smallest t >= max(p, q) with the correct root parity; that means cost k happens exactly when max(p, q)
 is either k or k-1, which becomes P_a(k)P_b(k) - P_a(k-2)P_b(k-2) for subtree sizes a, b, where
 P_s(r)=sum_{i=0}^r C(s,i). Then we sum this over all root splits a+b=n-1, multiplying by the number of
 ordered binary tree shapes Cat[a] * Cat[b]. The code precomputes Catalan numbers once, and for each test
 builds all needed P_s(k) values in linear time using P_{a+1}(r)=2P_a(r)-C(a,r), so the total complexity
 is O(maxN + sum n).
*/
