package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = NewFastScanner()
	out = bufio.NewWriter(os.Stdout)
)

type FastScanner struct {
	r *bufio.Reader
}

func NewFastScanner() *FastScanner {
	return &FastScanner{r: bufio.NewReaderSize(os.Stdin, 1<<20)}
}

func (fs *FastScanner) rInt() int {
	sign, val := 1, 0
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		if err != nil {
			return 0
		}
		c, err = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return sign * val
}

func (fs *FastScanner) rInt64() int64 {
	sign, val := int64(1), int64(0)
	c, err := fs.r.ReadByte()
	for (c < '0' || c > '9') && c != '-' {
		if err != nil {
			return 0
		}
		c, err = fs.r.ReadByte()
	}
	if c == '-' {
		sign = -1
		c, _ = fs.r.ReadByte()
	}
	for c >= '0' && c <= '9' {
		val = val*10 + int64(c-'0')
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return sign * val
}

func (fs *FastScanner) rString() string {
	buf := make([]byte, 0, 16)
	c, err := fs.r.ReadByte()
	for c == ' ' || c == '\n' || c == '\r' || c == '\t' {
		if err != nil {
			return ""
		}
		c, err = fs.r.ReadByte()
	}
	for c != ' ' && c != '\n' && c != '\r' && c != '\t' {
		buf = append(buf, c)
		c, err = fs.r.ReadByte()
		if err != nil {
			break
		}
	}
	return string(buf)
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxf32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func maxf64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minf32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func minf64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func absf32(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func absf64(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func readInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = in.rInt()
	}
	return a
}

func readInt64s(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = in.rInt64()
	}
	return a
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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

var (
	lp  [1 << 20]int
	hsh [1 << 20]uint64
)

func precomp() {
	var primes []int
	for i := 2; i < 1<<20; i++ {
		if lp[i] == 0 {
			primes = append(primes, i)
			lp[i] = i
		}
		for j := 0; j < len(primes) && primes[j]*i < 1<<20; j++ {
			lp[primes[j]*i] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}
	seed := uint64(123456789)
	for i := range 1 << 20 {
		seed ^= seed << 13
		seed ^= seed >> 7
		seed ^= seed << 17
		hsh[i] = seed
	}
}

func solve() {
	n := in.rInt()
	a := make([]uint64, n)
	for i := range n {
		x := in.rInt()
		for x > 1 {
			a[i] ^= hsh[lp[x]]
			x /= lp[x]
		}
	}
	b := make([]uint64, n)
	copy(b, a)
	for i := 1; i < n; i++ {
		b[i] ^= b[i-1]
	}
	cnt := make(map[uint64]int)
	for i := range n {
		cnt[a[i]]++
	}
	var ans int64
	for i := range n {
		ans += int64(cnt[b[i]])
	}
	fmt.Fprintln(out, ans)
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	precomp()
	t := in.rInt()
	for t > 0 {
		t--
		solve()
	}
}

// Solution by https://codeforces.com/profile/abc864197532

/*
The algorithm uses Zobrist-like hashing combined with prime factorization to efficiently match array
conditions. During precomputation, it computes the lowest prime factor for every number up to 10^6
using a linear sieve and assigns a pseudo-random 64-bit integer to each possible prime number. For
each testcase, it processes the input array by transforming every integer into the XOR sum of the
hashes of its prime factors (with multiplicity), storing these mapped values in array `a`. It then
creates a prefix XOR array `b` from `a`. The problem evaluates combinations by checking how many
elements in the prefix array `b` exactly match an element's hash in the factored array `a`. It
resolves this by counting occurrences of each hash in `a` using a frequency map, then aggregating
those counts over all values present in `b`.

Time Complexity: O(M + Σ(N log(max_val))) where M = 10^6 for the linear sieve and PRNG hashing.
For each test case, factoring takes logarithmic time per element, and map lookups/inserts take
average O(1) time. Overall time is dominated by I/O and factorization.

Space Complexity: O(M + N) where O(M) is required globally for the sieve (`lp`) and hash (`hsh`)
arrays up to 10^6, and O(N) auxiliary space per test case for storing the current array, its prefix
sums, and the frequency map.
*/
