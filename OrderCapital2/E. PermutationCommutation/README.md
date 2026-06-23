# E. Permutation Commutation
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Quack the Duck has a permutation<sup>∗</sup>
 a
 of length n
 and an incomplete sequence b<sub>1</sub>, b<sub>2</sub>, …, b<sub>n</sub>
.

Each element of b
 is either −1
 or an integer from 1
 to n
. Each integer from 1
 to n
 appears at most once in b
.

Quack hopes to complete b
 into a permutation that commutes with a
. In other words, after replacing every −1
 in b
, the equality a<sub>b<sub>i</sub></sub> = b<sub>a<sub>i</sub></sub>
 should hold for every 1 ≤ i ≤ n
.

Ja the Ghost wants to help Quack. Among all possible ways to complete b
, he wants to find the lexicographically smallest<sup>†</sup>
 one.

Determine whether such a completion exists. If it exists, output the lexicographically smallest valid permutation b
. Otherwise, report that it is impossible.

<sub><sup>∗</sup>
A permutation of length n
 is an array consisting of n
 distinct integers from 1
 to n
 in arbitrary order. For example, `[2,3,1,5,4]`
 is a permutation, but `[1,2,2]`
 is not a permutation (2
 appears twice in the array), and `[1,3,4]`
 is also not a permutation (n = 3
 but there is 4
 in the array).</sub>

<sub><sup>†</sup>
An array p
 is lexicographically smaller than an array q
 of the same size if and only if the following holds:
- p ≠ q
, and in the first position where p
 and q
 differ, the array p
 has a smaller element than the corresponding element in q
.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains an integer n
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
) — the length of the permutation.

The second line of each test case contains n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ n
) — the permutation a
.

The third line of each test case contains n
 integers b<sub>1</sub>, b<sub>2</sub>, …, b<sub>n</sub>
 (b<sub>i</sub> = −1
 or 1 ≤ b<sub>i</sub> ≤ n
) — the incomplete sequence b
.

It is guaranteed that each integer from 1
 to n
 appears at most once in b
.

It is guaranteed that the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, print "YES" if the answer exists, and "NO" otherwise.

You can output the answer in any case (upper or lower). For example, the strings "yEs", "yes", "Yes", and "YES" will be recognized as positive responses.

If the answer exists, in the next line output n
 integers p<sub>1</sub>, p<sub>2</sub>, …, p<sub>n</sub>
 — the lexicographically smallest valid sequence after replacing every −1
 in b
.

The sequence p
 must be a permutation, that is, each integer from 1
 to n
 must appear exactly once in p
. Also, it must satisfy a<sub>p<sub>i</sub></sub> = p<sub>a<sub>i</sub></sub>
 for every 1 ≤ i ≤ n
.

### Example:
**Input:**
```
12
3
2 3 1
-1 -1 -1
4
2 1 4 3
-1 -1 4 -1
4
2 1 4 3
3 1 -1 -1
4
2 1 4 3
1 -1 -1 2
5
2 3 1 5 4
2 -1 -1 -1 -1
5
2 3 1 5 4
4 -1 -1 -1 -1
6
2 3 1 5 6 4
4 -1 -1 -1 -1 -1
6
2 1 4 3 6 5
-1 3 -1 -1 -1 -1
6
3 5 6 2 1 4
-1 -1 -1 3 6 -1
7
2 3 1 5 4 6 7
-1 -1 -1 -1 -1 7 -1
8
2 3 4 1 6 7 8 5
5 7 -1 -1 -1 -1 -1 -1
8
2 3 4 1 6 7 8 5
5 -1 -1 -1 -1 -1 -1 -1
```
**Output:**
```
YES
1 2 3
YES
1 2 4 3
NO
NO
YES
2 3 1 4 5
NO
YES
4 5 6 1 2 3
YES
4 3 1 2 5 6
NO
YES
1 2 3 4 5 7 6
NO
YES
5 6 7 8 1 2 3 4
```
**Explaination:**  
In the first test case, b = `[1,2,3]`
 commutes with any permutation a
. Since all elements of b
 are unknown, this is also the lexicographically smallest possible valid permutation.

In the second test case, a = `[2,1,4,3]`
 and b<sub>3</sub> = 4
. Since a<sub>3</sub> = 4
, the condition for i = 3
 gives a<sub>b<sub>3</sub></sub> = b<sub>a<sub>3</sub></sub>
, so a<sub>4</sub> = b<sub>4</sub>
, hence b<sub>4</sub> = 3
. The remaining values are 1
 and 2
, and the lexicographically smallest valid choice is b<sub>1</sub> = 1
, b<sub>2</sub> = 2
. Thus the answer is `[1,2,4,3]`
.

In the third test case, a = `[2,1,4,3]`
, b<sub>1</sub> = 3
, and b<sub>2</sub> = 1
. For i = 1
, the condition requires a<sub>b<sub>1</sub></sub> = b<sub>a<sub>1</sub></sub>
. However, a<sub>b<sub>1</sub></sub> = a<sub>3</sub> = 4
, while b<sub>a<sub>1</sub></sub> = b<sub>2</sub> = 1
. Since 4 ≠ 1
, no valid completion exists.

