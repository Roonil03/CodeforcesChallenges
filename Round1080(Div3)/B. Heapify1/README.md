# B. Heapify 1
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are given a permutation a of length n<sup>∗</sup>.

You can perform the following operation any number of times (possibly zero):
- Select an index i (1 ≤ i ≤ n / 2), and swap a<sub>i</sub> and a<sub>2i</sub>. 

For example, when a = `[1,4,2,3,5]`, you can swap a<sub>2</sub> and a<sub>4</sub> to make it `[1,3,2,4,5]`, but you cannot swap a<sub>2</sub> and a<sub>3</sub>.

Please determine if the sequence a can be sorted in increasing order.

<sup>∗</sup><sub>A permutation of length n is an array consisting of n distinct integers from 1 to n in arbitrary order. For example, `[2,3,1,5,4]` is a permutation, but `[1,2,2]` is not a permutation (2 appears twice in the array), and `[1,3,4]` is also not a permutation (n=3 but there is 4 in the array). </sub>

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains a single integer n (1 ≤ n ≤ 2⋅10<sup>5</sup>).

The second line of each test case contains n distinct integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ n).

It is guaranteed that the sum of n over all test cases does not exceed 2⋅10<sup>5</sup>.
### Output

If a can be sorted in increasing order, output "YES" on a separate line. Otherwise, output "NO" on a separate line.

You can output the answer in any case. For example, the strings "yEs", "yes", and "Yes" will also be recognized as positive responses.

### Examples:
**Input:**
```
2
5
1 4 3 2 5
5
1 4 2 3 5
```
**Output:**
```
YES
NO
```
**Note**  
In the first test case, a is `[1,4,3,2,5]`. You can sort a in increasing order by swapping a<sub>2</sub> and a<sub>4</sub>. Therefore, the answer is "YES".

In the second test case, a is `[1,4,2,3,5]`. It is impossible to sort a in increasing order. Therefore, the answer is "NO".
