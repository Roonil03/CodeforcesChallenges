# B. Shrink
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes

### Description:
A shrink operation on an array a
 of size m
 is defined as follows:

- Choose an index i
 (2≤i≤m−1
) such that a<sub>i</sub>>a<sub>i−1</sub>
 and a<sub>i</sub>>a<sub>i+1</sub>
.
Remove a<sub>i</sub>
 from the array.
Define the score of a permutation∗
 p
 as the maximum number of times that you can perform the shrink operation on p
.

Yousef has given you a single integer n
. Construct a permutation p
 of length n
 with the maximum possible score. If there are multiple answers, you can output any of them.

<sub>∗
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

### Input:
The first line of the input contains an integer t
 (1≤t≤10<sup>3</sup>
) — the number of test cases.

Each test case contains an integer n
 (3≤n≤2⋅10<sup>5</sup>
) — the size of the permutation.

It is guaranteed that the sum of n
 over all test cases does not exceed 2⋅10<sup>5</sup>
.

### Output:
For each test case, output any permutation p<sub>1</sub>,p<sub>2</sub>,…,p<sub>n</sub>
 that maximizes the number of shrink operations.

### Examples:
**Input:**
```
2
3
6
```
**Output:**
```
1 3 2
2 3 6 4 5 1
```
**Explaination:**  
In the first test case:

- We choose p=`[1,3,2]`
.
- Choose index 2
, and remove p<sub>2</sub>
 from the array. The array becomes p=`[1,2]`
.

It can be shown that the maximum number of operations we can perform is 1
. Another valid answer is p=`[2,3,1]`
.

In the second test case:

- We choose p=`[2,3,6,4,5,1]`
.
- Choose index 5
, and remove p<sub>5</sub>
 from the array. The array becomes p=`[2,3,6,4,1]`
.
- Choose index 3
, and remove p<sub>3</sub>
 from the array. The array becomes p=`[2,3,4,1]`
.
- Choose index 3
, and remove p<sub>3</sub>
 from the array. The array becomes p=`[2,3,1]`
.
- Choose index 2
, and remove p<sub>2</sub>
 from the array. The array becomes p=`[2,1]`
.

The maximum number of operations we can perform is 4
. Any permutation with a score of 4
 is valid.