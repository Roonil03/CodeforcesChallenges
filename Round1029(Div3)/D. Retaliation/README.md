# D. Retaliation
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Yousef wants to explode an array a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
. An array gets exploded when all of its elements become equal to zero.

In one operation, Yousef can do exactly one of the following:

1. For every index i
 in a
, decrease a<sub>i</sub>
 by i
.
2. For every index i
 in a
, decrease a<sub>i</sub>
 by n−i+1
.
Your task is to help Yousef determine if it is possible to explode the array using any number of operations.

### Input:
The first line of the input contains an integer t
 (1≤t≤10<sup>4</sup>
) — the number of test cases.

The first line of each test case contains an integer n
 (2≤n≤2⋅10<sup>5</sup>
) — the size of the array.

The second line of each test case contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1≤a<sub>i</sub>≤10<sup>9</sup>
) — the elements of the array.

It is guaranteed that the sum of n
 over all test cases doesn't exceed 2⋅10<sup>5</sup>
.


### Output
For each test case, print "`YES`" if Yousef can explode the array, otherwise output "`NO`".

You can output the answer in any case (upper or lower). For example, the strings "`yEs`", "`yes`", "`Yes`", and "`YES`" will be recognized as positive responses.

### Examples:
**Input:**
```
6
4
3 6 6 3
5
21 18 15 12 9
10
2 6 10 2 5 5 1 2 4 10
7
10 2 16 12 8 20 4
2
52 101
2
10 2
```
**Output:**
```
NO
YES
NO
NO
YES
NO
```
**Explaination:**  
In the second test case, we can do the following:

- Perform 1
 operation of the first type. The array becomes `[20,16,12,8,4]`
.
- Perform 4
 operations of the second type. The array becomes `[0,0,0,0,0]`
.

In the first, third, fourth, and sixth test cases, it can be proven that it is impossible to make all elements equal to zero using any number of operations.