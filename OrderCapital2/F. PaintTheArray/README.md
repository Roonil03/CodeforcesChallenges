# F. Paint the Array
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Consider an array of length n
, and a fixed integer m
. A painting operation is defined as follows:
- Choose an interval of length m
 and paint it with the values 1, 2, …, m
 from left to right.
Formally, choose an integer l
 such that 1 ≤ l ≤ n − m + 1
. Then, for every 1 ≤ i ≤ m
, position l + i − 1
 is painted with value i
.

If a position is painted multiple times, only the value painted most recently remains.

An array is called valid if it can be obtained by performing some painting operations such that every position is painted at least once.

Given an array a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 with 1 ≤ a<sub>i</sub> ≤ m
, find the minimum number of modifications needed to make it valid. One modification changes one element to any integer between 1
 and m
.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains two integers n
 and m
 (1 ≤ m ≤ n ≤ 5 ⋅ 10<sup>5</sup>
) — the length of the array and the length of each painted interval.

The second line contains n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ m
).

It is guaranteed that the sum of n
 over all test cases does not exceed 5 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output a single integer — the minimum number of modifications needed to make it valid.

### Examples:
**Input:**
```
15
5 3
1 2 3 2 3
4 3
1 2 2 3
5 3
2 1 2 3 2
5 3
2 2 2 2 2
5 4
1 1 3 4 1
6 3
1 1 1 2 1 1
8 5
1 2 1 2 3 4 5 1
5 3
2 3 1 1 2
8 4
1 2 3 2 3 2 3 4
4 4
4 3 2 1
5 1
1 1 1 1 1
7 3
3 3 3 2 1 1 1
10 3
1 2 3 1 2 2 3 1 2 3
7 3
1 3 2 3 2 1 2
10 4
1 4 3 3 2 3 4 4 2 2
```
**Output:**
```
0
1
2
3
2
2
1
4
2
4
0
4
1
3
4
```
**Explaination:**  
In the transformations below, the underlined positions are exactly the positions painted in the latest operation.

In the first test case, the array is already valid. It can be obtained by the following painting operations:
`[−,−,−,−,−] → [−,−,1–,2–,3–] → [1–,2–,3–,2,3]`.  
Therefore no modification is needed, and the answer is 0
.

In the second test case, since n=4
 and m = 3
, every valid array must be obtained by painting both intervals `[1,3]`
 and `[2,4]`
. For example,
`[−,−,−,−] → [−,1–,2–,3–] → [1–,2–,3–,3]`.  
This gives the valid array `[1,2,3,3]`
. The given array `[1,2,2,3]`
 can be changed into it by modifying only the third element, so the answer is 1
.

In the third test case, one closest valid array is `[1,1,2,3,3]`
, which can be obtained as follows:
`[−,−,−,−,−] → [1–,2–,3–,−,−] → [1,2,1–,2–,3–] → [1,1–,2–,3–,3]`.  
The given array `[2,1,2,3,2]`
 differs from `[1,1,2,3,3]`
 in two positions. It can be proven that one modification is not enough, so the answer is 2.




