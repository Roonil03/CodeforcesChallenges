# E. Matrix Puzzle
## Time limit per test: 2.5 seconds
## Memory limit per test: 256 megabytes
### Description:
Camila and Daniel are the greatest puzzle rivals in Mexico. For years, they have challenged each other with the hardest number games in the country, and now they have turned their attention to you. You want to prove to them that you can compete at their level, and so they present you with the following problem.

You are given two square matrices $a$ and $b$ of size $n \times n$. The submatrix $(r_1, c_1, r_2, c_2)$ of a matrix consists of the elements in rows $r_1$ to $r_2$ and columns $c_1$ to $c_2$. The rows are numbered from $1$ to $n$ from top to bottom and the columns are numbered from $1$ to $n$ from left to right.

In one *move* on a submatrix of $a$, you pick one of its rows or columns and add any constant (possibly negative) to every element of that row or column inside the submatrix. The submatrix $(r_1, c_1, r_2, c_2)$ of $a$ is *matchable* if some sequence of moves makes it equal to the same submatrix of $b$. In other words, it is matchable if and only if you can choose a value $h_i$ for each row $i$ and a value $w_j$ for each column $j$ such that $a_{i,j} + h_i + w_j = b_{i,j}$ for all $r_1 \le i \le r_2$ and $c_1 \le j \le c_2$.

You will be asked to perform $q$ operations.

The operations are of two types:

* $\texttt{U}\ r_1\ c_1\ r_2\ c_2\ x$: Add a constant $x$ to every $a_{i,j}$ in the submatrix $(r_1, c_1, r_2, c_2)$.
* $\texttt{Q}\ r_1\ c_1\ r_2\ c_2$: Determine whether the submatrix $(r_1, c_1, r_2, c_2)$ of $a$ is *matchable*.

Queries do not modify either matrix: the moves are hypothetical and are chosen independently for each query. Only update operations change $a$, and $b$ never changes.

### Input
Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains two integers $n$ and $q$ ($2 \le n \le 1000$, $1 \le q \le 2 \cdot 10^4$) — the dimension of the matrices and the number of operations.

The next $n$ lines of each test case contain $n$ integers each ($-10^4 \le a_{i,j} \le 10^4$) — the initial state of matrix $a$.

The next $n$ lines of each test case contain $n$ integers each ($-10^4 \le b_{i,j} \le 10^4$) — the matrix $b$.

Then $q$ lines follow for that test case. Each line starts with a character $\texttt{U}$ or $\texttt{Q}$:

* $\texttt{U}$: then five integers $r_1, c_1, r_2, c_2, x$ ($1 \le r_1 \le r_2 \le n$, $1 \le c_1 \le c_2 \le n$, $-10^4 \le x \le 10^4$) — update the submatrix $(r_1, c_1, r_2, c_2)$ by adding $x$ to each of its elements.
* $\texttt{Q}$: then four integers $r_1, c_1, r_2, c_2$ ($1 \le r_1 \le r_2 \le n$, $1 \le c_1 \le c_2 \le n$) — query the submatrix $(r_1, c_1, r_2, c_2)$.

It is guaranteed that the sum of $n^2$ over all test cases does not exceed $10^6$.

It is guaranteed that the sum of $q$ over all test cases does not exceed $2 \cdot 10^4$.

### Output
For each query, output "$\texttt{YES}$" (without quotes) if the submatrix is matchable, and "$\texttt{NO}$" (without quotes) otherwise.

You can output the answer in any case (upper or lower). For example, the strings "$\texttt{yEs}$", "$\texttt{yes}$", "$\text{Yes}$", and "$\texttt{YES}$" will be recognized as positive responses.

### Examples:
**Input:**
```
3
2 2
0 0
0 -1
1 2
3 4
U 1 1 1 1 1
Q 1 1 2 2
3 5
1 1 1
2 2 2
3 3 3
2 3 4
6 7 8
10 11 12
Q 1 1 2 3
U 1 1 1 1 5
Q 1 1 2 3
U 1 2 1 3 5
Q 1 1 2 3
5 6
1 0 0 0 0
2 0 0 0 0
3 0 0 0 0
4 0 0 0 0
5 0 0 0 0
2 2 3 4 5
4 3 4 5 6
6 4 5 6 7
8 5 6 7 8
10 6 7 8 9
U 2 2 3 3 7
U 2 4 4 5 -5
Q 1 1 4 5
Q 2 4 4 5
Q 2 1 3 5
Q 1 1 2 3
```
**Output:**
```
Yes
Yes
No
Yes
No
Yes
Yes
No
```
**Explaination:**  
In the first test case, after the update, $a$ becomes

$$\begin{bmatrix} 1 & 0 \\ 0 & -1 \end{bmatrix}.$$

For the query, choosing $h_1 = 0$, $h_2 = 3$, $w_1 = 0$, and $w_2 = 2$ makes the submatrices match, so the answer is $\texttt{YES}$.

In the second test case, the first query has answer $\texttt{YES}$: choose $h_1 = 1$, $h_2 = 4$, $w_1 = 0$, $w_2 = 1$, and $w_3 = 2$. After the first update, the submatrices of $a$ and $b$ in the second query are

$$\begin{bmatrix} 6 & 1 & 1 \\ 2 & 2 & 2 \end{bmatrix}, \quad \begin{bmatrix} 2 & 3 & 4 \\ 6 & 7 & 8 \end{bmatrix},$$

and it can be proved that the answer is $\texttt{NO}$. After the second update, in the third query they are

$$\begin{bmatrix} 6 & 6 & 6 \\ 2 & 2 & 2 \end{bmatrix}, \quad \begin{bmatrix} 2 & 3 & 4 \\ 6 & 7 & 8 \end{bmatrix},$$

so choosing $h_1 = -4$, $h_2 = 4$, $w_1 = 0$, $w_2 = 1$, and $w_3 = 2$ makes them match, and the answer is $\texttt{YES}$.

In the third test case, after both updates, the submatrices of $a$ and $b$ in the first query are

$$\begin{bmatrix} 1 & 0 & 0 & 0 & 0 \\ 2 & 7 & 7 & -5 & -5 \\ 3 & 7 & 7 & -5 & -5 \\ 4 & 0 & 0 & -5 & -5 \end{bmatrix}, \quad \begin{bmatrix} 2 & 2 & 3 & 4 & 5 \\ 4 & 3 & 4 & 5 & 6 \\ 6 & 4 & 5 & 6 & 7 \\ 8 & 5 & 6 & 7 & 8 \end{bmatrix},$$

and it can be proved that the answer is $\texttt{NO}$. In the second query, they are

$$\begin{bmatrix} -5 & -5 \\ -5 & -5 \\ -5 & -5 \end{bmatrix}, \quad \begin{bmatrix} 5 & 6 \\ 6 & 7 \\ 7 & 8 \end{bmatrix},$$

so choosing $h_2 = 10$, $h_3 = 11$, $h_4 = 12$, $w_4 = 0$, and $w_5 = 1$ makes them match, and the answer is $\texttt{YES}$. In the third query, choosing $h_2 = 0$, $h_3 = 1$, $w_1 = 2$, $w_2 = -4$, $w_3 = -3$, $w_4 = 10$, and $w_5 = 11$ makes the submatrices match, and the answer is $\texttt{YES}$. It can be proved that the answer to the fourth query is $\texttt{NO}$.