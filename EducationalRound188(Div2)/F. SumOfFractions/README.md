# F. Sum of Fractions
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:

Let's call an increase of the fraction $\frac{x}{y}$ one of the following operations:

- either increase its numerator $x$ by $1$;
- or, if the denominator $y > 1$, decrease its denominator $y$ by $1$.

Note that the fraction is not reduced after the increase.  
For example, if you have a fraction $\frac{2}{7}$ and you decrease its denominator by $1$, you get $\frac{2}{6}$, not $\frac{1}{2}$.

---

Suppose you are given an integer array $b_1, b_2, \dots, b_m$ and an integer $k$.  
You apply the following algorithm:

1. Construct the array

   $$
   \frac{1}{b_1}, \frac{1}{b_2}, \dots, \frac{1}{b_m}
   $$

2. Choose any of the fractions and increase it.
3. Repeat the previous step $k$ times (the same fraction can be chosen multiple times).
4. Calculate the sum of the resulting fractions.

---

We will denote $\text{MSF}(b, k)$ as the maximum sum of fractions that can be obtained from the array $b$ by applying the increase operation exactly $k$ times.

---

Let $a[l \dots r]$ be the subarray

$$
a_l, a_{l+1}, \dots, a_r
$$

of the array $a$.

You are given two arrays of integers:

$$
a_1, a_2, \dots, a_n
$$

and

$$
k_1, k_2, \dots, k_m
$$

For each $k_i$, calculate

$$
\left(
\sum_{l=1}^{n}
\sum_{r=l}^{n}
\text{MSF}(a[l \dots r], k_i)
\right)
\mod 998244353
$$

In other words, for each $k_i$, calculate the sum of $\text{MSF}$ over all subarrays of the array $a$ and output the answer modulo $998244353$.

### Input
The first line contains two integers n
 and m
 (1 ≤n,m ≤5⋅10<sup>5</sup>
) — the sizes of the arrays a
 and k
.

The second line contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ 10<sup>8</sup>
) — the array a
.

The third line contains m
 integers k<sub>1</sub>,k<sub>2</sub>,…,k<sub>m</sub>
 (0 ≤ k<sub>1</sub> ≤ k<sub>2</sub> ≤ ⋯ ≤ k<sub>m</sub> ≤ 10<sup>8</sup>
) — the array k
 in non-decreasing order.

### Output
For each k<sub>i</sub>
, output a single integer — the sum of MSF
 over all subarrays modulo 998244353
.

It can be proven that the answer can be represented as an irreducible fraction P / Q
, where Q !≡ 0 (mod998244353)
. Accordingly, output the answer in the form P⋅Q<sup>−1</sup>(mod 998244353)
.

### Examples:
**Input:**
```
5 4
2 3 5 2 3
0 1 2 10
```
**Output:**
```
232923695
332748137
931694761
133099397
```
**Notes:**  
The answers for the corresponding k
 are: $\frac{379}{30}$
, $\frac{58}{3}$
, $\frac{473}{15}$
, and $\frac{2249}{15}$
.   