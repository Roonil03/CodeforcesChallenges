# D. Precision Alignment
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
K1o0n volunteered to work as a teaching assistant in the physics department, and he was assigned $n$ training laboratories. In the $i$-th laboratory, there are three instruments; their readings are $a_i$, $b_i$, and $c_i$. The instruments are distinguishable, and the order in this triple is fixed: the first ($a_i$), the second ($b_i$), the third ($c_i$).

One action is to choose the index of a laboratory $i$ and perform **exactly one** of the three adjustments:

* $c_i \mathrel{+}= \operatorname{sgn}(a_i - b_i)$
* $b_i \mathrel{+}= \operatorname{sgn}(a_i - c_i)$
* $a_i \mathrel{+}= \operatorname{sgn}(b_i - c_i)$

Here $\operatorname{sgn}(x)$ equals $1$ if $x > 0$, equals $-1$ if $x < 0$, and equals $0$ if $x = 0$.

In other words, an adjustment shifts the reading of one instrument by one unit, and the direction of the shift is determined by the difference of the readings of the other two, taken in the order they are written: the earlier one minus the later one.

The actions are performed one after another, and each adjustment looks at the readings that have been obtained by that moment. The same laboratory may be adjusted any number of times, and some may be left untouched altogether; in total, K1o0n is ready to make **at most $k$** actions across all laboratories.

The work will be evaluated by the worst laboratory:

$$S = \min_{1 \le i \le n} (a_i + b_i + c_i),$$

where the readings are taken after all actions. K1o0n wants to impress his colleagues, so he asks you to find the maximum $S$ he can achieve.

### Input
The first line contains an integer $t$ ($1 \le t \le 10^4$) — the number of testcases.

The first line of each testcase contains two integers $n$ and $k$ ($1 \le n \le 2 \cdot 10^5$, $0 \le k \le 10^{18}$) — the number of laboratories and the maximum number of actions.

The next $n$ lines of each testcase contain three integers $a_i$, $b_i$, and $c_i$ ($-10^9 \le a_i, b_i, c_i \le 10^9$).

It is guaranteed that the sum of $n$ over all testcases does not exceed $2 \cdot 10^5$.

### Output
For each testcase, output one integer — the maximum achievable value of $S$.

Note that the answer may not fit into a 32-bit integer type. Use a 64-bit integer type (`long long` in C++, `long` in Java).

### Examples:
**Input:**
```
5
3 4
-3 -1 -2
1 3 5
2 4 6
3 4
6 4 2
6 7 0
7 6 0
3 1000000000000000000
4 4 4
1 5 9
-1000000000 0 1000000000
2 4
1 2 3
0 0 5
4 3
2 1 0
4 4 5
1 5 9
7 -2 1
```
**Output:**
```
-2
14
12
6
6
```
**Explaination:**  
In the first testcase, to obtain a minimum sum of all triples equal to $-2$, you need to add $\operatorname{sgn}(b_1 - c_1) = 1$ to $a_1$ four times. After that, the triplet $(a_1, b_1, c_1)$ will become $(1, -1, -2)$, and the sum of the remaining triplets was originally greater than $-2$. It can be shown that it is impossible to obtain a sum of $-1$ or higher with that many operations.

In the second testcase, to obtain a minimum sum of all triples equal to $14$, it is necessary to add $\operatorname{sgn}(b_1 - c_1) = 1$ to $a_1$ $2$ times, add $\operatorname{sgn}(a_2 - c_2) = 1$ to $b_2$ once, and add $\operatorname{sgn}(a_3 - b_3) = 1$ to $c_3$.

In the fourth testcase, it seems that it's impossible to improve the value of the second lab; however, it's not. We can make $b_2 \mathrel{+}= \operatorname{sgn}(a_2 - c_2)$, and we are out of the "stuck" position.