# D. 1709
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given two arrays of integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 and b<sub>1</sub>,b<sub>2</sub>,…,b<sub>n</sub>
. It is guaranteed that each integer from 1
 to 2⋅n
 appears in exactly one of the arrays.

You need to perform a certain number of operations (possibly zero) so that both of the following conditions are satisfied:

- For each 1 ≤ i < n
, it holds that a<sub>i</sub> < a<sub>i+1</sub>
 and b<sub>i</sub> < b<sub>i+1</sub>
.
- For each 1 ≤ i ≤ n
, it holds that a<supb>i</sub> < b<sub>i</sub>
.
During each operation, you can perform exactly one of the following three actions:

- Choose an index 1 ≤ i < n
 and swap the values a<sub>i</sub>
 and ai+1
.
- Choose an index 1 ≤ i < n
 and swap the values values b<sub>i</sub>
 and bi+1
.
- Choose an index 1 ≤ i ≤ n
 and swap the values values a<sub>i</sub>
 and bi
.  
You do not need to minimize the number of operations, but the total number must not exceed 1709
. Find any sequence of operations that satisfies both conditions.

### Examples:
**Input**
```
0
1
3 1
1
2 1
1
3 2
9
3 1
3 2
3 3
1 1
2 1
2 2
1 2
1 1
2 1
6
2 2
1 1
1 2
2 1
3 1
3 2
```
**Output**
```
0
1
3 1
1
2 1
1
3 2
9
3 1
3 2
3 3
1 1
2 1
2 2
1 2
1 1
2 1
6
2 2
1 1
1 2
2 1
3 1
3 2
```
**Explaination:**  
In the first test case, a<sub>1</sub> < b<sub>1</sub>
, so no operations need to be applied.

In the second test case, a<sub>1</sub> < b<sub>1</sub>
. After applying the operation, these values will be swapped.

In the third test case, after applying the operation, a=`[1,3]`
 and b=`[2,4]`
.

In the fourth test case, after applying the operation, a=`[1,2]`
 and b=`[3,4]`
.