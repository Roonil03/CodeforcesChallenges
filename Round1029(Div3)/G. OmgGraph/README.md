# G. Omg Graph
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes

### Description:
You are given an undirected connected weighted graph. Define the cost of a path of length k
 to be as follows:

Let the weights of all the edges on the path be w1,...,wk
.
The cost of the path is (min<sup>k</sup><sub>i=1</sub>w<sub>i</sub>)+(max<sup>k</sup><sub>i=1</sub>w<sub>i</sub>)
, or in other words, the maximum edge weight + the minimum edge weight.
Across all paths from vertex 1
 to n
, report the cost of the path with minimum cost. Note that the path is not necessarily simple.

### Input:
The first line contains an integer t
 (1≤t≤10<sup>4</sup>
) — the number of test cases.

The first line of each test case contains two integers n
 and m
 (2≤n≤2⋅10<sup>5</sup>,n−1≤m≤min(2⋅10<sup>5</sup>,n(n−1)/2)
).

The next m
 lines each contain integers u,v
 and w
 (1≤u,v≤n,1≤w≤10<sup>9</sup>
) representing an edge from vertex u
 to v
 with weight w
. It is guaranteed that the graph does not contain self-loops or multiple edges and the resulting graph is connected.

It is guaranteed that the sum of n
 over all test cases does not exceed 2⋅10<sup>5</sup>
 and that the sum of m
 over all test cases does not exceed 2⋅10<sup>5</sup>
.

### Output:
For each test case, output a single integer, the minimum cost path from vertex 1
 to n
.

### Examples:
**Input:**
```
4
3 2
1 2 1
2 3 1
3 2
1 3 13
1 2 5
8 9
1 2 6
2 3 5
3 8 6
1 4 7
4 5 4
5 8 7
1 6 5
6 7 5
7 8 5
3 3
1 3 9
1 2 8
2 3 3
```
**Output:**
```
2
18
10
11
```
**Explaination:**  
For the second test case, the optimal path is 1→2→1→3
, the edge weights are 5,5,13
 so the cost is min(5,5,13)+max(5,5,13)=5+13=18
. It can be proven that there is no path with lower cost.