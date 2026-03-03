# E. Rigged Bracket Sequence
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
A regular bracket sequence is a sequence consisting of '`(`' and '`)`', which can be turned into a valid math expression by inserting `1`
 and `+`
 any number of times into the sequence. For example, the sequence "`()(()())`" is a regular bracket sequence, while "`())(()`" or "`(()`" are not regular bracket sequences.

You are given a regular bracket sequence S
.

Let us consider shifting a subsequence<sup>∗</sup>
 to the right. Formally, when a subsequence S<sub>i<sub>1</sub></sub>S<sub>i<sub>2</sub></sub>…S<sub>i<sub>k</sub></sub>
 is shifted to the right, the characters on the chosen indices simultaneously get reassigned as follows:
- S<sub>i<sub>1</sub></sub> ← S<sub>i<sub>k</sub></sub>
;
- S<sub>i<sub>2</sub></sub> ← S<sub>i<sub>1</sub></sub>
;
- S<sub>i<sub>3</sub></sub> ← S<sub>i<sub>2</sub></sub>
;
- …
- S<sub>i<sub>k</sub></sub> ← S<sub>i<sub>k−1</sub></sub>
.

In other words, the element of the j
-th chosen index gets reassigned to the `((j − 2) mod k + 1)`
-th chosen character.

For example, when S
 is "`()(()())`", shifting the subsequence S<sub>2</sub>S<sub>4</sub>
 changes S
 to "`((())())`". On the other hand, shifting S<sub>2</sub>S<sub>3</sub>S<sub>5</sub>
 changes S
 to "`())((())`".

Please count how many non-empty subsequences make S
 remain regular when shifted to the right. As the answer may be huge, you are only asked to output the answer modulo 998244353
.

<sub><sup>∗</sup>
A sequence a
 is a subsequence of a sequence b
 if a
 can be obtained from b
 by the deletion of several (possibly, zero or all) element from arbitrary positions. Two subsequences are considered different if the sets of positions of the deleted elements are different.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains a single integer n
 (2 ≤ n ≤ 300000
, n
 is even).

The second line of each test case contains a regular bracket sequence S
 of length n
 given as a string without spaces.

It is guaranteed that the sum of n
 over all test cases does not exceed 300000
.

### Output
For each test case, output the answer modulo 998244353
 on a separate line.

### Examples:
**Input:**
```
4
2
()
4
()()
6
(()())
10
()((())())
```
**Output:**
```
2
8
28
312
```
**Note:**  
For the second test case, the 8
 non-empty subsequences that make S
 remain regular when shifted to the right are as follows:
- S<sub>1</sub>
: changes S
 to "`()()`";
- S<sub>2</sub>
: changes S
 to "`()()`";
- S<sub>3</sub>
: changes S
 to "`()()`";
- S<sub>4</sub>
: changes S
 to "`()()`";
- S<sub>1</sub>S<sub>3</sub>
: changes S
 to "`()()`";
- S<sub>2</sub>S<sub>3</sub>
: changes S
 to "`(())`";
- S<sub>2</sub>S<sub>4</sub>
: changes S
 to "`()()`";
- S<sub>1</sub>S<sub>2</sub>S<sub>3</sub>
: changes S
 to "`(())`".
