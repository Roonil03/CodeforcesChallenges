# C. Lavencia Oranges
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
René is leaving Lavencia after two long, cold years, with $n$ Lavencia oranges and $n$ bags for the oranges (initially empty), both numbered from $1$ to $n$. René packs the oranges according to an array $a$: orange $i$ goes into bag $a_i$.

Lavencia has a regulatory process for exports. Jacqueline carries it out: she chooses an integer $s$ and confiscates all bags holding exactly $s$ oranges, picking whichever $s$ takes the most oranges from René. The *cost* of $a$ is the number of oranges Jacqueline confiscates.

More formally, for a given array $a$, let $c_i$ be the number of elements that appear exactly $i$ times in $a$. The *cost* of $a$ is:

$$\max_{0 \le i \le n} i \cdot c_i$$

Given $n$, find an array $a$ that minimizes the cost.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The only line of each test case contains a single integer $n$ ($1 \le n \le 2 \cdot 10^5$).

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

### Output
For each test case, output two lines.

On the first line, output a single integer — the minimum possible cost.

On the second line, output $n$ integers $a_1, a_2, \dots, a_n$ ($1 \le a_i \le n$) — any array that achieves this minimum.

### Examples:
**Input:**
```
3
4
5
8
```
**Output:**
```
2
1 2 3 3
3
1 1 2 2 2
3
4 2 1 5 1 3 1 3
```
**Explaination:**  
In the first test case, René places $1$, $1$, $2$, and $0$ oranges into the four bags, respectively. If Jacqueline picks size $1$, she confiscates bags $1$ and $2$, and René loses $2$ oranges. If she picks size $2$, she confiscates bag $3$, and he again loses $2$. Therefore, the cost is $2$. It can be proved that this is the minimum cost for this test case.

In the second test case, bag $1$ holds $2$ oranges, bag $2$ holds $3$, and the other bags are empty. Picking size $2$ costs René $2$ oranges, and picking size $3$ costs him $3$, so the cost is $3$. It can be proved that, in this case, $3$ is the minimum cost.