# E. Sum of Digits (and Again)
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
For a positive (strictly greater than 0
) integer x
, the string S(x)
 is formed through the following process:

initially, this string is empty;
then, the decimal representation of the number x
 without leading zeros is appended to it on the right;
after that, if x≤9
, the process ends. Otherwise, x
 is replaced by the sum of the digits of x
, and the process returns to step 2
;
For example:
- S(75)
 is 75123;
- S(30)
 is 303;
- S(9)
 is 9.

You are given a string s
 consisting of digits. Your task is to rearrange the characters in this string so that it forms the string S(x)
 for some number x
. Removing characters and/or adding new characters is not allowed. If the given string s
 is already the string S(x)
 for some number x
, you may leave it unchanged.

### Input
The first line contains one integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases.

Each test case consists of one string containing s
 (1 ≤ |s| ≤ 10<sup>5</sup>
) — a sequence of digits.

Additional constraints on the input:
- the sum of lengths of s
 over all test cases does not exceed 10<sup>5</sup>
;
- it is possible to rearrange the digits in s
 to obtain S(x)
 for some positive integer x
.
### Output
For each test case, output one string — s
 after rearranging the characters. If there are multiple valid answers, output any of them.

### Examples:
**Input:**
```
5
12735
1
011
99999299999999299959999999999999
4621467
```
**Output:**
```
75123
1
101
99999999999999999999999999992529
6442167
```