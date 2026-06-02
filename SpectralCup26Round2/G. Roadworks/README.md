# G. Roadworks
## Time limit per test: 3 seconds
## Memory limit per test: 512 megabytes
### Description:
In an under-construction village, n houses have been built in a row numbered from 1 to n. House i has hospitality h<sub>i</sub>.

The village has n − 1 roads, where road i connects houses i and i + 1 and will be built on day d<sub>i</sub>. Initially, no roads are built.

You start at house x and will stay in the village from day 1 to day k, initially with a satisfaction of 0. On day s, the following happens in order:
- All roads i with d<sub>i</sub> = s are built;
- You may move to an adjacent house, if the road to it has been built, or stay at your current house;
- Your satisfaction increases by h<sub>j</sub>, where j is the house you are currently at. 

Find the maximum satisfaction you can achieve after k days.

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains three integers n, k and x (2 ≤ n ≤ 2 ⋅ 10<sup>5</sup>, 1 ≤ k ≤ 10<sup>9</sup>, 1 ≤ x ≤ n) — the number of houses, the number of days and the starting house, respectively.

The second line contains n integers h<sub>1</sub>, h<sub>2</sub>, …, h<sub>n</sub> (0 ≤ h<sub>i</sub> ≤ 10<sup>9</sup>) — the hospitality of each house.

The third line contains n−1 integers d<sub>1</sub>, d<sub>2</sub>, …, d<sub>n−1</sub> (1 ≤ d<sub>i</sub> ≤ k) — the day each road is built.

It is guaranteed that the sum of n over all test cases does not exceed 2 ⋅ 10<sup>5</sup>. 

### Output

For each test case, output a single integer — the maximum satisfaction you can achieve after k days.

### Examples:
**Sample Input:**
```
4
5 10 3
14 2 3 5 6
10 6 2 7
4 8 1
0 0 0 1
7 1 2
2 1000000000 1
1 1000000000
1
9 27 6
17 13 5 8 14 3 4 17 20
10 1 2 13 3 15 6 23
```
**Sample Output:**
```
52
0
1000000000000000000
386
```
**Explaination:**  
In the first test case, the following is one optimal sequence of moves:

* You start at house $x = 3$ with a satisfaction of $0$.
* On day $1$, no roads are built yet, so you must remain at house $3$. Your satisfaction becomes $3$.
* On day $2$, road $3$ is built. You move to house $4$ and remain there on days $2, 3, 4, 5, 6$ and $7$. Your satisfaction becomes $33$. During this time, roads $2$ and $4$ are also built.
* On day $8$, you move back to house $3$. Your satisfaction becomes $36$.
* On day $9$, you move to house $2$. Your satisfaction becomes $38$.
* On day $10$, road $1$ is built. You move to house $1$. Your satisfaction becomes $52$.

It can be shown that it is impossible to achieve a satisfaction greater than $52$.

In the second test case, you cannot reach house $4$ within $8$ days, so the maximum achievable satisfaction is $0$.

In the third test case, you can immediately move to house $2$ and remain there for $1\,000\,000\,000$ days.