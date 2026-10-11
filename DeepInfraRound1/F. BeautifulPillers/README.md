# F. Beautiful Pillers
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
For an array $b$ of $m$ integers, a pair of indices $(l, r)$ ($l < r$) is called *good* if $b_l = b_r$, and there is no index $i$ ($l < i < r$) such that $b_i \ge b_l$.

We say that an index $i$ is *covered* by a good pair $(l, r)$ if $l \le i \le r$.

The array $b$ is called *beautiful* if every index is *covered* by at least one good pair.

You are given an array $a$ of $n$ integers. You can perform the following operation with it: choose two integers $i$ and $x$ ($1 \le i \le n$, $1 \le x \le 10^9$) and set $a_i = x$.

Find the minimum number of operations required to make the array *beautiful*. Also, output the operations. It can be shown that it is always possible to make the array *beautiful* using at most $n$ operations.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains a single integer $n$ ($2 \le n \le 2 \cdot 10^5$) — the length of the array $a$.

The second line of each test case contains $n$ integers $a_1, a_2, \dots, a_n$ ($1 \le a_i \le 10^9$) — the array $a$.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

### Output
For each test case, output a single integer $k$ ($0 \le k \le n$) — the minimum number of operations.

Then output $k$ lines describing the operations in the order they are applied. Each of these lines should contain two integers $i$ and $x$ ($1 \le i \le n, 1 \le x \le 10^9$) — the index and the new value of the corresponding operation.

If there are multiple optimal sequences of operations, output any of them.

### Examples:
**Input:**
```
5
3
1 2 3
8
5 3 2 5 1 7 9 9
5
1 1 1 1 1
8
1 2 3 4 4 3 2 1
2
1000000000 1000000000
```
**Output:**
```
1
1 3
1
5 7
0
2
1 3
8 3
0
```
**Explaination:**  
In the first example, the array $a = [1, 2, 3]$ has no good pairs at all, so it is not *beautiful*. After the operation with $i = 1$ and $x = 3$, the array becomes $[3, 2, 3]$. The pair $(1, 3)$ is good because its endpoints are both $3$ and its only internal value is $2 < 3$, so it *covers* all three indices.

In the second example, the array $a = [5, 3, 2, 5, 1, 7, 9, 9]$ has exactly two good pairs, $(1, 4)$ and $(7, 8)$, so the indices $5$ and $6$ are not *covered*. The operation with $i = 5$ and $x = 7$ turns the array into $[5, 3, 2, 5, 7, 7, 9, 9]$: the pairs $(1, 4)$ and $(7, 8)$ are still good, and the new good pair $(5, 6)$ *covers* the two remaining indices.

In the third example, the array is already *beautiful*, since $(i, i + 1)$ is a good pair for every $i$ from $1$ to $4$. Therefore, no operations are needed.