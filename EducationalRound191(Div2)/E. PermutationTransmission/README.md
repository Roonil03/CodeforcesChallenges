# E1. Permutation Transmission (Easy Version)
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
**This is the easy version of the problem. In this version, the upper bound on n and the sum of n over all test cases is 2000; furthermore, the maximum number of test cases is 200.**

There was a permutation p
 of size n<sup>
∗</sup>
.

It was sent through a communication channel as follows: first, all 0
-th bits of each number p<sub>i</sub>
 in the permutation were sent as one string of n
 characters 0 and/or 1; then the 1
-st bits were sent in the same format, and so on, up to the most significant bit of the number n
.

For example, for p = `[3,1,2,4]`
, the following three strings were sent:
1. "`1100`";
2. "`1010`";
3. "`0001`".

You received all these strings, but the order in which each string corresponded to a bit was lost, that is, the strings arrived in arbitrary order. In the example above, the strings could have arrived in the order "`1010`", "`0001`", and "`1100`".

Your task is to determine how many possible original permutations p
 could have been transmitted. It is possible that the data was corrupted during transmission, and then there is no valid original permutation p
.

<sub><sup>∗</sup>
A permutation of size n
 is an array of size n
 where each integer from 1
 to n
 appears exactly once.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 200
). The description of the test cases follows.

The first line of each test case contains one integer n
 (1 ≤ n ≤ 2000
).

The next ⌈log<sub>2</sub>(n + 1)⌉
 lines of each test case contain the transmitted data in the received order. Each line consists of n
 characters 0
 and/or 1
.

Additional input constraints:
- the sum of n
 over all test cases does not exceed 2000
.

### Output
For each test case, print one integer — the answer to the problem.

### Examples:
**Input:**
```
6
1
1
4
1010
0001
0110
7
0101011
1100101
0110110
5
10110
01100
11001
6
001011
111000
000000
7
1100001
1100001
1100011
```
**Output:**
```
1
2
6
0
0
0
```
**Explaination:**  
In the first example, there could be only one permutation, p = `[1]`
.

In the second example, there are 2
 possible variants: `[1,2,3,4]`
 and `[2,1,3,4]`
.

In the fourth example, there are no suitable original permutations.

# E2. Permutation Transmission (Difficult Version)
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
**This is the difficult version of the problem. In this version, the upper bound on n  and the sum of n over all test cases is 2 ⋅ 10<sup>5</sup>; furthermore, the maximum number of test cases is 10<sup>4</sup>.**

There was a permutation p
 of size n
<sup>∗</sup>
.

It was sent through a communication channel as follows: first, all 0
-th bits of each number p<sub>i</sub>
 in the permutation were sent as one string of n
 characters 0 and/or 1; then the 1
-st bits were sent in the same format, and so on, up to the most significant bit of the number n
.

For example, for p = `[3,1,2,4]`
, the following three strings were sent:
1. "`1100`";
2. "`1010`";
3. "`0001`".

You received all these strings, but the order in which each string corresponded to a bit was lost, that is, the strings arrived in arbitrary order. In the example above, the strings could have arrived in the order "`1010`", "`0001`", and "`1100`".

Your task is to determine how many possible original permutations p
 could have been transmitted. It is possible that the data was corrupted during transmission, and then there is no valid original permutation p
.

<sub><sup∗</sup>
A permutation of size n
 is an array of size n
 where each integer from 1
 to n
 appears exactly once.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains one integer n
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
).

The next ⌈log<sub>2</sub>(n + 1)⌉
 lines of each test case contain the transmitted data in the received order. Each line consists of n
 characters 0
 and/or 1
.

Additional input constraints:
- the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, print one integer — the answer to the problem.

### Examples
**Input:**
```
6
1
1
4
1010
0001
0110
7
0101011
1100101
0110110
5
10110
01100
11001
6
001011
111000
000000
7
1100001
1100001
1100011
```
**Output:**
```
1
2
6
0
0
0
```
**Explaination:**  
In the first example, there could be only one permutation, p = `[1]`
.

In the second example, there are 2
 possible variants: `[1,2,3,4]`
 and `[2,1,3,4]`
.

In the fourth example, there are no suitable original permutations.

