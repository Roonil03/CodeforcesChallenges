# D. Target Infection
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given $n$ pairwise disjoint infection intervals on the number line. The $i$-th interval is $[l_i, r_i]$, meaning that every integer position $x$ such that $l_i \le x \le r_i$ is infected. All integer positions not covered by any of the intervals are not infected.

There are $m$ people standing in a row. If the first person stands at position $s$, then the people stand at positions $s, s + 1, \dots, s + m - 1$. A person is infected if the position they stand at is infected.

For an integer $s$, let $f(s)$ be the number of infected people when the first person stands at position $s$.

You are given an integer $k$. Find any integer $s$ such that $f(s) = k$, or determine that no such $s$ exists.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains three integers $n$, $m$, and $k$ ($1 \le n \le 2 \cdot 10^5$, $1 \le m \le 10^9$, $0 \le k \le m$) — the number of infection intervals, the number of people, and the required number of infected people.

The $i$-th of the following $n$ lines contains two integers $l_i$ and $r_i$ ($1 \le l_i \le r_i \le 10^{15}$) — the endpoints of the $i$-th infection interval. It is guaranteed that $r_i + 1 < l_{i+1}$ for every $i$ from $1$ to $n - 1$, so the intervals are pairwise disjoint, no two of them are adjacent, and they are given from left to right.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

### Output
For each test case, if there is no integer $s$ such that $f(s) = k$, output $-1$.

Otherwise, output a single integer $s$ ($1 \le s \le 10^{18}$) — a position such that $f(s) = k$. It can be shown that if at least one such position exists, then a position within these bounds exists as well.

If there are multiple valid positions, output any of them.

### Examples:
**Input:**
```
6
1 5 5
7 11
2 4 3
1 2
5 6
3 5 4
1 3
5 6
8 8
1 3 2
10 12
3 6 0
1 1
4 5
10 12
1 1000000000 1
1000000000000000 1000000000000000
```
**Output:**
```
7
-1
2
9
42
999999876543210
```
**Explaination:**  
In the first example, one valid position is $s = 7$. The people then stand at the positions $7, 8, 9, 10, 11$, all of which are infected, so $f(7) = 5$.

In the second example, it can be shown that there is no suitable position.

In the third example, one valid position is $s = 2$. The people then stand at the positions $2, 3, 4, 5, 6$. Positions $2, 3, 5$, and $6$ are infected, while position $4$ is not, so $f(2) = 4$.