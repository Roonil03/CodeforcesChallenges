# A. Another Popcount Problem
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
You are given two integers $n$ and $k$.

Your task is to construct a sequence $a$ consisting of $k$ non-negative integers $a_1, a_2, \dots, a_k$ such that:

* $\sum_{i=1}^{k} a_i \leq n$
* The total number of set bits, i.e., $\sum_{i=1}^{k} \text{popcount}(a_i)$, is as large as possible.

You only need to output the maximum possible value of $\sum_{i=1}^{k} \text{popcount}(a_i)$.

Here, $\text{popcount}(x)$ denotes the number of 1 bits in the binary representation of $x$. For example, $\text{popcount}(6) = \text{popcount}((110)_2) = 2$, and $\text{popcount}(0) = 0$.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \leq t \leq 10^3$). The description of the test cases follows.

Each of the next $t$ lines contains two integers $n$ and $k$ ($1 \leq n, k \leq 10^6$) — the maximum allowed sum of the sequence and the length of the sequence, respectively.

### Output
For each test case, output a single integer — the maximum possible value of $\sum_{i=1}^{k} \text{popcount}(a_i)$.

### Examples:
**Input:**
```
6
2 1
3 1
6 2
14142 137205
1000000 100
1000000 1000000
```
**Output:**
```
1
2
4
14142
1322
1000000
```
**Explaination:**  
In the first test case, n=2
 and k=1
. We can choose a = `[1]`
 or a = `[2]`
. In both cases, the sum of popcounts is 1.

In the second test case, n = 3
 and k = 1
. We can choose a = `[3]`
, since (3)<sub>2</sub> = (11)<sub>2</sub>
, popcount(3)=2.

In the third test case, n = 6
 and k = 2
. We can choose a = `[3,3]`
. The sum is 3 + 3 = 6 ≤ 6
, and the total popcount is popcount(3)+popcount(3)=2+2=4.