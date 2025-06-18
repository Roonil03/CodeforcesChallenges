# A. Letter Home
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
You are given an array of distinct integers x<sub>1</sub>,x<sub>2</sub>,…,x<sub>n</sub>
 and an integer s
.

Initially, you are at position pos=s
 on the X
 axis. In one step, you can perform exactly one of the following two actions:
- Move from position pos
 to position pos+1
.
- Move from position pos
 to position pos−1
.

A sequence of steps will be considered successful if, during the entire journey, you visit each position x<sub>i</sub>
 on the X
 axis at least once. Note that the initial position pos=s
 is also considered visited.

Your task is to determine the minimum number of steps in any successful sequence of steps.

### Input
Each test consists of multiple test cases. The first line contains a single integer t
 (1 ≤ t ≤ 1000
) — the number of test cases. The description of the test cases follows.

The first line of each test case contains two integers n
 and s
 (1 ≤ n ≤ 10
, 1 ≤ s ≤ 100
) — the number of positions to visit and the starting position.

The second line of each test case contains n
 integers x<sub>1</sub>,x<sub>2</sub>,…,x<sub>n</sub>
 (1 ≤ x<sub>i</sub> ≤ 100
). It is guaranteed that for all 1 ≤ i < n
, it holds that x<sub>i</sub> < x<sub>i+1</sub>
.

### Output
For each test case, output the minimum number of steps in any successful sequence of steps.

### Examples:
**Input:**
```
12
1 1
1
1 2
1
1 1
2
2 1
2 3
2 2
1 3
2 3
1 2
3 1
1 2 3
3 2
1 3 4
3 3
1 2 3
4 3
1 2 3 10
5 5
1 2 3 6 7
6 6
1 2 3 9 10 11
```
**Output:**
```
0
1
1
2
3
2
2
4
2
11
8
15
```
**Explaination:**  
In the first test case, no steps need to be taken, so the only visited position will be 1
.

In the second test case, the following path can be taken: 2→1
. The number of steps is 1
.

In the third test case, the following path can be taken: 1→2
. The number of steps is 1
.

In the fifth test case, the following path can be taken: 2→1→2→3
. The number of steps is 3
.