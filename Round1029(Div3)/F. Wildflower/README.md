# F. Wildflower
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Yousef has a rooted tree∗
 consisting of exactly n
 vertices, which is rooted at vertex 1
. You would like to give Yousef an array a
 of length n
, where each ai
 (1≤i≤n)
 can either be 1
 or 2
.

Let s<sub>u</sub>
 denote the sum of a<sub>v</sub>
 where vertex v
 is in the subtree†
 of vertex u
. Yousef considers the tree special if all the values in s
 are pairwise distinct (i.e., all subtree sums are unique).

Your task is to help Yousef count the number of different arrays a
 that result in the tree being special. Two arrays b
 and c
 are different if there exists an index i
 such that b<sub>i</sub>≠c<sub>i</sub>
.

As the result can be very large, you should print it modulo 109+7
.

<sub>∗
 A tree is a connected undirected graph with n−1
 edges.

†
The subtree of a vertex v
 is the set of all vertices that pass through v
 on a simple path to the root. Note that vertex v
 is also included in the set.</sub>

 ### Input:
 The first line contains an integer t
 (1≤t≤10<sup>4</sup>)
 — the number of test cases.

Each test case consists of several lines. The first line of the test case contains an integer n
 (2≤n≤2⋅10<sup>5</sup>)
 — the number of vertices in the tree.

Then n−1
 lines follow, each of them contains two integers u
 and v
 (1≤u,v≤n,u≠v)
 which describe a pair of vertices connected by an edge. It is guaranteed that the given graph is a tree and has no loops or multiple edges.

It is guaranteed that the sum of n
 over all test cases doesn't exceed 2⋅10<sup>5</sup>
.

### Output
For each test case, print one integer x
 — the number of different arrays a
 that result in the tree being special, modulo 10<sup>9</sup>+7
.

### Examples:
**Input:**
```
7
2
1 2
8
1 2
2 3
3 8
2 4
4 5
5 6
6 7
10
1 2
2 3
3 4
4 5
5 6
4 7
7 8
4 9
9 10
7
1 4
4 2
3 2
3 5
2 6
6 7
7
1 2
2 3
3 4
3 5
4 6
6 7
7
5 7
4 6
1 6
1 3
2 6
6 7
5
3 4
1 2
1 3
2 5
```
**Output:**
```
4
24
0
16
48
0
4
```
**Explaination:**  
The tree given in the fifth test case:  
<img src="https://espresso.codeforces.com/bd9f04ef8898b0b29a6e5e86a17295ed03a19d03.png"><br>