# E. Lost Soul
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given two integer arrays a
 and b
, each of length n
.

You may perform the following operation any number of times:

- Choose an index i
 (1≤i≤n−1)
, and set a<sub>i</sub>:=b<sub>i+1</sub>
, or set b<sub>i</sub>:=a<sub>i+1</sub>
.

Before performing any operations, you are allowed to choose an index i
 (1≤i≤n)
 and remove both a<sub>i</sub>
 and b<sub>i</sub>
 from the arrays. This removal can be done at most once.

Let the number of matches between two arrays c
 and d
 of length m
 be the number of positions j
 (1≤j≤m)
 such that c<sub>j</sub>=d<sub>j</sub>
.

Your task is to compute the maximum number of matches you can achieve.

### Input:
The first line of the input contains an integer t
 (1≤t≤10<sup>4</sup>)
 — the number of test cases. The description of each test case follows.

The first line contains an integer n
 (2≤n≤2⋅10<sup>5</sup>)
 — the length of a
 and b
.

The second line contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1≤a<sub>i</sub>≤n)
 — the elements of a
.

The third line contains n
 integers b<sub>1</sub>,b<sub>2</sub>,…,b<sub>n</sub>
 (1≤b<sub>i</sub>≤n)
 — the elements of b
.

It is guaranteed that the sum of n
 over all test cases does not exceed 2⋅10<sup>5</sup>
.

### Output:
For each test case, print a single integer — the answer for the test case.

### Examples:
**Input:**
```
10
4
1 3 1 4
4 3 2 2
6
2 1 5 3 6 4
3 2 4 5 1 6
2
1 2
2 1
6
2 5 1 3 6 4
3 5 2 3 4 6
4
1 3 2 2
2 1 3 4
8
3 1 4 6 2 2 5 7
4 2 3 7 1 1 6 5
10
5 1 2 7 3 9 4 10 6 8
6 2 3 6 4 10 5 1 7 9
5
3 2 4 1 5
2 4 5 1 3
7
2 2 6 4 1 3 5
3 1 6 5 1 4 2
5
4 1 3 2 5
3 2 1 5 4
```
**Output:**
```
3
3
0
4
3
5
6
4
5
2
```
**Explaination:**  
In the first test case, we can do the following:

- We will choose not to remove any index.
- Choose index 3
, and set a<sub>3</sub>:=b<sub>4</sub>
. The arrays become: a=`[1,3,2,4]`
, b=`[4,3,2,2]`
.
- Choose index 1
, and set a<sub>1</sub>:=b<sub>2</sub>
. The arrays become: a=`[3,3,2,4]`
, b=`[4,3,2,2]`
.
- Choose index 1
, and set b<sub>1</sub>:=a<sub>2</sub>
. The arrays become: a=`[3,3,2,4]`
, b=`[3,3,2,2]`
. Notice that you can perform a<sub>i</sub>:=b<sub>i+1</sub>
 and b<sub>i</sub>:=a<sub>i+1</sub>
 on the same index i
.
The number of matches is 3
. It can be shown that this is the maximum answer we can achieve.

In the second test case, we can do the following to achieve a maximum of 3
:

- Let's choose to remove index 5
. The arrays become: a=`[2,1,5,3,4]`
, b=`[3,2,4,5,6]`
.
- Choose index 4
, and set b<sub>4</sub>:=a<sub>5</sub>
. The arrays become: a=`[2,1,5,3,4]`
, b=`[3,2,4,4,6]`
.
- Choose index 3
, and set a<sub>3</sub>:=b<sub>4</sub>
. The arrays become: a=`[2,1,4,3,4]`
, b=`[3,2,4,4,6]`
.
- Choose index 2
, and set a<sub>2</sub>:=b<sub>3</sub>
. The arrays become: a=`[2,4,4,3,4]`
, b=`[3,2,4,4,6]`
.
- Choose index 1
, and set b<sub>1</sub>:=a<sub>2</sub>
. The arrays become: a=`[2,4,4,3,4]`
, b=`[4,2,4,4,6]`
.
- Choose index 2
, and set b<sub>2</sub>:=a<sub>3</sub>
. The arrays become: a=`[2,4,4,3,4]`
, b=`[4,4,4,4,6]`
.
- Choose index 1
, and set a<sub>1</sub>:=b<sub>2</sub>
. The arrays become: a=`[4,4,4,3,4]`
, b=`[4,4,4,4,6]`
.

In the third test case, it can be shown that we can not get any matches. Therefore, the answer is 0
.