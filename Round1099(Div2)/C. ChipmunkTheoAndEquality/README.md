# C. Chipmunk Theo and Equality
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
While exploring the depths of the Internet, Theodore the chipmunk found a very interesting sequence of positive integers and decided to play with it.

In one operation, he chooses an element of the sequence and performs the following action:
- If the chosen element is even, he divides it by 2.
- If the chosen element is odd, he increases it by 1.

Theo really loves equality, so he wants to make all numbers in the sequence equal (otherwise, some numbers might feel offended). Since he needs to plan his lunch time, help him determine the minimum number of operations required to make all numbers equal.

### Input:
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains a single integer n
 (1 ≤ n ≤ 10<sup>5</sup>
) — the length of the sequence.

The second line of each test case contains n
 integers a<sub>1</sub>, a<sub>2</sub>, …, a<sub>n</sub>
 (1≤ai≤109
) — the elements of the sequence.

It is guaranteed that the sum of n
 over all test cases does not exceed 10<sup>5</sup>
.

### Output
For each test case, output a single integer — the minimum number of operations Theo needs to make all numbers in the sequence equal.

### Examples:
**Sample Input:**
```
5
3
3 2 4
7
3 6 7 16 8 8 7
3
1 4 2
5
10 10 10 10 10
6
1 1 3 1 1 1
```
**Sample Output:**
```
3
11
2
0
3
```
**Explaination:**  
In the first test case, we have a sequence: `[3,2,4]`

One possible sequence of operations:  
[**3**,2,4]→[**4**,2,4]→[2,2,**4**]→[2,2,**2**]  
(Operations are performed on numbers shown in bold.)  

In the second test case, the sequence is: [3,6,7,16,8,8,7]

Possible operations:

- Perform the operation once on the 1st element: 3→4 
- Perform the operation twice on the 2nd element: 6→3→4
- Perform the operation twice on the 4th element: 7→8→4
- Perform the operation twice on the 5th element: 16→8→4
- Perform the operation once on the 6th element: 8→4
- Perform the operation once on the 7th element: 8→4
- Perform the operation twice on the 8th element: 7→8→4

After these 11
 operations, all elements become 4
.


