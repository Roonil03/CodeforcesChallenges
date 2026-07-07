package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func precomp() {

}

func solve() {
	l := in.rInt()
	r := in.rInt()
	n := in.rInt()
	var lBit, rBit int
	for i := 29; i >= 0; i-- {
		if (1<<i)&l != 0 {
			lBit = i + 1
			break
		}
	}
	for i := 29; i >= 0; i-- {
		if (1<<i)&r != 0 {
			rBit = i + 1
			break
		}
	}
	var maxLen int
	for a := lBit + 1; a <= rBit; a++ {
		for b := a; b <= rBit; b++ {
			if curr := lcm(a, b); curr > maxLen {
				maxLen = curr
			}
		}
	}
	var v []string
	if maxLen > 0 {
		var tmp strings.Builder
		for i := 0; i < n; i++ {
			if i%maxLen == 0 {
				tmp.WriteByte('1')
			} else {
				tmp.WriteByte('0')
			}
		}
		v = append(v, tmp.String())
	}
	{
		val := 1 << (lBit - 1)
		var pattern strings.Builder
		pattern.WriteByte('1')
		for i := lBit - 2; i >= 0; i-- {
			if (1<<i)&l != 0 {
				val += 1 << i
				pattern.WriteByte('1')
			} else {
				if val+(1<<i) > r {
					pattern.WriteByte('0')
				} else {
					for pattern.Len() < lBit {
						pattern.WriteByte('0')
					}
					break
				}
			}
		}
		pStr := pattern.String()
		pLen := len(pStr)
		var tmp strings.Builder
		for i := 0; i < n; i++ {
			tmp.WriteByte(pStr[i%pLen])
		}
		v = append(v, tmp.String())
	}
	for a := lBit + 1; a <= rBit; a++ {
		var tmp strings.Builder
		assign := make([]int, lBit)
		for i := range assign {
			assign[i] = -1
		}
		assign[0] = 1
		chk := func(bit int) bool {
			var tot int64
			for i := lBit - 1; i >= 0; i-- {
				aB := assign[lBit-1-i]
				boundBit := ((1 << i) & l) > 0
				if aB == -1 {
					if i == lBit-1-bit {
						if boundBit {
							return false
						}
					} else {
						if !boundBit && tot+int64(1<<i) < int64(r) {
							return true
						} else {
							tot += int64(1 << i)
						}
					}
				} else {
					var bbInt int
					if boundBit {
						bbInt = 1
					}
					if bbInt < aB {
						return true
					} else if bbInt > aB {
						return false
					}
					tot += int64(aB) << i
				}
			}
			if tot >= int64(l) && tot <= int64(r) {
				return true
			}
			return false
		}
		for i := 0; i < n; i += a {
			idx := i % lBit
			if assign[idx] != -1 {
				tmp.WriteByte(byte('0' + assign[idx]))
			} else {
				if chk(idx) {
					assign[idx] = 0
					tmp.WriteByte('0')
				} else {
					assign[idx] = 1
					tmp.WriteByte('1')
				}
			}
			for j := i + 1; j < i+a && j < n; j++ {
				tmp.WriteByte('0')
			}
		}
		v = append(v, tmp.String())
	}
	best := v[0]
	for _, s := range v {
		if s < best {
			best = s
		}
	}
	fmt.Fprintln(out, best)
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
