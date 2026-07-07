# D. Two Digit Strings
## Time limit per test: 3 seconds
## Memory limit per test: 512 megabytes
### Description:
You are given two strings $a$ and $b$, consisting only of digits. In one operation, you can choose two adjacent characters in either of these two strings and replace them with their sum modulo 10. For example, from the string `57246` you can obtain the following strings in one operation:

* `57246` $\to$ `246`;
* `57246` $\to$ `5946`;
* `57246` $\to$ `5766`;
* `57246` $\to$ `5720`;

Note that after such an operation, the length of the string always decreases by exactly 1.

You can perform any number of such operations (even zero). You have to make these strings equal ($|a| = |b|$ and $a_i = b_i$ for all $i$ from 1 to $|a|$). Calculate the maximum possible length of the resulting equal strings.

### Input
The first line contains a single integer $t$ ($1 \le t \le 2 \cdot 10^3$) — the number of test cases.

The first line of each test case contains a string $a$ ($1 \le |a| \le 5 \cdot 10^3$), consisting only of digits.

The second line of each test case contains a string $b$ ($1 \le |b| \le 5 \cdot 10^3$), consisting only of digits.

Additional constraint on the input: $\sum(|a| + |b|)$ across all test cases does not exceed $10^4$.

### Output
For each test case, print one integer — the maximum possible length of the resulting equal strings. If you can't make the strings equal, print -1.

### Examples:
**Input:**
```
3
5147
44441
2194
5602
123450
012345
```
**Output:**
```
2
-1
5
```
**Explaination:**  
In the first example from the statement, the maximum length you can obtain is 2:

* `5147` $\to$ `647` $\to$ `61`;
* `44441` $\to$ `8441` $\to$ `241` $\to$ `61`.

In the second example from the statement, you can't make strings $a$ and $b$ equal, so the answer is -1.

In the third example from the statement, the maximum length you can obtain is 5:

* `123450` $\to$ `12345`;
* `012345` $\to$ `12345`.