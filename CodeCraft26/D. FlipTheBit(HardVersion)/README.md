# D. Flip the Bit (Hard Version)
## Time limit per test: 1.5 seconds
## Memory limit per test: 256 megabytes
### Description:
**This is the hard version of the problem. The difference between the versions is that in this version, there can be multiple special indices (1 ≤ k ≤ n
). You can hack only if you solved all versions of this problem.**

You are given a binary array a
 of length n
 and k
 special indices p<sub>1</sub>, p<sub>2</sub>, …, p<sub>k</sub>
 (1 ≤ p<sub>i</sub> ≤ n
). It is given that the values a<sub>i</sub>
 of all elements at special indices are the same (i.e., a<sub>p<sub>1</sub></sub> = a<sub>p<sub>2</sub></sub> = … = a<sub>p<sub>k</sub></sub>
).

In one operation, you can choose a range `[l,r]`
 (1 ≤ l ≤ r ≤ n
) such that the range contains at least one special index (l ≤ p<sub>i</sub> ≤ r
) and flip all bits a<sub>j</sub>
 for l ≤ j ≤ r
. Flipping a bit changes 0
 to 1
 and 1
 to 0
.

Let x
 denote the value at special indices before any operations are applied. Find the minimum number of operations required to make all elements of the array equal to x
.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains two integers n
 and k
 (1 ≤ k ≤ n ≤ 2 ⋅ 10<sup>5</sup>
) — the length of the array and the number of special indices.

The second line contains n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (0 ≤ a<sub>i</sub> ≤ 1
) — the elements of the array.

The third line contains k
 integers p<sub>1</sub>, p<sub>2</sub>, …, p<sub>k</sub>
 (1 ≤ p<sub>1</sub> < p<sub>2</sub> < … < p<sub>k</sub> ≤ n
) — the special indices. It is guaranteed that a<sub>p<sub>1</sub></sub> = a<sub>p<sub>2</sub></sub> = … = a<sub>p<sub>k</sub></sub>
.

It is guaranteed that the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output a single integer — the minimum number of operations required.

### Examples:
**Input:**
```
6
2 1
1 0
1
3 2
0 1 0
1 3
5 5
1 1 1 1 1
1 2 3 4 5
9 5
0 1 0 0 1 0 0 1 0
3 4 6 7 9
13 4
1 0 0 1 0 1 0 1 1 0 1 0 1
4 8 11 13
15 3
1 0 1 0 1 0 1 0 1 0 1 0 1 0 1
3 11 13
```
**Output:**
```
2
2
0
3
5
8
```
**Notes:**  
For the first test case, you can choose the range `[1,2]`
 and flip both the bits to get `[0,1]`
. Then you can choose the range `[1,1]`
 and flip the first bit to get `[1,1]`
.

For the second test case, you can choose the range `[1,2]`
 and flip the bits to get `[1,0,0]`
. Then you can choose the range `[1,1]`
 and flip the first bit to get `[0,0,0]`
.

For the fourth test case, you can choose the range `[2,8]`
 and flip the bits to get `[0,0,1,1,0,1,1,0,0]`
. Then you can choose the range `[3,4]`
 and flip the bits to get `[0,0,0,0,0,1,1,0,0]`
. Then you can choose the range `[6,7]`
 and flip the bits to get `[0,0,0,0,0,0,0,0,0]`
. Note that, while there can be other ways to achieve this, it can be proved that the minimum number of operations is 3
.
