# G. Grid Filling
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given a grid of size $n \times m$. The rows are numbered from $1$ to $n$ from top to bottom, and the columns are numbered from $1$ to $m$ from left to right. Two cells are *adjacent* if they share a side.

There are $h$ horizontal and $v$ vertical cuts in the grid. The $i$-th horizontal cut goes between rows $r_i$ and $r_i + 1$, and the $j$-th vertical cut goes between columns $c_j$ and $c_j + 1$; every cut crosses the whole grid. Together with the borders of the grid, the cuts split it into $(h + 1) \cdot (v + 1)$ rectangular *pieces*.

Write a positive integer in every cell of the grid so that both of the following conditions hold:

* any two adjacent cells contain different integers, **even if** these cells belong to different pieces;
* every piece consisting of $x$ cells contains each of the integers $1, 2, \dots, x$ exactly once.

Find any way to fill the grid satisfying these conditions, or determine that it is impossible.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 1000$). The description of the test cases follows.

The first line of each test case contains four integers $n$, $m$, $h$, and $v$ ($1 \le n, m \le 1000$, $0 \le h < n$, $0 \le v < m$) — the number of rows of the grid, the number of columns of the grid, the number of horizontal cuts, and the number of vertical cuts.

The next line contains $h$ integers $r_1, r_2, \dots, r_h$ ($1 \le r_1 < r_2 < \dots < r_h < n$) — the positions of the horizontal cuts. If $h = 0$, this line is not present.

The next line contains $v$ integers $c_1, c_2, \dots, c_v$ ($1 \le c_1 < c_2 < \dots < c_v < m$) — the positions of the vertical cuts. If $v = 0$, this line is not present.

It is guaranteed that the sum of $n$ over all test cases does not exceed $1000$.

It is guaranteed that the sum of $m$ over all test cases does not exceed $1000$.

### Output
For each test case, if there is no valid way to fill the grid, output $-1$.

Otherwise, output $n$ lines. The $i$-th of these lines should contain $m$ integers — the numbers written in the cells of the $i$-th row of the grid, from left to right.

If there are multiple valid ways to fill the grid, output any of them.

### Examples:
**Input:**
```
4
3 3 0 0
1 6 0 2
1 3
5 5 2 2
2 3
2 3
4 4 2 2
1 2
1 2
```
**Output:**
```
1 2 3
4 5 6
7 8 9
1 2 1 2 1 3
1 2 1 2 1
3 4 2 3 4
1 2 1 2 1
3 4 2 3 4
1 2 1 2 1
-1
```
**Explaination:**  
In the first example, one possible way to fill the grid is shown below:  
<img src="https://espresso.codeforces.com/eda15f0cc6ccf867e34bf8e9a617023beda6bea5.png"><br>
In the third example, the dashed lines represent the cuts. One possible way to fill the grid is shown below:   
<img src="https://espresso.codeforces.com/b5c8bbd174903f93312dd6854945c5de3930c5b8.png"><br>
In the fourth example, it can be shown that no valid filling exists.