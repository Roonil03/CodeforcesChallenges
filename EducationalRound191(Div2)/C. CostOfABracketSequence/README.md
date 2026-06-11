# C. Cost of a Bracket Sequence
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
Let the cost of an arbitrary bracket string be the length of it's longest subsequence<sup>∗</sup>
 that is a regular bracket sequence<sup>†</sup>
.

You are given a bracket string s
 and an integer k
. Your task is to remove at most k
 characters from the string s
 so that the cost of the resulting string is minimized.

<sub><sup>
∗</sup>
A sequence a
 is a subsequence of a sequence b
 if a
 can be obtained from b
 by the deletion of several (possibly, zero or all) element from arbitrary positions.

<sup>†</sup>
A bracket sequence is called regular if it is possible to obtain a correct arithmetic expression by inserting the characters +
 and 1
 into this sequence. For example, the sequences "`(())()`
", "`()`
", and "`(()(()))`
" are regular, while "`)(`
", "`(()`
", and "`(()))(`
" — are not.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>3</sup>
). The description of the test cases follows.

The first line of each test case contains two integers n
 and k
 (1 ≤ n ≤ 5000
; 0 ≤ k ≤ n
) — the length of the string s
 and the maximum number of deletions.

The second line of each test case contains a string s
 of length n
 consisting of the characters "`(`
" and/or "`)`
".

Additional input constraints:
- the sum of n
 over all test cases does not exceed 5000
.

### Output
For each test case, output a binary string of length n
. The i
-th character should be equal to "`1`" if the corresponding character of the string s
 is removed, and "`0`" otherwise.

The number of ones in the string must not exceed k
. The cost of the string obtained after removing the marked characters must be as small as possible.

If there are several answers, output any of them.

### Examples:
**Input:**
```
10
2 1
)(
2 0
()
4 1
(())
4 1
())(
5 1
((())
6 2
()()()
6 2
(()())
6 2
())(()
7 3
(()((()
10 3
(()())())(
```
**Output:**
```
00
00
1000
1000
00010
101000
001001
100001
1100001
0101001000
```
**Explaination:**  
In the first test case, the cost of the string is already 0
, so it is possible not to delete anything.

In the third test case, it is impossible to obtain a string of cost 0
 after one deletion, but it is possible to obtain a string of cost 2
 by deleting any character.