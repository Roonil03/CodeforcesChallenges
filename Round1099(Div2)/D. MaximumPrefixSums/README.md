# D. Maximum Prefix Sums
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
521 sounds the same as "I love you" in Chinese. On this special day, Little S wants to give Little A some well-prepared sequences to recall their friendship sealed for years.

Little S has prepared an array a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
. Let's define its prefix sums array b<sub>1</sub>, b<sub>2</sub>, …, b<sub>n</sub>
, where b<sub>i</sub> = a<sub>1</sub> + a<sub>2</sub> + … + a<sub>i</sub>
. Also define the prefix maximum array of b
: c<sub>1</sub>, c<sub>2</sub>, …, c<sub>n</sub>
, where c<sub>i</sub> = max(b<sub>1</sub>, b<sub>2</sub>, …, b<sub>i</sub>)
.

Now array a
 has been partially lost, but luckily Little S still keeps the array c
. Your task is to restore the array a
 and send it to Little A, or report that no such array exists — in this case, the kind and cute Little A won't get mad either.

Formally, you are given a binary string s
, a partially filled array a
, and an array c
, where:

- If you remember the value of a<sub>i</sub>
, then s<sub>i</sub> = 1
, and you are given the true value of a<sub>i</sub>
.
- If you do not remember the value of a<sub>i</sub>
, then s<sub>i</sub> = 0
, and you are given a<sub>i</sub> = 0
.

### Input:
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains an integer n
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
).

The second line of each test case contains a binary string s
 (s<sub>i</sub> ∈ `{0,1}`
) of length n
.

The third line of each test case contains n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (`|a`<sub>`i`</sub>`|` ≤ 10<sup>6</sup>
). If s<sub>i</sub> = 0
, then it is guaranteed that a<sub>i</sub> = 0
.

The fourth line of each test case contains n
 integers c<sub>1</sub>, c<sub>2</sub>, …, c<sub>n</sub>
 (`|c`<sub>`i`</sub>`|` ≤ 2 ⋅ 10<sup>11</sup>
).

It is guaranteed that the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, print "`Yes`" in the first line if the array a
 exists, otherwise print "`No`". You may print the answer in any case. For example, "`YeS`", "`YES`", "`NO`", "`nO`" will also be accepted.

If there is at least one solution, print n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (`|a`<sub>`i`</sub>`|` ≤ 10<sup>18</sup>
) in the second line. If there are multiple suitable arrays, you may print any of them

### Examples:
**Input:**
```
10
4
1110
1 2 -1 0
1 3 3 3
5
00001
0 0 0 0 5
-4 -4 -1 -1 -1
1
1
0
1
6
001111
0 0 2 -3 3 -6
-5 -2 0 0 0 0
5
11110
1 2 0 5 0
1 2 2 7 6
2
01
0 1
-1 -1
6
001010
0 0 5 0 3 0
3 3 4 9 13 16
6
000100
0 0 0 4 0 0
2 6 6 7 7 7
2
00
0 0
4 -1
8
11111111
6 1 1 2 0 5 1 9
6 7 8 10 10 15 16 25
```
**Output:**
```
Yes
1 2 -1 0
Yes
-4 0 3 -6 5
No
Yes
-5 3 2 -3 3 -6
No
No
No
Yes
2 4 -3 4 -100 0
No
Yes
6 1 1 2 0 5 1 9
```
**Explaination:**  
In the first test case, we have:

- a = `[1,2,−1,0]`
.
- b = `[1,3,2,2]`
.
- c = `[1,3,3,3]`
.

In the third test case, the correct array a
 should be equal to `[1]`
, so no solution exists.

