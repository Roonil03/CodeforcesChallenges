# C. XOR Problem
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given an integer $n$.

Define the *beauty* of an array $a$ as the value of $|a| - f(a)$, where $|a|$ is the length of $a$ and $f(a)$ is the number of non-empty subarrays<sup>*</sup> of $a$ with bitwise XOR equal to $0$. Formally, $f(a)$ is the number of pairs $(l, r)$ such that $1 \le l \le r \le |a|$ and $a_l \oplus a_{l+1} \oplus \dots \oplus a_r = 0$ $\dagger$.

Among all arrays $a$ consisting of integers from $0$ to $n$ that have the maximum possible beauty, find one with maximum length.

<sub><sup>*</sup>An array $c$ is a subarray of an array $d$ if $c$ can be obtained from $d$ by the deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end. Two subarrays are considered different if the sets of positions of the deleted elements are different.</sub>

$\dagger \oplus$ denotes the bitwise XOR operation.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The only line of each test case contains a single integer $n$ ($1 \le n \le 2 \cdot 10^5$) — the upper bound on the elements of the array.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

### Output
For each test case, output a single integer $k$ — the length of the array. On the next line, output $k$ integers $a_1, a_2, \dots, a_k$ ($0 \le a_i \le n$) — the elements of the array.

It can be shown that under the given constraints the sum of $k$ over all test cases does not exceed $10^6$.

If there are multiple suitable arrays, output any of them.

### Examples:
**Input:**
```
2
1
2
```
**Output:**
```
3
1 1 1
7
2 2 1 0 2 0 1
```
**Explaination:**  
In the first example, $n = 1$, so every element of the array is equal to $0$ or $1$. One possible solution is $a = [1, 1, 1]$: it has exactly two subarrays with bitwise XOR equal to $0$, namely $(l, r) = (1, 2)$ and $(l, r) = (2, 3)$. These two subarrays consist of the same elements, but they occupy different positions, so both of them are counted, and the beauty of $a$ is $3 - 2 = 1$. It can be proven that every array of length $4$ with elements equal to $0$ or $1$ has at least four subarrays with bitwise XOR equal to $0$, so its beauty is at most $0$.

In the second example, one possible solution is $a = [2, 2, 1, 0, 2, 0, 1]$. Its subarrays with bitwise XOR equal to $0$ are $(l, r) = (1, 2)$, $(2, 7)$, $(4, 4)$, and $(6, 6)$, so the beauty of $a$ is $7 - 4 = 3$. The array $[1, 2, 2, 1, 2]$ has beauty $5 - 2 = 3$ as well, but it is shorter. It can be shown that $3$ is the maximum possible beauty for $n = 2$, and that no array of length greater than $7$ achieves it.