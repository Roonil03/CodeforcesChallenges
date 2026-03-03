# C1. Lost Civilization (Easy Version)
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
**This is the easy version of the problem. The difference between the versions is that in this version, you must compute a value for one sequence. You can hack only if you solved all versions of this problem.**
### Description:
Let's define an algorithm to generate a sequence of m + k
 integers as follows:
1. First, receive a sequence x
 of m
 integers as input. If k = 0
, terminate immediately and return the sequence x
.
2. Then, select any index 1 ≤ i ≤ |x|
 and insert (x<sub>i</sub> + 1)
 immediately after the element x<sub>i</sub>
.
3. If x
 contains exactly m+k
 integers, terminate and return the sequence x
. Otherwise, return to the second step.

Alice knows that this algorithm was used by an ancient civilization in order to hide their secrets safely. Alice wants to learn the knowledge that they wanted to hide, but it is not an easy job to infer the input from the output of the algorithm.

Given a sequence a
 of n
 integers, determine the length of the shortest sequence that could be given as an input for the algorithm to generate a
.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains a single integer n
 (1 ≤ n ≤ 300000
).

The second line of each test case contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ 10<sup>9</sup>
).

It is guaranteed that the sum of n
 over all test cases does not exceed 300000
.

### Output
For each test case, output the length of the shortest sequence on a separate line.

### Examples:
**Input:**
```
5
5
1 2 3 4 5
5
1 3 5 7 9
5
1 2 5 6 5
7
1 2 4 5 3 7 8
9
9 8 9 2 3 4 4 5 3
```
**Output:**
```
1
5
3
4
3
```
**Note:**  
In the first test case, the sequence `[1]`
 can generate the sequence a = `[1,2,3,4,5]`
 by the following process:

`[1]` → `[1,2]` → `[1,2,3]` → `[1,2,3,4]` → `[1,2,3,4,5]`

In the second test case, the only sequence that can generate a = `[1,3,5,7,9]`
 is a
 itself.

# C2. Lost Civilization (Hard Version)
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
**This is the hard version of the problem. The difference between the versions is that in this version, you must compute the sum of values over all subsegments. You can hack only if you solved all versions of this problem.**
### Description:
Let's define an algorithm to generate a sequence of m + k
 integers as follows:
1. First, receive a sequence x
 of m
 integers as input. If k = 0
, terminate immediately and return the sequence x
.
2. Then, select any index 1 ≤ i ≤ |x|
 and insert (x<sub>i</sub> + 1)
 immediately after the element x<sub>i</sub>
.
3. If x
 contains exactly m+k
 integers, terminate and return the sequence x
. Otherwise, return to the second step.

Alice knows that this algorithm was used by an ancient civilization in order to hide their secrets safely. Alice wants to learn the knowledge that they wanted to hide, but it is not an easy job to infer the input from the output of the algorithm.

For a sequence b
 of n
 integers, let us define `f(b)`
 as the length of the shortest sequence that could be given as an input for the algorithm to generate b
.

Given a sequence a
 of n
 integers, please compute the value of the following sum

$$\sum_{l=1}^{n} \sum_{r=l}^{n} f([a_{l}, a_{l+1}, \dots, a_{r}])$$

In other words, you must find the sum of `f(c)`
 over all subsegments<sup>∗</sup>
 c
 of a
.

<sub><sup>∗</sup>
A sequence a
 is a subsegment of a sequence b
 if a
 can be obtained from b
 by the deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end. Two subsegments are considered different if the sets of positions of the deleted elements are different. </sub>

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains a single integer n
 (1 ≤ n ≤ 300000
).

The second line of each test case contains n
 integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>
 (1 ≤ a<sub>i</sub> ≤ 10<sup>9</sup>
).

It is guaranteed that the sum of n
 over all test cases does not exceed 300000
.

### Output
For each test case, output the answer on a separate line.

### Examples:
**Input:**
```
5
5
1 2 3 4 5
5
1 3 5 7 9
5
1 2 5 6 5
7
1 2 4 5 3 7 8
9
9 8 9 2 3 4 4 5 3
```
**Output:**
```
15
35
25
60
78
```
**Note:**  
In the first test case, all subsegments of a = `[1,2,3,4,5]`
 can be generated from a sequence of length 1
.

In the second test case, all subsegments of a = `[1,3,5,7,9]`
 cannot be generated from any shorter sequence than itself.