# D. Me When Median Problem
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given two arrays of positive integers a and b, both of length n. You will perform the following operation exactly n − 1 times:
- let m be the current length of a and b, note that the lengths will always be equal.
- select an integer i (1 ≤ i < m):
    - let S be the multiset {a<sub>i</sub>, a<sub>i+1</sub>, b<sub>i</sub>, b<sub>i+1</sub>}
    - sort the elements of S such that s<sub>1</sub> ≤ s<sub>2</sub> ≤ s<sub>3</sub> ≤ s<sub>4</sub>.
    - now replace a<sub>i</sub>, a<sub>i+1</sub> with s<sub>2</sub> and b<sub>i</sub>, b<sub>i+1</sub> with s<sub>3</sub>. More formally, replace a with [a<sub>1</sub>, a<sub>2</sub>, …, a<sub>i−1</sub>, s<sub>2</sub>, a<sub>i+2</sub>, …, a<sub>m</sub>], and replace b with [b<sub>1</sub>, b<sub>2</sub>, …, b<sub>i−1</sub>, s<sub>3</sub>, b<sub>i+2</sub>, …,b<sub>m</sub>]. 

After performing all operations, there will be exactly 1 element remaining in both a and b. Determine the maximum value of min(a<sub>1</sub>, b<sub>1</sub>) attainable if you perform operations optimally.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each testcase contains an integer n (1 ≤ n ≤ 10<sup>5</sup>) — the length of the arrays a and b.

The second line of each testcase contains n integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ 2 ⋅ n).

The third line of each testcase contains n integers b<sub>1</sub>, b<sub>2</sub>, …, b<sub>n</sub> (1 ≤ b<sub>i</sub> ≤ 2 ⋅ n).

It is guaranteed that the sum of n over all test cases does not exceed 10<sup>5</sup>.
### Output

For each testcase, output the maximum value of min(a<sub>1</sub>,b<sub>1</sub>) attainable.

### Examples:
**Sample Input:**
```
6
1
1
2
3
2 4 5
1 3 6
4
7 5 4 8
4 6 7 8
8
8 7 13 11 1 10 4 5
11 11 12 8 9 2 3 13
9
16 1 9 12 5 18 10 10 16
14 6 7 11 12 17 18 3 17
6
3 6 12 4 10 12
2 3 2 7 8 9
```
**Sample Output:**
```
1
3
6
8
14
8
```
**Explaination:**  
In the first example, we do not need to perform any operations, so the answer is just $\min(1, 2)$ which is $1$.

For the second example, we can do the following sequence of moves:

* select $i = 1$ and then:
    * $S = \{2, 4, 1, 3\}, s_1 = 1, s_2 = 2, s_3 = 3, s_4 = 4$
    * $a = [\textcolor{red}{2, 4}, 5] \to [2, 5]$
    * $b = [\textcolor{red}{1, 3}, 6] \to [3, 6]$
* select $i = 1$ and then
    * $S = \{2, 5, 3, 6\}, s_1 = 2, s_2 = 3, s_3 = 5, s_4 = 6$
    * $a = [\textcolor{red}{2, 5}] \to [3]$
    * $b = [\textcolor{red}{3, 6}] \to [5]$

The answer is then $\min(3, 5)$ which is $3$, it can be proven that this is optimal.