# E. Array Emptying
## time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given an array $a$ of length $n$ and an integer $k$.

In one operation, choose any element $x$ of the array.

* If $x \le k$, erase $x$ from the array.
* Otherwise, erase $x$ and insert two elements equal to $\left\lceil \frac{x}{2} \right\rceil$.

In both cases, increase $k$ by $1$ after performing the operation.

Find the minimum number of operations required to make the array empty.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains two integers $n$ and $k$ ($1 \le n \le 2 \cdot 10^5$, $0 \le k \le n$) — the length of the array and the initial value of $k$.

The second line contains $n$ integers $a_1, a_2, \dots, a_n$ ($1 \le a_i \le n$) — the elements of the array.

It is guaranteed that the sum of $n$ over all test cases does not exceed $2 \cdot 10^5$.

### Output
For each test case, output a single integer — the minimum number of operations required to make the array empty.

### Examples:
**Input:**
```
8
4 0
4 4 1 4
3 0
3 2 2
6 3
6 3 6 4 6 5
6 2
6 6 5 6 6 6
5 0
3 5 3 1 3
6 1
6 4 6 5 4 6
7 1
7 2 7 7 3 7 7
8 0
8 8 2 8 8 1 8 8
```
**Output:**
```
6
5
6
8
7
8
11
12
```
**Explaination:**  
### Note
In the first example, initially $k = 0$. We can perform the following operations:

1. Split an element equal to $4$ into two elements equal to $2$; $k$ becomes $1$.
2. Erase an element equal to $1$; $k$ becomes $2$.
3. Erase an element equal to $2$; $k$ becomes $3$.
4. Erase an element equal to $2$; $k$ becomes $4$.
5. Erase an element equal to $4$; $k$ becomes $5$.
6. Erase an element equal to $4$; $k$ becomes $6$.

The array becomes empty after $6$ operations.

In the fourth example, initially $k = 2$. We can perform the following operations:

1. Split an element equal to $6$ into two elements equal to $3$; $k$ becomes $3$.
2. Erase an element equal to $3$; $k$ becomes $4$.
3. Erase an element equal to $3$; $k$ becomes $5$.
4. Erase an element equal to $5$; $k$ becomes $6$.
5. Erase an element equal to $6$; $k$ becomes $7$.
6. Erase an element equal to $6$; $k$ becomes $8$.
7. Erase an element equal to $6$; $k$ becomes $9$.
8. Erase an element equal to $6$; $k$ becomes $10$.

The array becomes empty after $8$ operations.