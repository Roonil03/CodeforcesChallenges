# F. Binary Not Search and Queries
## Time limit per test: 4 seconds
## Memory limit per test: 512 megabytes
### Description:
For a sequence b
 consisting of m
 integers, the set `S(b)`
 is defined as the set of tuples `(i,j,k)`
 that satisfy the following conditions:
- i
, j
, k
 are integers;
- 1 ≤ k < m
;
- 1 ≤ i < j ≤ `m − k + 1`
;
- For every element v
 in b
, v
 appears the same number of times in `[b`<sub>`i`</sub>`,b`<sub>`i+1`</sub>`,…,b`<sub>`i+k−1`</sub>`]`
 and `[b`<sub>`j`</sub>`,b`<sub>`j+1`</sub>`,…,b`<sub>`j+k−1`</sub>`]`
.

For example, when b = `[1,2,1,2]`
, the tuple `(1,3,2)`
 is an element of `S(b)`
 because 1
 and 2
 both appear once in `[b`<sub>`1,b`<sub>`2`</sub>`]`
 and `[b`<sub>`3`</sub>`,b`<sub>`4`</sub>`]`
.

Additionally, we define two functions over sequences of positive integers:
- `k`<sub>`max`</sub>`(b)`
 is defined as the maximum value of k
 over all elements `(i,j,k)`
 of `S(b)`
;
- `f(b)`
 is defined as the number of different elements `(i,j,k)`
 of `S(b)`
 such that `k` = `k`<sub>`max`</sub>`(b)`
.

Exceptionally, when the set `S(b)`
 is empty, they are defined as `k`<sub>`max`</sub>`(b)` = 0
 and `f(b)` = 0
.

You are given a sequence a
 of n
 integers. Please answer q
 queries of the following kind:
- i<sub>x</sub>
: Change the value of a<sub>i</sub>
 to x
. Then, find the values of k<sub>max</sub>(a)
 and `f(a)`
.

Do note that the updates are persistent. In other words, the update from one query affects the later queries as well.

### Input:
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains two integers n
 and q
 (2 ≤ n ≤ 200000
, 1 ≤ q ≤ 100000
).

The second line of each test case contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ n
).

Each of the following q
 lines contains two integers i<sub>j</sub>
 and x<sub>j</sub>
 denoting the j
-th query (1 ≤ i<sub>j</sub> , x<sub>j</sub> ≤ n
).

It is guaranteed that the sum of n
 over all test cases does not exceed 200000
.

It is guaranteed that the sum of q
 over all test cases does not exceed 100000
.

### Output
For each test case, output q
 lines.

On the j
-th line, you must output the values of `k`<sub>`max`</sub>`(a)`
 and `f(a)`
 for the j
-th query.

It can be shown that both values will never exceed 10<sup>11</sup>
 under the constraints of this problem.

### Examples:
**Input:**
```
4
5 3
1 2 3 4 5
3 2
4 1
5 2
4 3
1 2 1 1
4 2
3 2
2 1
5 2
1 3 2 4 5
5 3
5 5
8 3
1 2 3 4 1 2 5 4
7 3
7 4
2 1
```
**Output:**
```
1 1
3 1
3 3
2 3
2 1
1 2
3 1
0 0
4 10
4 4
4 2
```
**Note:**  
Immediately after the first query of the second test case, a = `[1,2,1,2]`
. The elements of the set `S(a)`
 are as follows:
- `(1,3,1)`
: `[1,2,1,2]`
;
- `(2,4,1)`
: `[1,2,1,2]`
;
- `(1,2,2)`
: `[1,2,1,2]`
;
- `(1,3,2)`
: `[1,2,1,2]`
;
- `(2,3,2)`
: `[1,2,1,2]`
.

Therefore, k<sub>max</sub>(a) = 2
, and `f(a)` = 3
 because there are three elements `(i,j,k)`
 where k = k<sub>max</sub>(a) = 2
.

Immediately after the second query of the third test case, a = `[1,3,2,4,5]`. The set `S(a)`
 is empty at this point.

By definition, you should output k<sub>max</sub>(a) = 0
 and `f(a)` = 0
 because `S(a)`
 is currently empty.