# D. Robot Championship
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Fran and Sofia have become huge fans of robot aura fights, so they have decided to create their own robot aura tournament. Since Sofia is a great expert on robots, Fran will need your help to win the championship.

A match of the tournament takes place on an $n \times n$ board. The rows are numbered from **bottom to top** and the columns from left to right, so $(1, 1)$ is the bottom-left corner and $(n, n)$ is the top-right corner. The opponent's robot is hidden at some unknown cell $(i, j)$ and moves once every second, but you do not see it.

At the beginning of each second $s$, before the robot moves, you must choose an integer $v$ with $2 \le v \le 2n$ and shoot a laser along the diagonal

$$\{(r, c) \mid r + c = v\}.$$

The laser touches the robot if the robot is currently on a cell of that diagonal, in which case you win immediately. Otherwise, after your shot, suppose the robot is currently at $(r, c)$, the robot moves to one of the following cells, provided the move stays inside the board:

* up: $(r + 1, c)$;
* right: $(r, c + 1)$;
* down-left: $(r - 1, c - 1)$.

<img src="https://espresso.codeforces.com/76c8fb5da81a03caa0f7e955b4dd790574238278.png"><br>

The robot is at cell $(3, 2)$ (its diagonal is $r + c = 5$). The green laser is the diagonal $v = 7$, which does not touch the robot.
The robot is adversarial: it may start on any cell, and it may choose any valid sequence of moves. You have to output a fixed sequence of at most $3n$ diagonals $v_1, v_2, \dots, v_k$ such that, no matter where the robot starts and no matter how it moves, at least one of the lasers is guaranteed to touch it at some second during the tournament.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 44$). The description of the test cases follows.

The only line of each test case contains a single integer $n$ ($2 \le n \le 100$) — the size of the board.

It is guaranteed that the sum of $n^3$ over all test cases does not exceed $10^6$.

### Output
For each testcase, print two lines. The first line contains an integer $k$ ($1 \le k \le 3n$) — the number of lasers.

The second line contains $k$ integers $v_1, v_2, \dots, v_k$ ($2 \le v_i \le 2n$) — the diagonals in the order they are shot.

You may print any valid sequence.

### Examples:
**Input:**
```
2
2
3
```
**Output:**
```
3
3 3 3
3
4 4 4
```
**Explaination:**  
For $n = 2$, the sequence $3, 3, 3$ works. The first laser $3$ removes the diagonal $r + c = 3$ (cells $(1, 2)$ and $(2, 1)$). Cell $(1, 1)$ must move onto that same diagonal on the next turn, and cell $(2, 2)$ is forced to follow the path $(2, 2) \to (1, 1) \to$ the same diagonal. The second and third lasers catch both.

For $n = 3$, the sequence $4, 4, 4$ works. A robot on the diagonal $r + c = 4$ is hit by the first laser. From the diagonal $r + c = 3$ or from cell $(3, 3)$, every move leads onto $r + c = 4$, so the second laser hits it. From cell $(1, 1)$ or from the diagonal $r + c = 5$, every move leads onto $r + c = 3$ or to cell $(3, 3)$, so the third laser hits it.