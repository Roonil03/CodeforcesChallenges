# D. Good Schedule
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
Alice and Bob decided to watch a TV series consisting of n
 episodes, numbered from 1
 to n
. The series will be shown on television over the next n
 days. Unfortunately, they live in different cities, and the episode schedules may differ. On the i
-th day, the a<sub>i</sub>
-th episode will be shown in Alice's city, and the b<sub>i</sub>
-th episode in Bob's city.

They plan to select a segment of days `[L,R]`
 (1 ≤ L ≤ R ≤ n
) to watch the series. Initially, neither of them has seen any episodes. Each day i
 in this segment, the following happens:
- if Alice has already watched episodes 1, 2, …, a<sub>i−1</sub>
, but not a<sub>i</sub>
, then she watches a<sub>i</sub>
 on day i
; otherwise, she watches nothing on this day;
- if Bob has already watched episodes 1, 2, …, b<sub>i−1</sub>
, but not b<sub>i</sub>
, then he watches b<sub>i</sub>
 on day i
; otherwise, he watches nothing on this day.

To avoid spoilers in their conversations, Alice and Bob want to choose a segment `[L,R]`
 such that on every day in this segment, one of the following holds:
- either they both watch the same episode on that day;
- or neither of them watches anything on that day.

Help Alice and Bob calculate the number of suitable segments `[L,R]`
.

### Input
The first line contains a single integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases.

Each test case consists of three lines:
- the first line contains a single integer n
 (1 ≤ n ≤ 5 ⋅ 10<sup>5</sup>
);
- the second line contains n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ n
);
the third line contains n
 integers b<sub>1</sub>, b<sub>2</sub>, …, b<sub>n</sub>
 (1 ≤ b<sub>i</sub> ≤ n
).

Additional constraint on the input: the sum of n
 over all test cases does not exceed 5 ⋅ 10<sup>5</sup>
.

### Output
For each test case, print a single integer — the number of suitable segments [L,R]
.

### Examples:
**Input:**
```
4
3
1 2 1
1 2 2
2
2 1
1 2
5
1 3 2 1 4
1 4 2 3 2
9
1 1 3 1 1 3 2 3 1
1 3 1 1 3 1 2 1 3
```
**Output:**
```
4
0
7
12
```
**Explaination:**  
In the first example, the suitable segments are `[1,1]`
, `[1,2]`
, `[1,3]`
, and `[2,2]`
.

