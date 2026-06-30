# B. Crimson Triples
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
<i><div align=right><sub>
Chills run down your spine...

---

 -Terraria
</i></sub></div>

After summoning the next boss — The Brain of Cthulhu, you noticed that it surrounds itself with $n$ eyes, numbered from $1$ to $n$. In one attack, The Brain of Cthulhu chooses a triple of eyes (not necessarily distinct) with numbers $(a, b, c)$. The triple of eyes is called *crimson* if and only if the following property holds:

$$\gcd(\text{lcm}(a, b), \text{lcm}(b, c)) = \gcd(a, c)$$

To defeat the boss, you want to know how many ways The Brain of Cthulhu can choose a crimson triple of eyes. The triples of eyes $(a_1, b_1, c_1)$ and $(a_2, b_2, c_2)$ are considered different if $a_1 \neq a_2$, or $b_1 \neq b_2$, or $c_1 \neq c_2$.

<sub>

<sup>∗</sup>
$\gcd(x, y)$ denotes the greatest common divisor (GCD) of integers $x$ and $y$.  
<sup>†</sup>
$\text{lcm}(x, y)$ denotes the least common multiple (LCM) of integers $x$ and $y$.

</sub>
### Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 1000$). The description of the test cases follows.

The only line of each test case contains one integer $n$ ($1 \le n \le 2 \cdot 10^5$) — the number of eyes of The Brain of Cthulhu.

It is guaranteed that the sum of $n$ across all test cases does not exceed $2 \cdot 10^5$.

### Output

For each test case, output one integer — the number of ways to choose a crimson triple of eyes.

### Examples:
**Input:**
```
3
1
2
20
```
**Output:**
```
1
5
612
```
**Explaination:**  
In the first test case, there is $1$ possible crimson triple: $(1, 1, 1)$.

In the second test case, there are $5$ possible crimson triples: $(1, 1, 1), (1, 1, 2), (2, 1, 1), (2, 1, 2), (2, 2, 2)$.