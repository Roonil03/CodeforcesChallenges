# E. Product of Closures
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
Let the *closure* of a positive integer $x > 0$ be the following infinite binary string $C(x)$:

1. Write $x$ in binary form without leading zeros;
2. Concatenate the obtained string with itself infinitely many times and call the result $C(x)$.

For example, $C(1) = 11111\dots$, $C(4) = 10010010010010\dots$, $C(9) = 100110011001\dots$.

Let the *product* of two closures $C(x)$ & $C(y)$ be the infinite binary string obtained by the **bitwise binary AND** of the strings $C(x)$ and $C(y)$. For example, for $C(4)$ & $C(9)$ we get

$$
\begin{array}{r@{\quad}l}
C(4) & 100100100100100100\dots \\
\text{\& } C(9) & 100110011001100110\dots \\
\hline
& 100100000000100100\dots
\end{array}
$$

You are given three integers $l$, $r$, and $n$. Among all numbers in the segment $[l, r]$, find two integers $x$ and $y$ ($l \le x < y \le r$) such that $C(x) \& C(y)$ is **lexicographically smallest**, and print the first $n$ binary symbols of this product.

### Input
The first line contains one integer $t$ ($1 \le t \le 1000$) — the number of test cases.

The only line of each test case contains three integers $l$, $r$, and $n$ ($1 \le l < r < 2^{30}; 1 \le n \le 1000$) — the range of possible values and the length of the answer.

### Output
For each test case, print one binary string of length $n$ — the first $n$ symbols of the lexicographically smallest product of closures.

### Examples:
**Input:**
```
3
1 4 10
1073741822 1073741823 35
10 20 15
```
**Output:**
```
1000001000
11111111111111111111111111111011111
100000000010000
```
**Explaination:**  
In the first test case, the lexicographically smallest string is obtained by the product `C(2)` & `C(4)`.