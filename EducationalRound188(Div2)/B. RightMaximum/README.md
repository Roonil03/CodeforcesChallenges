# B. Right Maximum
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
You are given an array a
 consisting of n
 integers.

While the array is not empty, an operation is performed consisting of two steps:
- first, the maximum element in the array is chosen (if there are multiple maximum elements, the rightmost maximum is chosen);
- then, all elements after the chosen element, including it, are removed from the array.

Your task is to calculate the number operations that will be performed before the array becomes empty.

Input
The first line contains one integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases.

Each test case consists of two lines:
- the first line contains one integer n
 (2 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
);
- the second line contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ n
).

Additional constraint on the input: the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, print one integer — the number of operations that will be performed.

### Example:
**Input:**
```
4
5
2 1 2 3 1
6
1 2 3 4 5 6
3
3 2 1
4
1 3 3 1
```
**Output:**
```
3
6
1
3
```
**Notes:**  
In the first example, the array is `[2,1,2,3,1]`
. The following operations are performed on it:

first, the 4
-th element is chosen. The last two elements are removed, and the array becomes `[2,1,2]`
;
then, the 3
-rd element is chosen. The last element is removed, and the array becomes `[2,1]`
;
then, the 1
-st element is chosen. Both elements are removed, so the array becomes empty after 3
 operations.