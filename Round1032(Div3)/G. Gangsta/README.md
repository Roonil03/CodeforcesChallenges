# G. Gangsta
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given a binary string s<sub>1</sub>s<sub>2</sub>…s<sub>n
</sub> of length n
. A string s
 is called binary if it consists only of zeros and ones.

For a string p
, we define the function f(p)
 as the maximum number of occurrences of any character in the string p
. For example, f(00110)=3
, f(01)=1
.

You need to find the sum f(s<sub>l</sub>s<sub>l+1</sub>…s<sub>r</sub>)
 for all pairs 1 ≤ l ≤ r ≤ n
.

### Input
Each test consists of multiple test cases. The first line contains a single integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases. Then follows their descriptions.

The first line of each test case contains a single integer n
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
) — the length of the binary string.

The second line of each test case contains a string of length n
, consisting of 0
s and 1
s — the binary string s
.

It is guaranteed that the sum of n
 across all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output the sum f(s<sub>l</sub>s<sub>l+1</sub>…s<sub>r</sub>)
 for all pairs 1 ≤ l ≤ r ≤ n
.

### Examples:
**Input:**
```
6
1
0
2
01
4
0110
6
110001
8
10011100
11
01011011100
```
**Output:**
```
1
3
14
40
78
190
```
**Explaination:**  
In the first test case, the string s
 has one substring, and the value f(0)=1
.

In the second test case, all substrings of the string s
 are 0
, 01
, 1
. And the answer is 1+1+1=3
, respectively.

In the third test case, all substrings of the string s
 are 0
, 01
, 011
, 0110
, 1
, 11
, 110
, 1
, 10
, 0
. And the answer is 1+1+2+2+1+2+2+1+1+1=14
, respectively.