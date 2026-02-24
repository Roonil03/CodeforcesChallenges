# D. Cost of Tree
## Time limit per test: 3 seconds
## Memory limit per test: 256 megabytes
### Description:
For a tree T with root r, where each node u has a value au associated with it, the cost of the tree defined as:
$$\sum_{u \in T} (a_u \cdot d(r,u))$$

Here, this sum is taken over all nodes u in the tree T, and d(r,u) denotes the number of edges on the shortest path from node r to node u on a tree.

You are given a tree consisting of n nodes, rooted at node 1. Each node i has a value ai assigned to it. For each r from 1 to n, please solve the following problem independently:

Consider the subtree of node r with respect to node 1. Formally, the subtree of node r is the tree consisting of all nodes u such that the shortest path from 1 to u contains r.

Find the maximum cost of the subtree after performing at most one operation of the following type on the subtree:

- Choose any node u (u ≠ r). Remove the edge from u to the parent of node u<sup>∗</sup>. Then, add an edge from u to any node v that is still reachable from r. It can be shown that after this operation, the graph remains a tree. 

As an example, below shows an example of an operation with r = 1,u = 5, and v = 4. 

<img src="https://espresso.codeforces.com/08cbc8b0340fa37c3036d5842d029b6d770d5775.png"><br>

<sub><sup>∗</sup>Formally, remove the edge from u to p, where p is the unique node satisfying d(u,p)=1 and d(u,r)=d(p,r)+1</sub>

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each testcase contains a single integer n (1 ≤ n ≤ 2⋅10<sup>5</sup>) — the count of nodes in the tree.

The second line of each testcase contains n integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ 2⋅10<sup>5</sup>).

Then n−1 lines follow, the i-th line containing two integers u and v (1 ≤ u, v ≤ n) — the two nodes that the i-th edge connects.

It is guaranteed that the given edges form a tree.

It is guaranteed that the sum of n does not exceed 2⋅10<sup>5</sup> over all test cases.
### Output

For each test case, print n numbers – the answers for r = 1, 2, …, n.

### Examples:
**Input:**
```
3
5
1 3 2 1 2
1 2
2 3
3 4
3 5
7
1 2 3 1 3 2 1
1 2
2 3
3 4
4 5
4 6
3 7
5
5 4 3 2 1
1 2
2 3
3 4
4 5
```
**Output:**
```
18 10 5 0 0 
40 28 18 8 0 0 0 
20 10 4 1 0 
```
**Note:**  
In the first test case, for r = 1, it is optimal to choose u = 5 and v = 4. The cost of the tree is then `1⋅0+3⋅1+2⋅2+1⋅3+2⋅4 = 18`. It can be shown that a larger cost cannot be obtained over all legal operations.

For r = 4 for example, there is only 1 node in the subtree, so there is no operation possible. The only possible cost of the subtree is 0.