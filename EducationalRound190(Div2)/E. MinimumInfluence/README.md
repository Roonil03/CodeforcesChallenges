# E. Minimum Influence
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
Imagine that you are the owner of a news website and want to study how some selected news items affect your users.

You have $n$ news items, and for each of them, you have already determined two parameters: how much it touches politics $p_i$ and how much it touches culture $c_i$.

You also have $m$ users whose reaction to the news you want to study. For each person, you have already determined three parameters: tolerance to political news $tp_j$, tolerance to cultural news $tc_j$, and the "zone of influence" $d_j$.

The influence of politics $I_p(i, j)$ and culture $I_c(i, j)$ in news item $i$ on user $j$ can be calculated by the following formulas:

$$
I_p(i, j) = \begin{cases} 
0 & \text{if } p_i < tp_j \\ 
p_i & \text{if } tp_j \le p_i < tp_j + d_j \\ 
tp_j + d_j & \text{if } p_i \ge tp_j + d_j 
\end{cases}, \quad 
I_c(i, j) = \begin{cases} 
0 & \text{if } c_i < tc_j \\ 
c_i & \text{if } tc_j \le c_i < tc_j + d_j \\ 
tc_j + d_j & \text{if } c_i \ge tc_j + d_j 
\end{cases}
$$

In other words, while the amount of politics $p_i$ is less than the tolerance level $tp_j$, it does not affect the user. Otherwise, the topic starts to irritate the user, but not more than up to $tp_j + d_j$. The same goes for culture.

The total influence of news item $i$ on user $j$ is $I(i, j) = I_p(i, j) + I_c(i, j)$.

For each user $j$, determine the **minimum** influence $I(i, j)$ among all news items $i$.

### Input
The first line contains one integer n
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
) — the number of news items.

The second line contains n
 integers p<sub>1</sub>, p<sub>2</sub>, …, p<sub>n</sub>
 (0 ≤ p<sub>i</sub> ≤ 10<sup>6</sup>
) — the amount of political content of each news item.

The third line contains n
 integers c<sub>1</sub>, c<sub>2</sub>, …, c<sub>n</sub>
 (0 ≤ c<sub>i</sub> ≤ 10<sup>6</sup>
) — the amount of cultural content of each news item.

The fourth line contains one integer m
 (1 ≤ m ≤ 4 ⋅ 10<sup>5</sup>
) — the number of users.

The fifth line contains m
 integers tp<sub>1</sub>, tp<sub>2</sub>, …, tp<sub>m</sub>
 (0 ≤ tp<sub>j</sub> ≤ 10<sup>6</sup>
) — the political tolerance of each user.

The sixth line contains m
 integers tc<sub>1</sub>, tc<sub>2</sub>, …, tc<sub>m</sub>
 (0 ≤ tc<sub>j</sub> ≤ 10<sup>6</sup>
) — the cultural tolerance of each user.

The seventh line contains m
 integers d<sub>1</sub>, d<sub>2</sub>, …, d<sub>m</sub>
 (0 ≤ d<sub>j</sub> ≤ 10<sup>6</sup>
) — the zone of influence of each user.

### Output
For each user, output one integer — the minimum influence `I(i,j)`
 among all news items.

### Examples:
**Input:**
```
6
2 4 1 6 0 10
3 2 6 1 9 0
5
0 0 0 1 5
0 0 9 5 2
9 4 8 2 2
```
**Output:**
```
5
4
1
2
2
```
<br><br>
**Input:**
```
5
75 19 53 12 10
34 75 67 84 95
5
55 14 46 97 14
78 61 56 23 33
10 4 7 11 3
```
**Output:**
```
0
18
53
34
36
```