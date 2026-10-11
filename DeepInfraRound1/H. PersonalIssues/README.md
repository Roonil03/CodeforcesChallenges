# H1. Personal Issues (Easy Version)
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
**This is the easy version of the problem. The difference between the versions is that in this version, $t \le 1000$, $n \le 3000$ and the sum of $n$ over all test cases does not exceed $3000$. You can hack only if you solved all versions of this problem.**

The rivals Axial and Nife have had yet another argument, and they decided to settle it by playing a game on a tree.

They are given a tree with $n$ vertices numbered from $1$ to $n$. Each vertex contains either $0$ or $1$. Initially, two distinct vertices $x$ and $y$ contain $0$, and all other vertices contain $1$.

The players take turns, with Axial going first. On each turn, the current player chooses a vertex $v$ containing $1$ and changes its value to $0$. Before changing the value of $v$, let $c_v$ be the number of neighbors of $v$ that contain $0$. If $c_v = 1$, the current player gets one point. Otherwise, they get no points. The game ends when all vertices contain $0$.

Let $p$ and $q$ be the final scores of Axial and Nife, respectively. Axial wants to maximize $p - q$, while Nife wants to minimize it. Assuming both players play optimally, determine the value of $p - q$ at the end of the game.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 1000$). The description of the test cases follows.

The first line of each test case contains three integers $n$, $x$, and $y$ ($2 \le n \le 3000$, $1 \le x, y \le n$, $x \ne y$) — the number of vertices in the tree and the two vertices that initially contain $0$.

Each of the next $n - 1$ lines contains two integers $u$ and $v$ ($1 \le u, v \le n$, $u \ne v$), denoting an edge between vertices $u$ and $v$.

It is guaranteed that the given graph is a tree.

It is guaranteed that the sum of $n$ over all test cases does not exceed $3000$.

### Output
For each test case, output a single integer — the value of $p - q$ at the end of the game.

### Examples:
**Input:**
```
4
3 1 2
1 2
2 3
4 1 4
1 2
2 3
3 4
3 1 3
1 2
2 3
6 2 3
1 2
1 3
1 4
1 5
1 6
```
**Output:**
```
1
1
0
0
```
**Explaination:**  
In the figures below, blue vertices initially contain 0
, while white vertices initially contain 1
.

For the first example, the initial state is:

<img src="https://espresso.codeforces.com/e9cd98c760fe6648d25e0a629a0cd74beffac9d8.png"><br>

Axial changes vertex 3
 to 0
. It has exactly one neighbor containing 0
, so Axial gets one point. Therefore, p−q=1
.

For the second example, the initial state is:

<img src="https://espresso.codeforces.com/8509b22a1a90150db9d0ced2994f4b274ae098eb.png"><br>

Axial can change either vertex 2
 or vertex 3
 to 0
 and get one point. After that, the remaining vertex has two neighbors containing 0
, so Nife gets no points. Therefore, p−q=1
.

For the third example, the initial state is:

<img src="https://espresso.codeforces.com/7ef05afa3fec9ce0ae24a054df3667379ed20537.png"><br>

The only possible move is to change vertex 2
 to 0
. It has two neighbors containing 0
, so Axial gets no points. Therefore, p−q=0
.

# H2. Personal Issues (Hard Version)
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
**This is the hard version of the problem. The difference between the versions is that in this version, $t \le 10^4$, $n \le 2 \cdot 10^5$ and the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$. You can hack only if you solved all versions of this problem.**

The rivals Axial and Nife have had yet another argument, and they decided to settle it by playing a game on a tree.

They are given a tree with $n$ vertices numbered from $1$ to $n$. Each vertex contains either $0$ or $1$. Initially, two distinct vertices $x$ and $y$ contain $0$, and all other vertices contain $1$.

The players take turns, with Axial going first. On each turn, the current player chooses a vertex $v$ containing $1$ and changes its value to $0$. Before changing the value of $v$, let $c_v$ be the number of neighbors of $v$ that contain $0$. If $c_v = 1$, the current player gets one point. Otherwise, they get no points. The game ends when all vertices contain $0$.

Let $p$ and $q$ be the final scores of Axial and Nife, respectively. Axial wants to maximize $p - q$, while Nife wants to minimize it. Assuming both players play optimally, determine the value of $p - q$ at the end of the game.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains three integers $n$, $x$, and $y$ ($2 \le n \le 2 \cdot 10^5$, $1 \le x, y \le n$, $x \ne y$) — the number of vertices in the tree and the two vertices that initially contain $0$.

Each of the next $n - 1$ lines contains two integers $u$ and $v$ ($1 \le u, v \le n$, $u \ne v$), denoting an edge between vertices $u$ and $v$.

It is guaranteed that the given graph is a tree.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

### Output
For each test case, output a single integer — the value of $p - q$ at the end of the game.

### Examples:
**Input:**
```
4
3 1 2
1 2
2 3
4 1 4
1 2
2 3
3 4
3 1 3
1 2
2 3
6 2 3
1 2
1 3
1 4
1 5
1 6
```
**Output:**
```
1
1
0
0
```
**Explaination:**  
In the figures below, blue vertices initially contain 0
, while white vertices initially contain 1
.

For the first example, the initial state is:

<img src="https://espresso.codeforces.com/e9cd98c760fe6648d25e0a629a0cd74beffac9d8.png"><br>

Axial changes vertex 3
 to 0
. It has exactly one neighbor containing 0
, so Axial gets one point. Therefore, p−q=1
.

For the second example, the initial state is:

<img src="https://espresso.codeforces.com/8509b22a1a90150db9d0ced2994f4b274ae098eb.png"><br>

Axial can change either vertex 2
 or vertex 3
 to 0
 and get one point. After that, the remaining vertex has two neighbors containing 0
, so Nife gets no points. Therefore, p−q=1
.

For the third example, the initial state is:

<img src="https://espresso.codeforces.com/7ef05afa3fec9ce0ae24a054df3667379ed20537.png"><br>

The only possible move is to change vertex 2
 to 0
. It has two neighbors containing 0
, so Axial gets no points. Therefore, p−q=0
.
