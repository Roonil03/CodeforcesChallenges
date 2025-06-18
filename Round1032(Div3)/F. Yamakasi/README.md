# F. Yamakasi
## Time limit per test: 3 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given an array of integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 and two integers s
 and x
. Count the number of subsegments of the array whose sum of elements equals s
 and whose maximum value equals x
.

More formally, count the number of pairs 1 ≤ l ≤ r ≤ n
 such that:

- a<sub>l</sub>+a<sub>l+1</sub>+…+a<sub>r</sub> = s
.
- max(a<sub>l</sub>,a<sub>l+1</sub>,…,a<sub>r</sub>) = x
.
### Input
Each test consists of multiple test cases. The first line contains a single integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases. The description of the test cases follows.

The first line of each test case contains three integers n
, s
, and x
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
, −2 ⋅ 10<sup>14</sup> ≤ s ≤ 2 ⋅ 10<sup>14</sup>
, −10<sup>9</sup> ≤ x ≤ 10<sup>9</sup>
).

The second line of each test case contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (−10<sup>9</sup> ≤ a<sub>i</sub> ≤ 10<sup>9</sup>
).

It is guaranteed that the sum of n
 across all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output the number of subsegments of the array whose sum of elements equals s
 and whose maximum value equals x
.

### Examples:
**Input:**
```
9
1 0 0
0
1 -2 -1
-2
3 -1 -1
-1 1 -1
6 -3 -2
-1 -1 -1 -2 -1 -1
8 3 2
2 2 -1 -2 3 -1 2 2
9 6 3
1 2 3 1 2 3 1 2 3
13 7 3
0 -1 3 3 3 -2 1 2 2 3 -1 0 3
2 -2 -1
-2 -1
2 -2 -1
-1 -2
```
**Output:**
```
1
0
2
0
2
7
8
0
0
```
**Explaination:**  
In the first test case, the suitable subsegment is l=1
, r=1
.

In the third test case, the suitable subsegments are l=1
, r=1
 and l=3
, r=3
.

In the fifth test case, the suitable subsegments are l=1
, r=3
 and l=6
, r=8
.

In the sixth test case, the suitable subsegments are those for which r=l+2
.

In the seventh test case, the following subsegments are suitable:
- l=1
, r=7
.
- l=2
, r=7
.
- l=3
, r=6
.
- l=4
, r=8
.
- l=7
, r=11
.
- l=7
, r=12
.
- l=8
, r=10
.
- l=9
, r=13
.