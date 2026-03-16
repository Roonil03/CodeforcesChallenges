package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readLine() string {
	line, err := in.ReadString('\n')
	if err != nil {
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

const mod = 998244353

func solve() {
	temp := readInts(2)
	n, m := temp[0], temp[1]
	a := readInt64s(n)
	k := readInt64s(m)
	l := make([]int, n)
	r := make([]int, n)
	st := []int{}
	for i := range n {
		for len(st) > 0 && a[st[len(st)-1]] >= a[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			l[i] = -1
		} else {
			l[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	st = st[:0]
	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && a[st[len(st)-1]] > a[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			r[i] = n
		} else {
			r[i] = st[len(st)-1]
		}
		st = append(st, i)
	}
	invmp := make(map[int64]int64)
	for _, v := range a {
		if _, ok := invmp[v]; !ok {
			invmp[v] = inv(v)
		}
	}
	p1 := int64(0)
	e := make([]Element, n)
	for i := 0; i < n; i++ {
		ca := (int64(i+1) * int64(n-i)) % mod
		ta := (ca * invmp[a[i]]) % mod
		p1 = (p1 + ta) % mod
		cb := (int64(i-l[i]) * int64(r[i]-i)) % mod
		e[i] = Element{
			val:   a[i],
			count: cb,
			inv:   invmp[a[i]],
		}
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].val < e[j].val
	})
	pre := make([]int64, n+1)
	preVal := make([]int64, n+1)
	preInv := make([]int64, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = (pre[i] + e[i].count) % mod
		preVal[i+1] = (preVal[i] + (e[i].val%mod)*e[i].count) % mod
		preInv[i+1] = (preInv[i] + e[i].inv*e[i].count) % mod
	}
	for _, v := range k {
		id := sort.Search(n, func(i int) bool {
			return e[i].val > v+1
		})
		sumcs2 := pre[id]
		sumvals2 := preVal[id]
		suminvs2 := preInv[id]
		suminvs1 := (preInv[n] - preInv[id] + mod) % mod
		kMod := v % mod
		kPlus2 := (v + 2) % mod
		ts1 := (kMod * suminvs1) % mod
		ts2 := (kPlus2 * sumcs2) % mod
		ts2 = (ts2 - sumvals2 + mod) % mod
		ts2 = (ts2 - suminvs2 + mod) % mod
		fmt.Fprintln(out, (p1+ts1+ts2)%mod)
	}
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t := 1
	for t > 0 {
		t--
		solve()
	}
}

func pow(a, b int64) int64 {
	res := int64(1)
	a %= mod
	for b > 0 {
		if b%2 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		b /= 2
	}
	return res
}

func inv(n int64) int64 {
	return pow(n, mod-2)
}

type Element struct {
	val   int64
	count int64
	inv   int64
}
