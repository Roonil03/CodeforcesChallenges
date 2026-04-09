# H. Closer
## Time limit per test: 3 seconds
## Memory limit per test: 512 megabytes
### Description:
At your networking event, there are n
 deals waiting to happen.

Each deal needs exactly two people, and every person is waiting for exactly one deal. So there are 2n
 people in total, and each deal ID 1…n
 appears on exactly two people's badges.

For corporate reasons, the organisers place each person on a unique vertex in a tree with 2n
 vertices. Vertex v
 hosts a person with badge a<sub>v</sub>
.

Deal i
 is sealed if, when the event begins, the two people with badge i
 are on adjacent vertices. Each sealed deal contributes wi
 to your earnings.

Before the event starts, you are allowed one reshuffle:
1. Choose any (potentially empty) set of edges such that no two chosen edges share a vertex (i.e., the edges form a matching).
2. For every chosen edge `(u,v)`
, swap the people at vertices u
 and v
.

After reshuffling, what is the maximum you can earn from this event?

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains a single integer n
 (1 ≤ n ≤ 10<sup>5</sup>
).

The next line contains n
 integers w<sub>1</sub>, w<sub>2</sub>, …, w<sub>n</sub>
 (0 ≤ w<sub>i</sub> ≤ 10<sup>5</sup>
), where w<sub>i</sub>
 denotes the profit from deal i
.

The next line contains 2n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>2n</sub>
. It is guaranteed that each value 1, 2, …, n
 appears exactly twice in a
.

The next 2n − 1
 lines each contain two integers u
 and v
 (1 ≤ u, v, ≤ 2n , u ≠ v
) — denoting that an edge connects nodes u
 and v
.

It is guaranteed that the sum of n
 over all test cases does not exceed 10<sup>5</sup>
.

### Output
For each test case, output a single integer — the maximum total earning after a reshuffle.

### Examples:
**Input:**
```
3
2
10 3
1 2 2 1
1 2
2 3
3 4
3
8 7 1
3 1 2 3 1 2
1 2
1 3
1 4
1 5
1 6
3
4 9 7
1 2 3 1 2 3
1 2
2 3
3 4
4 5
5 6
```
**Output:**
```
10
8
11
```
**Note:**  
Test 1. Swap edges `(1,2)`
 and `(3,4)`
 (they don't share vertices). Then the two 1s become adjacent on edge `(2,3)`
, earning w<sub>1</sub> = 10
, which beats keeping the existing `2—2` adjacency worth 3.

Test 2. The tree is a star. Swapping `(1,2)`
 moves badge 1 to the center, making it adjacent to the other 1, earning w<sub>1</sub> = 8
.

Test 3. Swap `(1,2)`, `(3,4)`, `(5,6)`
 to seal two deals at once: edge `(2,3)`
 becomes `1—1` (gain 4) and edge `(4,5)`
 becomes `3—3` (gain 7), total 11.