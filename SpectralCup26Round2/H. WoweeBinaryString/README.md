# H. Wowee Binary String
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
You find yourself with a string s of length n consisting of only the characters 0, 1 and ?. In other words, s is an incomplete binary string. You will do the following operations in order:
1. replace all ? in s with either 0 or 1.
2. then repeat the following any number of times (possibly none):
    - select a substring of s with an even number of occurrences of the character 1 and delete it. More formally, select two integers l, r where 1 ≤ l ≤ r ≤ |s| and 1 occurs in s<sub>l</sub>, s<sub>l+1</sub>, …, s<sub>r</sub> an even number of times, then replace s with s<sub>1</sub>, …, s<sub>l−1</sub>, s<sub>r+1</sub>, …,s<sub>|s|</sub> 

Find how many different binary strings can be obtained after performing the operations. As the number could be humungous, find it modulo 998244353.

### Input:
Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 100). The description of the test cases follows.

The first line of each testcase contains an integer n (1 ≤ n ≤ 3000) — the length of the string.

The second line of each test case contains the incomplete binary string s.

It is guaranteed that the sum of n over all test cases does not exceed 3000. 

### Output

For each testcase, output the number of binary strings that can be obtained, modulo 998244353.

### Examples:
**Sample Input:**
```
7
1
?
5
?????
5
?1001
8
10010001
7
00??10?
7
?1?0?1?
16
0??000?00??1?1?0
```
**Sample Output:**
```
3
63
14
18
71
95
5399
```
**Explaination:**  
In the first two testcases, any binary string of length no longer than $n$ can be made.

In the third testcase, the following strings can be made: $\epsilon, 0, 1, 01, 11, 001, 011, 101, 111, 0101, 1001, 1101, 01001, 11001$, where $\epsilon$ represents the empty string.

An example of how $1$ can be made is as follows:

* $\textcolor{red}{?}1001 \to 11001$
* $1\textcolor{red}{10}01 \to 01$
* $\textcolor{red}{0}1 \to 1$