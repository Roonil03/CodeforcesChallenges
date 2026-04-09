# E. Definitely Larger
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given a permutation p
<sup>∗</sup>
 of integers from 1
 to n
.

For an arbitrary permutation q
 of length n
, we say that index j
 dominates index i
 if and only if all of the following conditions hold:

- j > i
,
- p<sub>j</sub> > p<sub>i</sub>
, and
- q<sub>j</sub> > q<sub>i</sub>
.

You are also given an array d
 of length n
, where d<sub>i</sub>
 denotes the number of indices that dominate index i
.

Your task is to construct a permutation q
 of integers from 1
 to n
 such that for every index i
 (1 ≤ i ≤ n
), the number of indices j
 that dominate i
 is exactly d<sub>i</sub>
.

If such q
 exists, output any valid one. Otherwise, report that it does not exist.
<sub>
<sup>∗</sup>
A permutation of length n
 is an array consisting of n
 distinct integers from 1
 to n
 in arbitrary order. For example, `[2,3,1,5,4]`
 is a permutation, but `[1,2,2]`
 is not a permutation (2
 appears twice in the array), and `[1,3,4]`
 is also not a permutation (n=3
 but there is 4
 in the array).</sub>
### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 1000
). The description of the test cases follows.

The first line of each test case contains an integer n
 (1 ≤ n ≤ 5000
).

The second line of each test case contains n
 integers p<sub>1</sub>, p<sub>2</sub>, …, p<sub>n</sub>
. It is guaranteed that p
 forms a permutation of integers from 1
 to n
.

The third line of each test case contains n
 integers d<sub>1</sub>, d<sub>2</sub>, …, d<sub>n</sub>
 (0 ≤ d<sub>i</sub> ≤ n
).

It is guaranteed that the sum of n
 over all test cases does not exceed 5000
.

### Output
For each test case:
- If it is possible to construct q
, output a line containing n
 integers q<sub>1</sub>, q<sub>2</sub>, …, q<sub>n</sub>
 — a valid permutation. If multiple valid answers exist, you may output any of them.
Otherwise, output −1
.

### Examples:
**Input:**
```
7
3
2 3 1
1 0 0
4
3 4 1 2
2 1 1 0
5
2 3 1 4 5
2 2 1 1 0
1
1
0
5
3 1 4 2 5
1 1 1 1 0
4
3 4 2 1
1 1 1 0
8
7 6 3 1 2 5 4 8
1 1 2 2 2 1 1 0
```
**Output:**
```
1 2 3
-1
2 1 4 3 5
1
2 4 1 3 5
-1
1 2 4 6 5 3 7 8
```
**Note:**  
In the first test case, a valid output is q = `[1,2,3]`
:
- For i = 1
, we have p<sub>1</sub> = 2
 and q<sub>1</sub> = 1
. The index j = 2
 dominates i = 1
 because 2 > 1
, p<sub>2</sub> = 3 > p<sub>1</sub>
, and q<sub>2</sub> = 2 > q<sub>1</sub>
. Index j = 3
 does not dominate i = 1
 because p<sub>3</sub> = 1 < p<sub>1</sub>
. Thus, exactly 1
 index dominates i = 1
, which matches d<sub>1</sub> = 1
.
- For i = 2
, we have p<sub>2</sub> = 3
 and q<sub>2</sub> = 2
. The only larger index is j = 3
, but p<sub>3</sub> = 1 < p<sub>2</sub>
, so it does not dominate. Thus, 0
 indices dominate i = 2
, matching d<sub>2</sub> = 0
.
- For i = 3
, there are no larger indices j > 3
, so 0
 indices dominate, matching d<sub>3</sub> = 0
.

In the second test case, p = `[3,4,1,2]`
 and d = `[2,1,1,0]`
. Let's look at i = 2
, where p<sub>2</sub> = 4
. Since there is no index j > 2
 with p<sub>j</sub> > p<sub>2</sub>
, no index can dominate i = 2
. However, the array d
 requires d<sub>2</sub> = 1
, which makes it impossible to construct a valid permutation q
. Hence, the answer is -1.



