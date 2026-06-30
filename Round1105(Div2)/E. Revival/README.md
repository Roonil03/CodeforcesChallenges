# E. Revival
## Time limit per test: 2.5 seconds
## Memory limit per test: 256 megabytes
### Description:
Consider a permutation $p$ of $1, 2, \dots, n$. Let $s_i$ denote the number of inversions in the entire prefix $p_1, p_2, \dots, p_i$, defined as:

$$s_i = \sum_{1 \le x < y \le i} [p_x > p_y]$$

where the square parentheses denote the Iverson bracket notation.

For each position $i$ ($1 \le i \le n$), you are given a condition in the form of either $p_i = x$ or $s_i = x$. Your task is to reconstruct the original permutation $p$.

### Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

For each test case: The first line contains a single integer $n$ ($1 \le n \le 2 \cdot 10^5$). Each of the following $n$ lines contains a character $c$ ($c \in \{'p', 's'\}$) and an integer $x$:

* If $c = \text{'p'}$, it denotes that $p_i = x$, where $1 \le x \le n$.
* If $c = \text{'s'}$, it denotes that $s_i = x$, where $0 \le x \le \frac{i(i-1)}{2}$.

The sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

It is guaranteed that a valid permutation always exists.

### Output

For each test case, output $n$ integers representing the permutation $p$.

If there are multiple valid permutations, you can output any of them.

### Examples:
**Input:**
```
5
3
p 1
p 2
p 3
3
s 0
s 1
s 2
3
p 1
s 0
p 2
5
p 1
p 4
s 0
p 2
s 4
6
s 0
s 1
s 3
s 6
s 10
s 15
```
**Output:**
```
1 2 3
3 1 2
1 3 2
1 4 5 2 3
6 5 4 3 2 1
```
**Explaination:**  
In the first test case, the values of all elements are explicitly given, so the unique valid permutation is $\{1, 2, 3\}$.

In the second test case, we need to reconstruct the permutation from the prefix inversion counts:

* For position $2$, the number of inversions increases by $s_2 - s_1 = 1 - 0 = 1$. This means $p_2$ must be smaller than exactly one element before it, so $p_1 > p_2$.
* For position $3$, the number of inversions increases by $s_3 - s_2 = 2 - 1 = 1$. This means $p_3$ is smaller than exactly one element before it.

The only permutation satisfying these conditions is $\{3, 1, 2\}$.

In the third test case, we are given $p_1 = 1$ and $p_3 = 2$. The only remaining available value for $p_2$ is $3$. We can verify that the prefix $p_1, p_2$ (which is $1, 3$) has $0$ inversions, perfectly satisfying the condition $s_2 = 0$. Thus, the answer is $\{1, 3, 2\}$.

In the fifth test case, the given $s_i$ values perfectly match $\frac{i(i-1)}{2}$, which is the maximum possible number of inversions for a prefix of length $i$. This implies that every element must be smaller than all elements before it, meaning the permutation is strictly decreasing. Hence, the answer is $\{6, 5, 4, 3, 2, 1\}$.