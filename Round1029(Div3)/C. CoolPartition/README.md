# C. Cool Partition
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
Yousef has an array a
 of size n
. He wants to partition the array into one or more contiguous segments such that each element a<sub>i</sub>
 belongs to exactly one segment.

A partition is called cool if, for every segment b<sub>j</sub>
, all elements in b<sub>j</sub>
 also appear in b<sub>j+1</sub>
 (if it exists). That is, every element in a segment must also be present in the segment following it.

For example, if a=`[1,2,2,3,1,5]`
, a cool partition Yousef can make is b<sub>1</sub>=`[1,2]`
, b<sub>2</sub>=`[2,3,1,5]`
. This is a cool partition because every element in b<sub>1</sub>
 (which are 1
 and 2
) also appears in b<sub>2</sub>
. In contrast, b<sub>1</sub>=`[1,2,2]`
, b<sub>2</sub>=`[3,1,5]`
 is not a cool partition, since 2
 appears in b<sub>1</sub>
 but not in b<sub>2</sub>
.

Note that after partitioning the array, you do not change the order of the segments. Also, note that if an element appears several times in some segment b<sub>j</sub>
, it only needs to appear at least once in b<sub>j+1</sub>
.

Your task is to help Yousef by finding the maximum number of segments that make a cool partition.

### Input:
The first line of the input contains integer t
 (1≤t≤10<sup>4</sup>
) — the number of test cases.

The first line of each test case contains an integer n
 (1≤n≤2⋅10<sup>5</sup>
) — the size of the array.

The second line of each test case contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1≤a<sub>i</sub>≤n
) — the elements of the array.

It is guaranteed that the sum of n
 over all test cases doesn't exceed 2⋅10<sup>5</sup>
.

### Output:
For each test case, print one integer — the maximum number of segments that make a cool partition.

### Examples:
**Input:**
```
8
6
1 2 2 3 1 5
8
1 2 1 3 2 1 3 2
5
5 4 3 2 1
10
5 8 7 5 8 5 7 8 10 9
3
1 2 2
9
3 3 1 4 3 2 4 1 2
6
4 5 4 5 6 4
8
1 2 1 2 1 2 1 2
```
**Output:**
```
2
3
1
3
1
3
3
4
```
**Explaination:**  
The first test case is explained in the statement. We can partition it into b<sub>1</sub>=`[1,2]`
, b<sub>2</sub>=`[2,3,1,5]`
. It can be shown there is no other partition with more segments.

In the second test case, we can partition the array into b<sub>1</sub>=`[1,2]`
, b<sub>2</sub>=`[1,3,2]`
, b<sub>3</sub>=`[1,3,2]`
. The maximum number of segments is 3
.

In the third test case, the only partition we can make is b<sub>1</sub>=`[5,4,3,2,1]`
. Any other partition will not satisfy the condition. Therefore, the answer is 1
.