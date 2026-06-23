# D. Fullmetal Bitchemist
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
A binary string is a string consisting only of characters 0
 and 1
. The two characters 0
 and 1
 are called opposite values.

Consider a binary string t
. Let `|t|`
 be the length of t
. When `|t|` ≥ 2
, for every 1 ≤ i < `|t|`
, the characters t<sub>i</sub>
 and t<sub>i + 1</sub>
 are adjacent.

A binary string t
 is called beautiful if it can be reduced to a string of length exactly 1
 by applying the following operation any number of times, possibly zero:
- Choose two equal adjacent characters, remove both of them, and insert one character with the opposite value in their place.

For example, the string `10001`
 can become `1101`
 by replacing the first adjacent pair `00`
 with `1`
. Then it can become `001`
, then `11`
, and finally `0`
. Therefore, `10001`
 is beautiful.

On the other hand, `111`
 is not beautiful. After one operation, it becomes 01
, and then no operation can be applied.

You are given a binary string s
. Compute the number of non-empty beautiful substrings<sup>∗</sup>
 of s
.

<sub><sup>∗</sup>
A string a
 is a substring of a string b
 if a
 can be obtained from b
 by the deletion of several (possibly, zero or all) characters from the beginning and several (possibly, zero or all) characters from the end.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains an integer n
 (1 ≤ n ≤ 10<sup>6</sup>
) — the length of s
.

The second line of each test case contains a binary string s
 of length n
.

It is guaranteed that the sum of n
 over all test cases does not exceed 10<sup>6</sup>
.

### Output
For each test case, output a single integer — the number of beautiful substrings of s.

### Examples:
**Input:**
```
10
1
0
2
01
5
01001
3
001
6
011110
9
010110110
12
010000101001
16
1010011010010110
20
11110101101101001110
30
000101100011111001111100000010
```
**Output:**
```
1
2
10
5
15
30
47
81
139
316
```
**Explaination:**  
In the first test case, the only non-empty substring is 0
, which is already a binary string of length 1
. Therefore, it is beautiful.

In the second test case, the beautiful substrings are 0
 and 1
. The substring 01
 is not beautiful, because its two characters are not equal, so no operation can be applied.

In the third test case, the beautiful substrings are:

- `s[1,1]` = 0
;
- `s[2,2]` = 1
;
- `s[3,3]` = 0
;
- `s[4,4]` = 0
;
- `s[5,5]` = 1
;
- `s[3,4]` = 00
, which can become 1
;
- `s[2,4]` = 100
, which can become 11
, then 0
;
- `s[3,5]` = 001
, which can become 11
, then 0
;
- `s[1,4]` = 0100
, which can become 011
, then 00
, then 1
;
- `s[1,5]` = 01001
, which can become 0111
, then 001
, then 11
, and finally 0

Thus, there are 10
 beautiful substrings.