# C. Those Who Are With Us
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
You are given a matrix of integers with n
 rows and m
 columns. The cell at the intersection of the i
-th row and the j
-th column contains the number a<sub>ij</sub>
.

You can perform the following operation exactly once:

- Choose two numbers 1 ≤ r ≤ n
 and 1 ≤ c ≤ m
.
- For all cells (i,j)
 in the matrix such that i=r
 or j=c
, decrease a<sub>ij</sub>
 by one.
You need to find the minimal possible maximum value in the matrix a
 after performing exactly one such operation.

### Input
Each test consists of multiple test cases. The first line contains a single integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases. The description of the test cases follows.

The first line of each test case contains two integers n
 and m
 (1 ≤ n ⋅ m ≤ 10<sup>5</sup>
) — the number of rows and columns in the matrix.

The next n
 lines of each test case describe the matrix a
. The i
-th line contains m
 integers a<sub>i1</sub>,a<sub>i2</sub>,…,a<sub>im</sub>
 (1 ≤ a<sub>ij</sub> ≤ 100
) — the elements in the i
-th row of the matrix.

It is guaranteed that the sum of n ⋅ m
 across all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output the minimum maximum value in the matrix a
 after performing exactly one operation.

### Examples:
**Input:**
```
10
1 1
1
1 2
1 2
2 1
2
1
2 2
4 2
3 4
3 4
1 2 3 2
3 2 1 3
2 1 3 2
4 3
1 5 1
3 1 3
5 5 5
3 5 1
4 4
1 3 3 2
2 3 2 2
1 2 2 1
3 3 2 3
2 2
2 2
1 2
3 2
1 2
2 1
1 2
3 3
2 1 1
1 2 1
1 1 2
```
**Output:**
```
0
1
1
3
2
4
3
1
1
2
```
**Explaination:**  
In the first three test cases, you can choose r=1
 and c=1
.

In the fourth test case, you can choose r=1
 and c=2
.
<img src="https://espresso.codeforces.com/c17767a6920a311231f92f0369c10b9b6a34d491.png"><br>

In the fifth test case, you can choose r=2
 and c=3
.
<img src="https://espresso.codeforces.com/45bddc965ec3fa397cb2b518c9236cd1e45f846e.png"><br>

In the sixth test case, you can choose r=3
 and c=2
.