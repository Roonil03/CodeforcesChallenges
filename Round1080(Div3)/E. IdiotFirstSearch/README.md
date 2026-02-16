# E. Idiot First Search
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
There is a binary tree of n+1 vertices (n is odd), with vertices labeled 0,1,…,n. At most one letter can be written on each vertex of the tree, and all vertices initially have nothing written on them. The root of the tree is vertex 0.

In the tree, vertex 0 is the parent of vertex 1, while all other vertices have either 2 children or 0 children.

Bob is lost in one vertex of the tree and wishes to escape the tree by reaching vertex 0. This is very easy for most people with common sense. However, since Bob is an idiot, he created a new way of traversing the tree; introducing the "Idiot First Search".

When Bob is on vertex v (1≤v≤n), Bob's movement is determined as follows:
- If vertex v is a leaf, Bob always moves to the parent of v; otherwise, check the next few conditions.
- If nothing is written on vertex v, Bob writes 'L' on vertex v and moves to the left child of v;
- If '`L`' is written on vertex v, Bob overwrites it to 'R' and moves to the right child of v;
- If '`R`' is written on vertex v, Bob erases it and moves to the parent of v. 

It takes exactly 1 second for Bob to move to an adjacent vertex, so Bob will take exactly x seconds to perform x moves.

It has been shown that regardless of which vertex Bob starts on, Bob can reach vertex 0 in a finite (though possibly inexplicably large) amount of time. We don't know who proved it; surely it can't be Bob, but it is definitely proven.

For each vertex k=1,2,…,n, please determine the total time it takes to reach vertex 0 if Bob started on vertex k, in seconds. As the values may be huge, you are only asked to compute them modulo 10<sup>9</sup>+7.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains a single integer n (1 ≤ n ≤ 300001, n is odd).

Each of the next n lines contains two integers li and ri denoting the children of vertex i (0≤ l<sub>i</sub>, r<sub>i</sub> ≤ n).

For each vertex, l<sub>i</sub>=r<sub>i</sub>=0 is given if the vertex has no children. Otherwise, l<sub>i</sub> and r<sub>i</sub> are the left and right children of vertex i.

It is guaranteed that the input defines a valid binary tree satisfying the conditions given in the statement.

It is guaranteed that the sum of n over all test cases does not exceed 300001.
### Output

For each test case, output n integers τ<sub>1</sub>,τ<sub>2</sub>,…,τ<sub>n</sub> separated by spaces.

Here, τ<sub>k</sub> denotes the total time it takes to reach vertex 0 if Bob started on vertex k, modulo 10<sup>9</sup>+7.

### Examples:
**Input:**
```
3
1
0 0
5
2 3
0 0
4 5
0 0
0 0
7
2 3
4 5
0 0
6 7
0 0
0 0
0 0
```
**Output:**
```
1
9 10 14 15 15
13 22 14 27 23 28 28
```

**Note:**
On the first test case, there are only two vertices, vertex 0 and vertex 1. It takes only 1 second for Bob to reach vertex 0 from vertex 1.

On the second test case, the tree is given as follows.

<img src="https://espresso.codeforces.com/dfccc9cd793030983f22e97a4aefba2172a64dd9.png"><br>

It takes 14 seconds for Bob to reach vertex 0 from vertex 3. The moves are as follows:

$$3 \xrightarrow{L} 4 \xrightarrow{X} 3 \xrightarrow{R} 5 \xrightarrow{X} 3 
\xrightarrow{X} 1 \xrightarrow{L} 2 \xrightarrow{X} 1 \xrightarrow{R} 3 
\xrightarrow{L} 4 \xrightarrow{X} 3 \xrightarrow{R} 5 \xrightarrow{X} 3 
\xrightarrow{X} 1 \xrightarrow{X} 0
$$

Here, the letters above the arrows denote the letter on the vertex before moving to the adjacent vertex, where X denotes nothing written.