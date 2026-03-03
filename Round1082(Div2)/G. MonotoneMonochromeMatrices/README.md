# G1. Monotone Monochrome Matrices (Easy Version)
## Time limit per test: 4 seconds
## Memory limit per test: 512 megabytes
**This is the easy version of the problem. The difference between the versions is that in this version, the constraints on n
 and q
 are smaller. You can hack only if you solved all versions of this problem.**
### Description:

A monochrome matrix of size `n×n`
 is a matrix of n
 rows and n
 columns, where each cell is colored either black or white. Let the color of cell (r,c)
 in a monochrome matrix C
 be denoted as C[r,c]
.

Let's call such a matrix C
 monotone if it satisfies the following condition:
- There exist no two rows 1 ≤ i < j ≤ n
 and two columns 1 ≤ k < l ≤ n
 that satisfy the following three conditions:
    - `C[i,k]` = `C[j,l]`
;
    - `C[j,k]` = `C[i,l]`
;
    - `C[i,k]` ≠ `C[j,k]`
.

There is a monochrome matrix M
 of size `n×n`
, where all cells are initially white. Please solve q
 queries of the following kind:
- `r c`
: Change the color of the cell `(r,c)`
 in M
 to black. Then, determine if M
 is monotone or not.

For each query, it is guaranteed that the color of the cell `(r,c)`
 was white before the query.

Do note that the updates are persistent. In other words, the change in color from one query affects the later queries as well.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>3</sup>
). The description of the test cases follows.

The first line of each test case contains two integers n
 and q
 (2 ≤ n ≤ 25000
, 1 ≤ q ≤ min(n<sup>2</sup>,200000)
).

Each of the q
 following lines contains two integers r<sub>i</sub>
, c<sub>i</sub>
 denoting the i
-th query (1 ≤ r<sub>i</sub>, c<sub>i</sub> ≤ n
).

For each query, it is guaranteed that the color of the cell `(r,c)`
 was white before the query.

It is guaranteed that the sum of n
 over all test cases does not exceed 25000.

It is guaranteed that the sum of q
 over all test cases does not exceed 200000.

### Output
For each query, output "YES" or "NO" on a separate line based on the answer of the query.

You can output the answer in any case. For example, the strings "yEs", "yes", and "Yes" will also be recognized as positive responses.

### Examples:
**Input:**
```
2
3 9
2 2
3 3
2 3
3 1
3 2
1 1
1 2
2 1
1 3
5 17
2 1
4 5
4 1
3 3
3 1
3 5
1 3
1 5
1 1
5 3
5 5
5 1
1 4
5 2
5 4
1 2
2 5
```
**Output:**
```
YES
NO
YES
NO
YES
NO
NO
YES
YES
YES
NO
YES
NO
NO
YES
NO
NO
YES
NO
NO
YES
YES
NO
YES
YES
YES
```
**Note:**  
In the first test case, M
 has size 3×3
. The states of M
 after each query are as shown below.

$$
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \square \\
\square & \square & \square 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \square \\
\square & \square & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \blacksquare \\
\square & \square & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \square & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \square & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \blacksquare & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \blacksquare & \square \\
\blacksquare & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
$$

On queries where M
 is not monotonic, the four squares highlighted in red denote cells that violate the condition stated above.

# G2. Monotone Monochrome Matrices (Hard Version)
## Time limit per test: 4 seconds
## Memory limit per test: 512 megabytes
**This is the hard version of the problem. The difference between the versions is that in this version, the constraints on n
 and q
 are very large. You can hack only if you solved all versions of this problem.**
### Description:
A monochrome matrix of size `n×n`
 is a matrix of n
 rows and n
 columns, where each cell is colored either black or white. Let the color of cell (r,c)
 in a monochrome matrix C
 be denoted as C[r,c]
.

Let's call such a matrix C
 monotone if it satisfies the following condition:
- There exist no two rows 1 ≤ i < j ≤ n
 and two columns 1 ≤ k < l ≤ n
 that satisfy the following three conditions:
    - `C[i,k]` = `C[j,l]`
;
    - `C[j,k]` = `C[i,l]`
;
    - `C[i,k]` ≠ `C[j,k]`
.

There is a monochrome matrix M
 of size `n×n`
, where all cells are initially white. Please solve q
 queries of the following kind:
- `r c`
: Change the color of the cell `(r,c)`
 in M
 to black. Then, determine if M
 is monotone or not.

For each query, it is guaranteed that the color of the cell `(r,c)`
 was white before the query.

Do note that the updates are persistent. In other words, the change in color from one query affects the later queries as well.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>3</sup>
). The description of the test cases follows.

The first line of each test case contains two integers n
 and q
 (2 ≤ n ≤ 2000000
, 1 ≤ q ≤ min(n<sup>2</sup>,2000000)
).

Each of the q
 following lines contains two integers r<sub>i</sub>
, c<sub>i</sub>
 denoting the i
-th query (1 ≤ r<sub>i</sub>, c<sub>i</sub> ≤ n
).

For each query, it is guaranteed that the color of the cell `(r,c)`
 was white before the query.

It is guaranteed that the sum of n
 over all test cases does not exceed 2000000
.

It is guaranteed that the sum of q
 over all test cases does not exceed 2000000
.
### Output
For each query, output "YES" or "NO" on a separate line based on the answer of the query.

You can output the answer in any case. For example, the strings "yEs", "yes", and "Yes" will also be recognized as positive responses.

### Examples:
**Input:**
```
2
3 9
2 2
3 3
2 3
3 1
3 2
1 1
1 2
2 1
1 3
5 17
2 1
4 5
4 1
3 3
3 1
3 5
1 3
1 5
1 1
5 3
5 5
5 1
1 4
5 2
5 4
1 2
2 5
```
**Output:**
```
YES
NO
YES
NO
YES
NO
NO
YES
YES
YES
NO
YES
NO
NO
YES
NO
NO
YES
NO
NO
YES
YES
NO
YES
YES
YES
```
**Note:**  
In the first test case, M
 has size 3×3
. The states of M
 after each query are as shown below.

$$
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \square \\
\square & \square & \square 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \square \\
\square & \square & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \blacksquare \\
\square & \square & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \square & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\square & \square & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \square & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \blacksquare & \square \\
\square & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \blacksquare & \square \\
\blacksquare & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
\rightarrow
\begin{bmatrix}
\blacksquare & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare \\
\blacksquare & \blacksquare & \blacksquare 
\end{bmatrix}
$$

On queries where M
 is not monotonic, the four squares highlighted in red denote cells that violate the condition stated above.