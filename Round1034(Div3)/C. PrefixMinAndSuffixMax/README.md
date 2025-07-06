# C. Prefix Min and Suffix Max
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given an array a
 of distinct integers.

In one operation, you may either:

- choose a nonempty prefix∗
 of a
 and replace it with its minimum value, or
- choose a nonempty suffix†
 of a
 and replace it with its maximum value.
Note that you may choose the entire array a
.

For each element a<sub>i</sub>
, determine if there exists some sequence of operations to transform a
 into `[a`<sub>`i`</sub>`]`
; that is, make the array a
 consist of only one element, which is a<sub>i</sub>
. Output your answer as a binary string of length n
, where the i
-th character is 1
 if there exists a sequence to transform a
 into `[a`<sub>`i`</sub>`]`
, and 0
 otherwise.

<sub>∗
 A prefix of an array is a subarray consisting of the first k
 elements of the array, for some integer k
.

†
A suffix of an array is a subarray consisting of the last k
 elements of the array, for some integer k
.</sub>
 ### Input
The first line contains an integer t
 (1≤t≤10<sup>4</sup>
)  — the number of test cases.

The first line of each test case contains one integer n
 (2≤n≤2⋅10<sup>5</sup>
) — the size of the array a
.

The second line of each test case contains n
 integers, a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1≤a<sub>i</sub>≤10<sup>6</sup>
). It is guaranteed that all a<sub>i</sub>
 are distinct.

It is guaranteed that the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output a binary string of length n
  — the i
-th character should be 1
 if there exists a sequence of operations as described above, and 0
 otherwise.

### Examples:
**Input:**
```
3
6
1 3 5 4 7 2
4
13 10 12 20
7
1 2 3 4 5 6 7
```
**Output:**
```
100011
1101
1000001
```
**Explaination:**  
In the first sample, you can first choose the prefix of size 3
. Then the array is transformed into
```
1	4	7	2
```
Next, you can choose the suffix of size 2
. Then the array is transformed into
```
1	4	7
```
Finally, you can choose the prefix of size 3
. Then the array is transformed into
```
1
```
So we see that it is possible to transform a
 into `[1]`
.
It can be shown that it is impossible to transform a
 into `[3]`
.