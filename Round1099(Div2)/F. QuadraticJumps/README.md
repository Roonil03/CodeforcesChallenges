# F. Quadratic Jumps
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given two integers n
 and q
. Consider a graph with n
 vertices, where vertices i
 and j
 are connected by an edge if and only if `|j − i|`
 is a perfect square<sup>∗</sup>
.

You are given q
 pairs of numbers a<sub>i</sub>, b<sub>i</sub>
. For each of these q
 pairs, you need to find the shortest distance between vertices a<sub>i</sub?>
 and b<sub>i</sub>
 in this graph. It can be proved that the graph is connected, so the distance between a<sub>i</sub>
 and b<sub>i</sub>
 is not infinite.

<sub><sup>∗</sup>
An integer x
 is a perfect square if there exists an integer y
 such that x = y<sup>2</sup>
.</sub>

### Input:
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 1000
). The description of the test cases follows.

The first line of each test case contains two integers n
 and q
 (2 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
, 1 ≤ q ≤ 10<sup>5</sup>
) — the number of vertices in the graph and the number of pairs of vertices for which the distance must be found.

Then the next q
 lines describe the pairs of vertices for which the shortest distance must be found. Each pair is described by two numbers a, b
 (1 ≤ a < b ≤ n
) — the numbers of the vertices between which the shortest distance must be found.

It is guaranteed that the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
, and the sum of q
 over all test cases does not exceed 10<sup>5</sup>
.

### Output
For each test case, output the shortest distance between each of the q
 pairs of vertices.

### Examples:
**Input:**
```
2
5 4
1 2
1 3
1 4
1 5
8 3
3 7
2 5
1 7
```
**Output:**
```
1
2
2
1
1
2
3
```
**Explaination:**  
This is what the graph looks like for the first test case:

<img src="https://espresso.codeforces.com/40bfd22066fe9acd7f50f81ceae09ea022b9a8cf.png"><br>

- For the first pair of vertices, the shortest path is 1 → 2
.
- For the second pair of vertices, the shortest path is 1 → 2 → 3
.
- For the third pair of vertices, the shortest path is 1 → 5 → 4
.
- For the fourth pair of vertices, the shortest path is 1 → 5
.

This is what the graph looks like for the second test case:

<img src="https://espresso.codeforces.com/d1730912dc669e44d9231a4d2fd63ef588cc7c11.png"><br>

- For the first pair of vertices, the shortest path is 1 → 2
.
- For the second pair of vertices, the shortest path is 2 → 6 → 5
.
- For the third pair of vertices, the shortest path is 1 → 2 → 6 → 7
.

