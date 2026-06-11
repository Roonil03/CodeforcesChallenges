# B. Different Distances
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
Construct an array of 4 ⋅ n
 integers such that:

- each number 1, 2, …, n
 appears in the array exactly 4
 times;
- let p<sub>x,i</sub>
 be the position of the i
-th occurrence of number x
 in the array; then, for each x
 from 1
 to n
, the numbers (p<sub>x,2</sub> − p<sub>x,1</sub>), (p<sub>x,3</sub> − p<sub>x,2</sub>), (p<sub>x,4</sub> − p<sub>x,3</sub>)
 must be pairwise distinct.

For example, for n = 3
, one possible array is `[1,1,2,1,2,3,1,3,2,2,3,3]`
, because:
- p<sub>1,2</sub> − p<sub>1,1</sub> = 1, p<sub>1,3</sub> − p<sub>1,2</sub> = 2, p<sub>1,4</sub> − p<sub>1,3</sub> = 3
 — all numbers are distinct;
- p<sub>2,2</sub> − p<sub>2,1</sub> = 2, p<sub>2,3</sub> − p<sub>2,2</sub> = 4, p<sub>2,4</sub> − p<sub>2,3</sub> = 1
 — all numbers are distinct;
- p<sub>3,2</sub> − p<sub>3,1</sub> = 2, p<sub>3,3</sub> − p<sub>3,2</sub> = 3, p<sub>3,4</sub> − p<sub>3,3</sub> = 1
 — all numbers are distinct.
### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 200
). The description of the test cases follows.

The only line of each test case contains one integer n
 (2 ≤ n ≤ 200
).

### Output
For each test case, output the required array. If there are several suitable arrays, output any of them. It can be shown that at least one suitable array always exists under the constraints of the problem.

### Examples:
**Input:**
```
3
3
4
5
```
**Output:**
```
1 1 2 1 2 3 1 3 2 2 3 3
1 4 3 4 4 3 3 1 2 1 2 3 4 2 2 1 
5 3 2 4 1 2 2 4 4 5 3 5 2 1 3 1 5 4 1 3 
```