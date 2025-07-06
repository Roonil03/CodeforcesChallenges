# B. Tournament
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given an array of integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
. A tournament is held with n
 players. Player i
 has strength a<sub>i</sub>
.

While more than k
 players remain,

- Two remaining players are chosen at random;
-Then the chosen player with the lower strength is eliminated. If the chosen players have the same strength, one is eliminated at random.

Given integers j
 and k
 (1≤j,k≤n
), determine if there is any way for player j
 to be one of the last k
 remaining players.


### Input
The first line contains an integer t
 (1≤t≤104
)  — the number of test cases.

The first line of each test case contains three integers n
, j
, and k
 (2≤n≤2 ⋅ 10<sup>5</sup>
, 1≤j,k≤n
).

The second line of each test case contains n
 integers, a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1≤a<sub>i</sub>≤n
).

It is guaranteed that the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output on a single line "`YES`" if player j
 can be one of the last k
 remaining players, and "NO" otherwise.

You can output the answer in any case (upper or lower). For example, the strings "`yEs`", "`yes`", "`Yes`", and "`YES`" will be recognized as positive responses.

### Examples:
**Input:**
```
3
5 2 3
3 2 4 4 1
5 4 1
5 3 4 5 2
6 1 1
1 2 3 4 5 6
```
**Output:**
```
YES
YES
NO
```
**Explaination:**  
In the first sample, suppose that players 2
 and 5
 are chosen. Then player 2
 defeats player 5
. Now, the remaining player strengths are
```
3	2	4	4
```
Next, suppose that players 3
 and 4
 are chosen. Then player 3
 might defeat player 4
. Now, the remaining player strengths are
```
3	2	4
```
Player 2
 is one of the last three players remaining.  
In the third sample, it can be shown that there is no way for player 1
 to be the last player remaining.


