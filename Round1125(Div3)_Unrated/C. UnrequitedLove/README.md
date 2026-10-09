# C. Unrequited Love
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Descriptive:
K1o0n reached the stage where, for complete happiness, he only lacked a musical instrument, and bought a synthesizer second-hand. The instrument has $n$ keys, giving $a_1, a_2, \dots, a_n$ units of audience love.

Playing one note at a time is boring, so K1o0n learned a single technique — a *triad*: three keys with one key between each pair, pressed simultaneously. A triad starting at $x$ ($1 \le x \le n - 4$) is the keys $x$, $x + 2$, and $x + 4$, which brings $a_x + a_{x+2} - a_{x+4}$ units of audience love.

The keys are old, and each can withstand exactly one press. Therefore, K1o0n wants to choose exactly two different triads that **do not** share any keys. In addition, he wants both triads to bring the **same** amount of audience love.

Count the number of ways to choose two such triads.

### Input
The first line contains an integer $t$ ($1 \le t \le 10^4$) — the number of testcases.

The first line of each testcase contains an integer $n$ ($6 \le n \le 2 \cdot 10^5$) — the number of keys.

The second line of each testcase contains $n$ integers $a_i$ ($-10^4 \le a_i \le 10^4$) — the units of audience love.

It is guaranteed that the sum of $n$ over all testcases does not exceed $2 \cdot 10^5$.

### Output
For each testcase, output one integer — the number of unordered pairs of disjoint triads with equal audience love values.

Note that the answer may not fit into a 32-bit type. Use a 64-bit type (`long long` in C++, `long` in Java).

### Examples:
**Input:**
```
3
6
5 1 4 2 2 -4
7
0 0 0 0 0 0 0
12
-4 3 0 -4 -4 -1 -4 1 -3 -2 -2 -1
```
**Output:**
```
1
2
5
```
**Explaination:**  
In the first testcase, such a pair is unique: $x = 1$, $y = 2$: $5 + 4 - 2 = 1 + 2 - (-4)$.

In the second testcase, there are two such pairs: $x = 1$, $y = 2$ and $x = 3$, $y = 2$, and each triad brings $0$ units of audience love.