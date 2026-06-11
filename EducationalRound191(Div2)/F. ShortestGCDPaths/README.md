# F. Shortest GCD Paths
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
You are given three integers n
, a
, b
. Consider a weighted undirected graph with n
 vertices, where for every pair of distinct vertices `(u,v)`
 there is an edge with weight:

$$w(u, v) = \frac{\max(u, v)}{\gcd(u, v)}$$

Here $\gcd(x, y)$ denotes the greatest common divisor (GCD) of integers $x$ and $y$.

Find the shortest path from vertex $a$ to vertex $b$ in this graph.

### Input:
The only line contains three integers n
, a
, b
 (2 ≤ n ≤ 10<sup>9</sup>, 1 ≤ a, b ≤ n, a ≠ b
).

### Output
Print one integer — the length of the shortest path from vertex a
 to vertex b
.

### Examples:
#### Example 1:
**Input:**
```
10 9 8
```
**Output:**
```
7
```
**Explaination:**  
Consider the first example.

The shortest path in it is `9 → 6 → 8`
 with total cost `w(9,6) + w(6,8)` = `3 + 4` = 7
.
#### Example 2:
**Input:**
```
100 16 27
```
**Output:**
```
10
```
### Example 3:
**Input:**
```
100 55 11
```
**Output:**
```
5
```
