package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func precomp() {

}

func h1(s, m int64, n int, l, r, pref []int64) int64 {
	e := s + m - 1
	ll, hh := 0, n
	for ll < hh {
		mid := (ll + hh) >> 1
		if r[mid] >= s {
			hh = mid
		} else {
			ll = mid + 1
		}
	}
	l1 := ll
	ll, hh = 0, n
	for ll < hh {
		mid := (ll + hh) >> 1
		if l[mid] > e {
			hh = mid
		} else {
			ll = mid + 1
		}
	}
	r1 := ll
	if l1 >= r1 {
		return 0
	}
	res := pref[r1] - pref[l1]
	if l[l1] < s {
		res -= (s - l[l1])
	}
	if r[r1-1] > e {
		res -= (r[r1-1] - e)
	}
	return res
}

func solve() {
	n := in.rInt()
	// arr := readInts(n)
	m := in.rInt64()
	k := in.rInt64()
	l := make([]int64, n)
	r := make([]int64, n)
	pref := make([]int64, n+1)
	for i := range n {
		l[i] = in.rInt64()
		r[i] = in.rInt64()
		pref[i+1] = pref[i] + (r[i] - l[i] + 1)
	}
	pt := make([]int64, 0, 4*n+2)
	pt = append(pt, 1, r[n-1]+1)
	for i := range n {
		if l[i] >= 1 {
			pt = append(pt, l[i])
		}
		if r[i]+1 >= 1 {
			pt = append(pt, r[i]+1)
		}
		if l[i]-m >= 1 {
			pt = append(pt, l[i]-m)
		}
		if r[i]-m+1 >= 1 {
			pt = append(pt, r[i]-m+1)
		}
	}
	sort.Slice(pt, func(i, j int) bool {
		return pt[i] < pt[j]
	})
	temp := pt[:0]
	for i, v := range pt {
		if i == 0 || v != pt[i-1] {
			temp = append(temp, v)
		}
	}
	pt = temp
	evms := make([]int64, len(pt))
	for i, v := range pt {
		evms[i] = h1(v, m, n, l, r, pref)
	}
	for i := range len(pt) - 1 {
		v1 := evms[i]
		v2 := evms[i+1]
		if (v1 <= k && k <= v2) || (v2 <= k && k <= v1) {
			res := pt[i]
			if v1 <= k && k <= v2 {
				res += (k - v1)
			} else {
				res += (v1 - k)
			}
			fmt.Fprintln(out, res)
			return
		}
	}
	fmt.Fprintln(out, -1)
	// printSlice(arr)
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
