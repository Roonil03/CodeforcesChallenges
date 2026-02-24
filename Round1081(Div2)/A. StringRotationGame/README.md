# A. String Rotation Game
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
Define a block in a string as a contiguous substring of characters of the same type that cannot be extended either to the left or the right. For example, in the string aabcccdaa, there are five blocks:
- `aa` (1-st to 2-nd characters)
- `b` (3-rd character)
- `ccc` (4-th to 6-th characters)
- `d` (7-th character)
- `aa` (8-th to 9-th characters). 

You are playing a game where you are given a string s of length n. You can cyclically rotate<sup>∗</sup> the string however you want. Your score is then calculated as the number of blocks in the final string. Please find the maximum score possible.

<sub><sup>∗</sup>Formally, choose an index 1 ≤ i ≤ n, and replace the string s<sub>1</sub>s<sub>2</sub>…s<sub>n</sub> with the string s<sub>i+1</sub>s<sub>i+2</sub>…s<sub>n</sub>s<sub>1</sub>s<sub>2</sub>…s<sub>i</sub>. For example, the string abcde can be rotated to string deabc by choosing i=3.</sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 500). The description of the test cases follows.

The first line of each test case contains a single integer n (1 ≤ n ≤ 100).

The second line of each test case contains the string s of length n.

Strings s consist of lowercase Latin characters only.
### Output

For each testcase, output a single integer denoting the maximum score you can achieve.

### Examples:
**Input:**
```
4
4
abcd
4
abbc
4
abba
6
abbccc
```
**Output:**
```
4
4
3
4
```
**Note**  
In the first test case, score of the original string abcd is 4. It can be shown that a score greater than 4 cannot be achieved.

In the second test case, cyclically rotating the string by 2 positions will give us string bcab. Score of this string is 4. It can be shown that a score greater than 4 cannot be achieved.