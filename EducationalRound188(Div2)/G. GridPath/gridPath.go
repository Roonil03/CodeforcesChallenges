// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	in  = bufio.NewReader(os.Stdin)
// 	out = bufio.NewWriter(os.Stdout)
// )

// func readLine() string {
// 	line, err := in.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}
// 	return strings.TrimSpace(line)
// }

// func readInts(n int) []int {
// 	parts := strings.Fields(readLine())
// 	res := make([]int, n)
// 	for i := range n {
// 		res[i], _ = strconv.Atoi(parts[i])
// 	}
// 	return res
// }

// func readInt64s(n int) []int64 {
// 	parts := strings.Fields(readLine())
// 	res := make([]int64, n)
// 	for i := range n {
// 		res[i], _ = strconv.ParseInt(parts[i], 10, 64)
// 	}
// 	return res
// }

// func printSlice(arr []int) {
// 	for i, v := range arr {
// 		if i > 0 {
// 			fmt.Fprint(out, " ")
// 		}
// 		fmt.Fprint(out, v)
// 	}
// 	fmt.Fprintln(out)
// }

// type Matrix struct {
// 	f [305][305]int64
// }

// func (a *Matrix) Mul(b *Matrix, MOD int64, size int) *Matrix {
// 	ans := &Matrix{}
// 	for i := 0; i <= size; i++ {
// 		for j := 0; j <= size; j++ {
// 			if a.f[i][j] > 0 {
// 				val := a.f[i][j]
// 				for k := 0; k <= size; k++ {
// 					if b.f[j][k] > 0 {
// 						ans.f[i][k] = (ans.f[i][k] + val*b.f[j][k]) % MOD
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return ans
// }

// func solve() {
// 	line := readLine()
// 	parts := strings.Fields(line)
// 	rows, _ := strconv.Atoi(parts[0])
// 	cols, _ := strconv.Atoi(parts[1])
// 	MOD, _ := strconv.ParseInt(parts[2], 10, 64)
// 	mSize := cols + cols + 1
// 	v := make([]int64, mSize+1)
// 	for i := 0; i <= cols+1; i++ {
// 		v[i] = 1
// 	}
// 	v[mSize] = MOD - 1
// 	f0 := &Matrix{}
// 	f0.f[0][0] = int64(cols*(cols+1)/2) % MOD
// 	f0.f[0][mSize] = 1
// 	f0.f[mSize][mSize] = 1
// 	su := func(a *int64, b int64) {
// 		*a += b
// 		if *a >= MOD {
// 			*a -= MOD
// 		}
// 	}
// 	buildState := func(l, r int) {
// 		for i := r; i <= cols; i++ {
// 			su(&f0.f[0][i], 1)
// 		}
// 		for i := l; i >= 1; i-- {
// 			su(&f0.f[0][cols+i], 1)
// 		}
// 		var vec []int
// 		if l > 1 {
// 			vec = append(vec, l-1)
// 		}
// 		if r < cols {
// 			vec = append(vec, cols+r+1)
// 		}
// 		for _, x := range vec {
// 			su(&f0.f[x][0], MOD-1)
// 			for i := r; i <= cols; i++ {
// 				su(&f0.f[x][i], MOD-1)
// 			}
// 			for i := l; i >= 1; i-- {
// 				su(&f0.f[x][cols+i], MOD-1)
// 			}
// 		}
// 	}
// 	for i := 1; i <= cols; i++ {
// 		for j := i; j <= cols; j++ {
// 			buildState(i, j)
// 		}
// 	}
// 	tExp := rows
// 	for tExp > 0 {
// 		if tExp&1 != 0 {
// 			v2 := make([]int64, mSize+1)
// 			for i := 0; i <= mSize; i++ {
// 				if v[i] == 0 {
// 					continue
// 				}
// 				for j := 0; j <= mSize; j++ {
// 					if f0.f[i][j] > 0 {
// 						v2[j] = (v2[j] + v[i]*f0.f[i][j]) % MOD
// 					}
// 				}
// 			}
// 			for i := 0; i <= mSize; i++ {
// 				v[i] = v2[i]
// 			}
// 		}
// 		if tExp == 1 {
// 			break
// 		}
// 		f0 = f0.Mul(f0, MOD, mSize)
// 		tExp >>= 1
// 	}
// 	v[mSize] = (v[mSize] + v[0]) % MOD
// 	fmt.Fprintln(out, v[mSize])
// }

// func main() {
// 	out = bufio.NewWriterSize(os.Stdout, 1<<20)
// 	defer out.Flush()
// 	t := 1
// 	for t > 0 {
// 		t--
// 		solve()
// 	}
// }
