# C1. We Be Flipping (Easy Version)
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
**This is the easy version of the problem. The difference between the versions is that in this version, you must minimise the sum. You can hack only if you solved all versions of this problem.**

You have an array a of length n which consists of non-zero (but possibly negative) integers. You will perform the following operation at most n times (possibly none):
- select an index i (1 ≤ i ≤ n) such that a<sub>i</sub> > 0
- then for each j where 1 ≤ j ≤ i do a<sub>j</sub> := −a<sub>j</sub>. 

Output a valid sequence of operations of length at most n which minimises the sum of a at the end.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each testcase contains an integer n (2 ≤ n ≤ 2 ⋅ 10<sup>5</sup>) — the length of the array a.

The second line of each testcase contains n integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub> (−10<sup>9</sup> ≤ a<sub>i</sub> ≤ 10<sup>9</sup>, a<sub>i</sub> ≠ 0).

It is guaranteed that the sum of n over all test cases does not exceed 2 ⋅ 10<sup>5</sup>.
### Output

For each testcase, output a single integer k (0 ≤ k ≤ n) — the number of operations you will perform.

Now output k integers b<sub>1</sub>, …, b<sub>k</sub> where b<sub>i</sub> is the index you perform the ith operation on.

After performing the operations, the sum of a should be minimal.

### Examples:
**Sample Input:**
```
3
5
-1 -2 -3 -5 -4
5
-1 -2 3 -5 4
4
5 7 10 19
```
**Sample Output:**
```
0

4
3 5 4 2
1
4
```
**Explaination:**  
In the first testcase, the sum is already minimised. So no operations are required.

In the second testcase, the operations are made as follows:

* $[-1, -2, 3, -5, 4] \xrightarrow{i=3} [\textcolor{red}{1, 2, -3}, -5, 4]$
* $[1, 2, -3, -5, 4] \xrightarrow{i=5} [\textcolor{red}{-1, -2, 3, 5, -4}]$
* $[-1, -2, 3, 5, -4] \xrightarrow{i=4} [\textcolor{red}{1, 2, -3, -5}, -4]$
* $[1, 2, -3, -5, -4] \xrightarrow{i=2} [\textcolor{red}{-1, -2}, -3, -5, -4]$

This has a sum of $-15$, which is the minimum possible.

# C2. We Be Flipping (Hard Version)
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
**This is the hard version of the problem. The difference between the versions is that in this version, you must maximise the sum. You can hack only if you solved all versions of this problem.**

You have an array a of length n which consists of non-zero (but possibly negative) integers. You will perform the following operation at most n times (possibly none):
- select an index i (1 ≤ i ≤ n) such that a<sub>i</sub> > 0
- then for each j where 1 ≤ j ≤ i do a<sub>j</sub> := −a<sub>j</sub>. 

Output a valid sequence of operations of length at most n which maximises the sum of a at the end.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each testcase contains an integer n (2 ≤ n ≤ 2 ⋅ 10<sup>5</sup>) — the length of the array a.

The second line of each testcase contains n integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub> (−10<sup>9</sup> ≤ a<sub>i</sub> ≤ 10<sup>9</sup>, a<sub>i</sub> ≠ 0).

It is guaranteed that the sum of n over all test cases does not exceed 2 ⋅ 10<sup>5</sup>.
### Output

For each testcase, output a single integer k (0 ≤ k ≤ n) — the number of operations you will perform.

Now output k integers b<sub>1</sub>, …, b<sub>k</sub> where b<sub>i</sub> is the index you perform the ith operation on.

After performing the operations the sum of a should be maximal.

### Examples:
**Sample Input:**
```
5
5
-1 -2 -3 -5 -4
4
5 7 10 19
5
1 -3 2 -1 10
4
16 -13 -18 -16
11
2 -10 -11 3 -10 15 7 18 16 17 -9
```
**Sample Output:**
```
0

0

2
1 3
0

6
6 3 1 5 4 7
```
**Explaination:**  
In the first testcase, no operations are possible.

In the second testcase, the sum is already maximal.

In the third testcase, operations are made as follows:

* $[1, -3, 2, -1, 10] \xrightarrow{i=1} [\textcolor{red}{-1}, -3, 2, -1, 10]$
* $[-1, -3, 2, -1, 10] \xrightarrow{i=3} [\textcolor{red}{1, 3, -2}, -1, 10]$

This has sum 11, which can be proven to be maximal.