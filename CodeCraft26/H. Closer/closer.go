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

type Pair struct {
	label int
	val   int64
}

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
	for i := 0; i < n; i++ {
		res[i], _ = strconv.Atoi(parts[i])
	}
	return res
}

func readInt64s(n int) []int64 {
	parts := strings.Fields(readLine())
	res := make([]int64, n)
	for i := 0; i < n; i++ {
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

func lookup(pairs []Pair, label int) (int64, bool) {
	i := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].label >= label
	})
	if i < len(pairs) && pairs[i].label == label {
		return pairs[i].val, true
	}
	return 0, false
}

func totalForLabel(base int64, pairs []Pair, label int) int64 {
	if v, ok := lookup(pairs, label); ok {
		return base + v
	}
	return base
}

func solve() {
	n, _ := strconv.Atoi(readLine())
	N := 2 * n
	wInput := readInt64s(n)
	w := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		w[i] = wInput[i-1]
	}
	aInput := readInts(N)
	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = aInput[i-1]
	}
	adj := make([][]int, N+1)
	for i := 0; i < N-1; i++ {
		uv := readInts(2)
		u, v := uv[0], uv[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	parent := make([]int, N+1)
	order := make([]int, 0, N)
	stack := []int{1}
	parent[1] = -1
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, u)
		for _, v := range adj[u] {
			if v == parent[u] {
				continue
			}
			parent[v] = u
			stack = append(stack, v)
		}
	}
	baseChildren := make([]int64, N+1)
	baseNoParent := make([]int64, N+1)
	bonusPairs := make([][]Pair, N+1)
	bestPairs := make([][]Pair, N+1)
	noChoose := func(v int, parentLabel int) int64 {
		best := baseNoParent[v]
		if x, ok := lookup(bestPairs[v], parentLabel); ok {
			cand := x + w[parentLabel]
			if cand > best {
				best = cand
			}
		}
		return best
	}
	chooseByParent := func(v int, parentOrigLabel int) int64 {
		res := totalForLabel(baseChildren[v], bonusPairs[v], parentOrigLabel)
		if a[v] == parentOrigLabel {
			res += w[parentOrigLabel]
		}
		return res
	}
	for idx := len(order) - 1; idx >= 0; idx-- {
		u := order[idx]
		base := int64(0)
		tmpBonus := make(map[int]int64)
		for _, v := range adj[u] {
			if v == parent[u] {
				continue
			}
			base += baseNoParent[v]
			for _, pr := range bestPairs[v] {
				delta := pr.val + w[pr.label] - baseNoParent[v]
				if delta > 0 {
					tmpBonus[pr.label] += delta
				}
			}
		}
		baseChildren[u] = base
		if len(tmpBonus) > 0 {
			bp := make([]Pair, 0, len(tmpBonus))
			for label, val := range tmpBonus {
				bp = append(bp, Pair{label: label, val: val})
			}
			sort.Slice(bp, func(i, j int) bool {
				return bp[i].label < bp[j].label
			})
			bonusPairs[u] = bp
		}
		noneVal := totalForLabel(baseChildren[u], bonusPairs[u], a[u])
		best := noneVal
		bestMap := make(map[int]int64)
		bestMap[a[u]] = noneVal
		for _, v := range adj[u] {
			if v == parent[u] {
				continue
			}
			labelAfter := a[v]
			totalIfLabel := totalForLabel(baseChildren[u], bonusPairs[u], labelAfter)
			val := totalIfLabel - noChoose(v, labelAfter) + chooseByParent(v, a[u])
			if val > best {
				best = val
			}
			if cur, ok := bestMap[labelAfter]; !ok || val > cur {
				bestMap[labelAfter] = val
			}
		}
		baseNoParent[u] = best
		bp := make([]Pair, 0, len(bestMap))
		for label, val := range bestMap {
			bp = append(bp, Pair{label: label, val: val})
		}
		sort.Slice(bp, func(i, j int) bool {
			return bp[i].label < bp[j].label
		})
		bestPairs[u] = bp
	}
	fmt.Fprintln(out, baseNoParent[1])
}

func main() {
	out = bufio.NewWriterSize(os.Stdout, 1<<20)
	defer out.Flush()
	t, _ := strconv.Atoi(readLine())
	for t > 0 {
		t--
		solve()
	}
}

/*
The key idea is to stop thinking in terms of “which duplicate pair becomes adjacent” and instead
think in terms of the final labels sitting on each tree edge after one matching-swap operation:
a deal contributes exactly when some edge ends up with equal labels on both endpoints. Root the
tree, let every vertex’s state be “stay” or “swap with exactly one neighbor”, and do DP bottom-up.
For a node u, the parent only needs two summaries: the best value in u’s subtree if u does not
swap with its parent, and for each possible final label that can appear on u (only a[u] or one
of its children’s labels), the best value under that exact final label. Then the edge to the
parent is easy: if the parent does not swap with u, you only gain the edge profit when u’s final
label matches the parent’s final label; if the parent does swap with u, then u is forced to take
the parent’s original label and all of u’s children become independent under that fixed label.
The “bonus by label” maps stay sparse because each node can only end with labels coming from
itself or its children, so the total stored information is linear, and with sorting/binary-search
on those sparse lists the full solution runs in O(n log n) over each test batch.
*/
