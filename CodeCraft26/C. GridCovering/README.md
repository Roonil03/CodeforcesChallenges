# C. Grid Covering
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Prakul has been working hard setting problems for Codecraft. When he's deep in thought, he likes jumping around his room in weird but specific ways. After some time, he wonders if jumping in this manner can cover all the tiles of his room.

Prakul's room can be considered as a grid A
 consisting of n
 rows and m
 columns. He starts walking from A<sub>1,1</sub>
. If he is currently at A<sub>i,j</sub>
 he can do either of the following moves:
- Jump b
 steps right to A<sub>i,((j+b−1) mod m) + 1</sub>
, or
- Jump a
 steps down to A<sub>((i + a − 1) mod n) + 1 , j</sub>
.

A special restriction is that he can start with either move, but must alternate between them.

Note that his room is quite special as it wraps around itself. Moving one step right from the m
-th column places him in the 1
-st column. Similarly, moving one step down from the n
-th row places him in the 1
-st row.

Since he still has to work on problemsetting, he needs your help. Determine if he can visit all tiles in A
 in a finite number of jumps.


### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The only line of each test case contains four integers n, m, a, b
 (1 ≤ n, m, a, b ≤ 10<sup>9</sup>
).

### Output
For each test case, print "`YES`" if Prakul can cover all the tiles of his room, and "`NO`" otherwise.

You can output the answer in any case (upper or lower). For example, the strings "`yEs`", "`yes`", "`Yes`", and "`YES`" will be recognized as positive responses.

### Examples:
**Input:**
```
10
1 1 1 1
2 2 1 1
4 2 2 1
6 9 6 7
67 42 42 67
3411 4134 32 23
90234 143124 232 323
69387963 98793214 9791 4324786
985865 578977 899368 447605
1000000000 1000000000 1000000000 1000000000
```
**Output:**
```
YES
YES
NO
NO
YES
NO
NO
NO
YES
NO
```
**Note:**  
n the second test case: n=2
, m=2
, a=1
, b=1
. One possible way to cover the grid is if Prakul starts by moving down:
- Start at `(1,1)`
.
- Jump down a = 1
 tile: `(1,1)` → `(2,1)`
.
- Jump right b = 1
 tile: `(2,1)` → `(2,2)`
.
- Jump down a = 1
 tile: `(2,2)` → `(1,2)`
.
- Jump right b = 1
 tile: `(1,2)` → `(1,1)`
.

The sequence of visited tiles is `{(1,1),(2,1),(2,2),(1,2)}`
. In this case, all 2 ⋅ 2 = 4
 tiles are covered, so the answer would be "`YES`". Notice that Prakul wraps around the grid in the last two moves.  
In the third test case: n = 4
, m = 2
, a = 2
, b = 1
. If Prakul starts by moving down, the moves look like:
- Start at `(1,1)`
.
- Jump down a = 2
 tiles: `(1,1)` → `(3,1)`
.
- Jump right b = 1
 tile: `(3,1)` → `(3,2)`
.
- Jump down a = 2
 tiles: `(3,2)` → `(1,2)`
.
- Jump right b = 1
 tile: `(1,2)` → `(1,1)`
.
- ⋯  

In this case, the tiles `{(2,1),(2,2),(4,1),(4,2)}`
 remain unvisited forever. Since not all tiles can be covered, the answer is "`NO`". It can be shown that even if Prakul started by moving right in the first move, the answer would not have changed.

