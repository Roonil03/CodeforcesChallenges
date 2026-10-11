# A. Robot Odd Moves
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
You are given a point $(a, b)$ of the integer grid.

A robot is standing at the point $(0, 0)$. The robot may perform the following two-step operation any number of times (possibly zero), where $(x, y)$ denotes the point the robot is standing at:

1. Move one unit up or one unit down: $(x, y) \to (x, y + 1)$ or $(x, y) \to (x, y - 1)$.
2. Then move one unit to the right: $(x, y) \to (x + 1, y)$.

The robot *visits* a point if at some moment it is standing at this point: before the first operation, between the two moves of an operation, or after an operation.

Find the minimum number of operations after which the robot has visited $(a, b)$, or determine that it is impossible.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 100$). The description of the test cases follows.

The only line of each test case contains two integers $a$ and $b$ ($0 \le a, b \le 1000$) — the coordinates of the given point.

### Output
For each test case, output a single integer — the minimum number of operations after which the robot has visited $(a, b)$, or $-1$ if it is impossible.

### Examples:
**Input:**
```
9
0 0
0 1
0 2
1 0
2 1
2 3
5 6
1000 999
1000 1000
```
**Output:**
```
0
1
-1
2
3
3
6
1001
1000
```
**Explaination:**  
In the first example, the robot is standing at $(0, 0)$ before the first operation, so it has already visited $(0, 0)$ after $0$ operations.

In the second example, the robot moves up to $(0, 1)$ and then right to $(1, 1)$ during its first operation. Thus, it visits $(0, 1)$ after $1$ operation.

In the third example, it can be proven that the robot cannot visit $(0, 2)$.

In the fourth example, the robot can first move up and right, and then down and right. During the second operation, it visits $(1, 0)$. It cannot visit this point during the first operation, so the answer is $2$.