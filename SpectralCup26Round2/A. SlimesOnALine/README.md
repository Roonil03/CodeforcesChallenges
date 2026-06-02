# A. Slimes on a Line
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
There are n slimes on a line, where slime i is at position ai on the line. You will perform the following operation some number of times (possibly none):
- select an integer x, then for each j (1 ≤ j ≤ n):
    - if a<sub>j</sub> < x then do a<sub>j</sub> := a<sub>j</sub> + 1.
    - if a<sub>j</sub> > x then do a<sub>j</sub> := a<sub>j</sub>−1.
    - if a<sub>j</sub> = x then do nothing. 

Determine the minimum number of operations to make all slimes occupy the same position.

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 100). The description of the test cases follows.

The first line of each testcase contains an integer n (2 ≤ n ≤ 1000) — the number of slimes.

The second line of each testcase contains n integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ 1000) — the initial positions of the slimes.

It is guaranteed that the sum of n over all test cases does not exceed 1000. 

### Output

For each testcase, output the minimum number of operations required to make all slimes occupy the same position.

### Examples:
**Sample Input:**
```
10
5
1 2 3 4 5
5
3 3 3 3 3
6
5 6 7 1 2 3
2
2 5
4
1 3 8 7
4
6 2 1 8
3
1 3 9
5
1 10 1 10 10
8
10 8 5 9 1 6 9 10
2
1 1000
```
**Sample Output:**
```
2
0
3
2
4
4
4
5
5
500
```
**Explaination:**
Note  
*Test Case 1:* We can perform 2 operations, both with x=3. The first operation updates the array of positions to a = `[2,3,3,3,4]`, and then the second operation updates it to a = `[3,3,3,3,3]`.  

*Test Case 2:* All the slimes are already at position 3, and hence 0 operations are needed.
