# G. Modular Sorting
## Time limit per test: 5 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given an integer m
 (2≤m≤5⋅105
) and an array a
 consisting of nonnegative integers smaller than m
.

Answer queries of the following form:

- 1
 i
 x
: assign a<sub>i</sub>:=x
- 2
 k
: in one operation, you may choose an element ai
 and assign a<sub>i</sub>:=(a<sub>i</sub>+k)(mod m)
∗
 — determine if there exists some sequence of (possibly zero) operations to make a
 nondecreasing†
.
Note that instances of query 2
 are independent; that is, no actual operations are taking place. Instances of query 1
 are persistent.
<sub>
∗
a(mod m)
 is defined as the unique integer b
 such that 0 ≤ b < m
 and a−b
 is an integer multiple of m
.

†
An array a
 of size n
 is called nondecreasing if and only if a<sub>i</sub>≤a<sub>i+1</sub>
 for all 1≤i<n
.</sub>

### Input
The first line contains an integer t
 (1≤t≤10<sup>4</sup>
)  — the number of test cases.

The first line of each test case contains three integers, n
, m
, and q
 (2≤n≤10<sup>5</sup>
, 2≤m≤5⋅10<sup>5</sup>
, 1≤q≤10<sup>5</sup>
) — the size of the array a
, the integer m
, and the number of queries.

The second line of each test case contains n
 integers, a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (0 ≤ a<sub>i</sub> < m
).

Then follows q
 lines. Each line is of one of the following forms:

- 1
 i
 x
 (1≤i≤n
, 0≤x<m
)
- 2
 k
 (1≤k<m
)

It is guaranteed that the sum of n
 and the sum of q
 over all test cases each do not exceed 10<sup>5</sup>
.

### Output
For each instance of query 2
, output on a single line "YES" if there exists some sequence of (possibly zero) operations to make a
 nondecreasing, and "NO" otherwise.

You can output the answer in any case (upper or lower). For example, the strings "yEs", "yes", "Yes", and "YES" will be recognized as positive responses.

### Examples:
**Input:**
```
2
7 6 6
4 5 2 2 4 1 0
2 4
1 4 5
2 4
2 3
1 7 2
2 3
8 8 3
0 1 2 3 4 5 6 7
2 4
1 3 4
2 4
```
**Output:**
```
YES
NO
NO
YES
YES
NO
```
**Explaination:**  
In the first sample, the array is initially:
```
4	5	2	2	4	1	0
```
By applying the operation twice on a<sub>1</sub>
, twice on a<sub>2</sub>
, once on a<sub>5</sub>
, twice on a<sub>6</sub>
, and once on a<sub>7</sub>
, the array becomes:
```
0	1	2	2	2	3	4
```
which is in nondecreasing order.
After the second query, the array becomes:
```
4	5	2	5	4	1	0
```
and it can be shown that it is impossible to sort this with operations of the form a<sub>i</sub> := (a<sub>i</sub>+4)(mod 6)
, and it is also impossible to sort this with operations of the form a<sub>i</sub>:=(a<sub>i</sub> + 3)(mod 6)
.