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

var f [1000001]int
var g [1000001]int

func precomp() {
	for i := 1; i <= 1000000; i++ {
		sum := 0
		temp := i
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		f[i] = sum
	}
	for i := 1; i <= 1000000; i++ {
		if i <= 9 {
			g[i] = i
		} else {
			g[i] = f[i] + g[f[i]]
		}
	}
}

func h1(v int) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(v))
	cur := v
	for cur > 9 {
		cur = f[cur]
		sb.WriteString(strconv.Itoa(cur))
	}
	return sb.String()
}

func solve() {
	s := readLine()
	if len(s) == 1 {
		fmt.Fprintln(out, s)
		return
	}
	co := [10]int{}
	t := 0
	for i := 0; i < len(s); i++ {
		d := int(s[i] - '0')
		co[d]++
		t += d
	}
	for v := 1; v < t && v <= 1000000; v++ {
		if v+g[v] == t {
			s1 := h1(v)
			scn := [10]int{}
			p := true
			for i := 0; i < len(s1); i++ {
				d := int(s1[i] - '0')
				scn[d]++
				if scn[d] > co[d] {
					p = false
					break
				}
			}
			if p {
				rc := co
				for i := 0; i < 10; i++ {
					rc[i] -= scn[i]
				}
				var x strings.Builder
				for d := 1; d <= 9; d++ {
					if rc[d] > 0 {
						x.WriteByte(byte('0' + d))
						rc[d]--
						break
					}
				}
				for d := 0; d <= 9; d++ {
					for rc[d] > 0 {
						x.WriteByte(byte('0' + d))
						rc[d]--
					}
				}
				fmt.Fprint(out, x.String())
				fmt.Fprintln(out, s1)
				return
			}
		}
	}
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	precomp()
	t, _ := strconv.Atoi(readLine())
	for t > 0 {
		t--
		solve()
	}
}
