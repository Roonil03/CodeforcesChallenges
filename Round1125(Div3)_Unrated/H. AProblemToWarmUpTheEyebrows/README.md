# H. A Problem to Warm Up the Eyebrows
## Time limit per test: 1.5 seconds
## Memory limit per test: 256 megabytes
### Description:
K1o0n came up with this amusing game for his friend Noobish_Monk. They take a rectangular table of integers, and for every solid rectangular piece of this table Noobish_Monk either gets something or, on the contrary, loses something — depending on the shape of the piece. Figuring out how much he gets in total became Noobish_Monk's task, but once again he misread the statement $998244353$ times, which caused him to get stuck on an insignificant piece.

We invite you to try your hand at this game as well!

You are given a table of integers $A$ with $n$ rows and $m$ columns. A submatrix is any solid rectangular piece of the table; it is determined by the indices of its first and last rows and the indices of its first and last columns. Suppose a submatrix occupies $a$ consecutive rows and $b$ consecutive columns, and the sum of all numbers in it is $S$. If $a + b$ is even, then Noobish_Monk gets $S^2$, and if it is odd, he loses $S^2$. This quantity is called the price of the submatrix:

$$(-1)^{a+b} S^2$$

The price of the table is the sum of the prices of all its submatrices, including the whole table. Help him: find the price of table $A$. It can be negative and can also be very large in magnitude, so output it **modulo $10^9 + 7$**.

### Input
The first line contains an integer $t$ ($1 \le t \le 10^4$) — the number of testcases.

The first line of each testcase contains two integers $n$ and $m$ ($1 \le n, m, 1 \le n \cdot m \le 10^6$) — the dimensions of the matrix.

The next $n$ lines contain $m$ integers each: the $j$-th number in the $i$-th line equals $A_{i, j}$ ($-10^9 \le A_{i, j} \le 10^9$).

It is guaranteed that the sum of $n \cdot m$ over all testcases does not exceed $10^6$.

### Output
For each testcase, output a single integer in a separate line from the range $[0, 10^9 + 6]$ — the remainder of the table price divided by $10^9 + 7$.

### Examples:
**Input:**
```
4
3 3
1 5 2
7 9 8
3 6 4
2 4
1 -1 -2 2
-2 2 1 -1
3 2
4 5
8 2
0 0
4 4
1 1 1 1
-1 -1 -1 -1
1 1 1 1
-1 -1 -1 -1
```
**Output:**
```
100
6
999999967
999999971
```
