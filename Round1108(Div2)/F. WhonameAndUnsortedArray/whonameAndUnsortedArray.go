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

func precomp() {}

func solve() {
	n := in.rInt()
	arr := readInts(n)
	for i := range arr {
		arr[i]--
	}
	isS := true
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			isS = false
			break
		}
	}
	if isS {
		fmt.Fprintln(out, 0)
		fmt.Fprintln(out)
		return
	}
	if n == 2 {
		fmt.Fprintln(out, -1)
		return
	}
	var ans []int
	op := func(p int) {
		t1 := arr[p-1]
		for i := p - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = t1
		t2 := arr[p]
		for i := p; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[n-1] = t2
	}
	gop := func(p int) {
		ans = append(ans, p)
		op(p)
	}
	for i := 2; i < n; i++ {
		p := 0
		for j := 0; j < n; j++ {
			if arr[j] == i {
				p = j
				break
			}
		}
		if p != 0 {
			gop(p)
		} else {
			gop(2)
			gop(1)
			gop(n - 1)
		}
	}
	if arr[0] == 1 && arr[1] == 0 {
		if n%2 == 1 {
			for i := 0; i < n-2; i++ {
				gop(2)
			}
		} else {
			fmt.Fprintln(out, -1)
			return
		}
	}
	fmt.Fprintln(out, len(ans))
	printSlice(ans)
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

// Solution referenced from https://codeforces.com/profile/Nachia

/*
The algorithm constructively sorts the array by sequentially processing target values from 2 to $N-1$. For each target $i$,
it finds its current index $p$. The simulated cyclic shift operation `op(p)` works by effectively moving the element at $p-1$
to the front while cyclically sending the element at $p$ to the back of the array. By applying this specific transformation
at the target's index, the targeted element cascades structurally toward its correct suffix position. If the target is initially
stuck at index 0, three specific cascading rotations (`p=2`, `p=1`, `p=N-1`) are sequentially applied to break it loose and
reposition it without violating the invariant of previously sorted elements. Once items 2 through $N-1$ are correctly arranged,
the sequence either falls in perfectly sorted order or has 0 and 1 inverted. If inverted, the array can only be repaired if $N$
is odd, done by repeatedly applying $N-2$ consecutive rotations with $p=2$; otherwise, it is impossible and the code rejects it
(returning -1). The time complexity per testcase is $O(N^2)$ since finding each element and simulating the array shifts takes
$O(N)$ for $O(N)$ elements. The space complexity is $O(N)$ as it stores the permuted array and the recorded sequence of operations
in a scalable list.
*/
