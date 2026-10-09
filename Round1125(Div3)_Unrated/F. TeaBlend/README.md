# F. Tea Blend
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Recently, K1o0n made an impulsive purchase of $n$ packs of tea and is now wondering when he will manage to try all the blends. He is picky about tea, so he immediately assigned each pack a strength — an integer $a_i$.

The packs are arranged on the shelf in the order they were bought, and K1o0n is too lazy to sort them out: when brewing tea, he pours the first $j$ packs in order into the teapot, and then adds one more spoon from the pack with number $i$. We denote the resulting blend by $(i, j)$, and define its strength as the product of the strengths of everything that went into it:

$$F(i, j) = a_i \cdot \prod_{k=1}^j a_k = a_i \cdot (a_1 \cdot a_2 \cdot \dots \cdot a_j)$$

K1o0n pours the finished teapot into identical cups: he can pour it into $d$ cups if and only if $d$ — is a natural divisor of the number $F(i, j)$, and each cup gets the same amount. He calls a blend *capricious* if the number of suitable $d$ is odd.

Count how many *capricious* blends K1o0n can brew, that is, how many pairs $(i, j)$ with $1 \le i \le n$ and $1 \le j \le n$ give an odd number of natural divisors of the number $F(i, j)$.

### Input
The first line contains an integer $t$ ($1 \le t \le 10^4$) — the number of testcases.

The first line of each testcase contains an integer $n$ ($1 \le n \le 2 \cdot 10^5$) — the number of tea packs.

The second line of each testcase contains $n$ integers $a_1, a_2, \dots, a_n$ ($1 \le a_i \le 10^6$) — the strengths of the teas.

It is guaranteed that the sum of $n$ over all testcases does not exceed $2 \cdot 10^5$.

### Output
For each testcase, output one integer — the number of *capricious* blends.

Note that the answer may not fit into a 32-bit type. Use a 64-bit type (`long long` in C++, `long` in Java).

### Examples:
**Input:**
```
4
4
2 3 6 4
4
974153 980051 988027 1
8
12 3 5 45 1 7 63 2
1
13
```
**Output:**
```
4
4
11
1
```
