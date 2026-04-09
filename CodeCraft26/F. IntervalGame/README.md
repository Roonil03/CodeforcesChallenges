# F. Interval Game
## Time limit per test: 3 seconds
## Memory limit per test: 256 megabytes
### Description:
Alice and Bob are playing a game involving two intervals, `[l`<sub>`1`</sub>`,r`<sub>`1`</sub>`]`
 and `[l`<sub>`2`</sub>`,r`<sub>`2`</sub>`]`
, constrained by integers x<sub>1</sub>
 and x<sub>2</sub>
 respectively. The intervals satisfy 1 ≤ l<sub>1</sub> ≤ r<sub>1</sub> ≤ x<sub>1</sub>
 and 1 ≤ l<sub>2</sub> ≤ r<sub>2</sub> ≤ x<sub>2</sub>
.

The game is played as follows. In one turn, a player must choose one of the two intervals (let's say the i
-th interval `[l`<sub>`i`</sub>`,r`<sub>`i`</sub>`]`
) and perform exactly one of the following operations:
- Replace l<sub>i</sub>
 with a strictly smaller integer l′<sub>i</sub>
 such that 1 ≤ l′<sub>i</sub> < l<sub>i</sub>
.
- Replace r<sub>i</sub>
 with a strictly larger integer r′<sub>i</sub>
 such that r<sub>i</sub> < r′<sub>i</sub> ≤ x<sub>i</sub>
.

The player who cannot make a valid move loses the game.  
The game setup proceeds in two steps:
1. First, Alice chooses the values of l<sub>1</sub>
 and r<sub>1</sub>
 to determine the first interval. It must be a valid interval within `[1,x`<sub>`1`</sub>`]`
.
2. Then, the second interval `[l`<sub>`2`</sub>`,r`<sub>`2`</sub>`]`
 is chosen uniformly at random from all possible valid intervals within `[1,x`<sub>`2`</sub>`]`
.
Once the intervals are chosen, the game begins with Alice moving first.

Assuming both Alice and Bob play optimally, Alice wants to choose the initial interval `[l`<sub>`1`</sub>`,r`<sub>`1`</sub>`]`
 to maximize her probability of winning. Your task is to find such an interval. If there are multiple optimal intervals, you may output any one of them.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The only line of each test case contains two integers x<sub>1</sub>
 and x<sub>2</sub>
 (1 ≤ x<sub>1</sub>, x<sub>2</sub> ≤ 5 ⋅ 10<sup>5</sup>
) — the upper bounds for the two intervals.

It is guaranteed that the sum of x<sub>1</sub>
 over all test cases does not exceed 5 ⋅ 10<sup>5</sup>
.

It is guaranteed that the sum of x<sub>2</sub>
 over all test cases does not exceed 5 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output two integers l<sub>1</sub>
 and r<sub>1</sub>
 (1 ≤ l<sub>1</sub> ≤ r<sub>1</sub> ≤ x<sub>1</sub>
) — the optimal interval chosen by Alice.

If there are multiple optimal intervals, you may output any one of them.

### Examples:
**Input:**
```
6
1 1
1 10
2 1
2 2
20 64578
185367 133524
```
**Output:**
```
1 1
1 1
2 2
1 2
5 12
37381 52023
```
**Note:**  
For the first test case, Alice is forced to choose `[1,1]`
 as it is the only valid option.

For the third test case, if Alice chooses `[2,2]`
, she is guaranteed to win (probability 1
); she can replace l<sub>1</sub>
 with 1
, after which Bob cannot make a valid move. Note that `[1,1]`
 would also be a valid winning answer here.

For the fourth test case, if Alice chooses `[1,2]`
, she achieves a winning probability of 2/3
. By contrast, choosing `[1,1]`
 or `[2,2]`
 would result in a winning probability of only 1/3
.
