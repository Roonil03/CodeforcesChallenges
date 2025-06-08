// The Question is solved post the contest times

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	to, weight int
}

type PQItem struct {
	node int
	dist int64
}

type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*PQItem)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func dijkstraMaxEdge(graph [][]Edge, start int) []int64 {
	n := len(graph)
	dist := make([]int64, n)
	visited := make([]bool, n)
	for i := range dist {
		dist[i] = 1e18
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PQItem{start, 0})
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PQItem)
		node := current.node
		d := current.dist
		if visited[node] {
			continue
		}
		visited[node] = true
		dist[node] = d
		for _, edge := range graph[node] {
			if visited[edge.to] {
				continue
			}
			newDist := maxInt64(d, int64(edge.weight))
			heap.Push(pq, &PQItem{edge.to, newDist})
		}
	}
	return dist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < t; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(line[0])
		m, _ := strconv.Atoi(line[1])
		graph := make([][]Edge, n)
		for j := 0; j < m; j++ {
			scanner.Scan()
			line := strings.Fields(scanner.Text())
			u, _ := strconv.Atoi(line[0])
			v, _ := strconv.Atoi(line[1])
			w, _ := strconv.Atoi(line[2])
			u--
			v--
			graph[u] = append(graph[u], Edge{v, w})
			graph[v] = append(graph[v], Edge{u, w})
		}
		distFromStart := dijkstraMaxEdge(graph, 0)
		distFromEnd := dijkstraMaxEdge(graph, n-1)
		ans := int64(1e18)
		for u := 0; u < n; u++ {
			for _, edge := range graph[u] {
				v := edge.to
				w := int64(edge.weight)

				maxEdge := maxInt64(maxInt64(distFromStart[u], distFromEnd[v]), w)
				cost := maxEdge + w
				ans = minInt64(ans, cost)
			}
		}
		fmt.Println(ans)
	}
}

/*
This algorithm implements a heuristic bottleneck approach with the following steps:

1. Modified Dijkstra: Run a modified Dijkstra's algorithm from node 1 and node n, where the "distance" to
each node represents the minimum possible maximum edge weight on any path to that node

2. Edge Enumeration: For each edge (u,v) with weight w, consider it as a potential "bottleneck" edge in the
optimal path

3. Path Construction: For each such edge, construct a path: (1 → u) + (u → v) + (v → n) where:

	Maximum edge weight = max(distFromStart[u], distFromEnd[v], w)

	Assumed minimum edge weight = w

	Total cost = max(distFromStart[u], distFromEnd[v], w) + w

Time Complexity: O((V + E) log V) - Two Dijkstra runs plus O(E) edge enumeration
Space Complexity: O(V + E) - Graph storage plus distance arrays

This approach is much faster than the previous algorithm but trades optimality for efficiency. It works
well when the minimum edge weight on the optimal path coincides with one of the edges being considered,
but may produce suboptimal results in cases where the true minimum edge is on a subpath rather than the
connecting edge.
*/
