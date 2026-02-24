# E. Swap to Rearrange
## Time limit per test: 5 seconds
## Memory limit per test: 512 megabytes
### Description:
You are given 2 arrays, a and b, both of length n. You can perform the following operation:

- Choose an index i (1 ≤ i ≤ n) and swap a<sub>i</sub> with b<sub>i</sub>.

You can perform the operation any number of times (possibly zero), but each index can be chosen by at most one operation. Your task is to make a a rearrangement of b after all operations, or state that it is impossible. You do not have to minimize the number of operations.

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains a single integer n (1 ≤ n ≤ 10<sup>6</sup>).

The second line of each test case contains n integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ n).

The third line of each test case contains n integers b<sub>1</sub>,b<sub>2</sub>,…,b<sub>n</sub>. (1 ≤ b<sub>i</sub> ≤ n).

It is guaranteed that the sum of n over all test cases does not exceed 10<sup>6</sup>.
### Output

For each test case, output −1 if it is impossible to make a a rearrangement of b. Otherwise, output two lines in the following format:

- In the first line, print the number of operations s (0 ≤ s ≤ n).
- In the second line of each test case, print s numbers – the indices you select in each operation in order. You should guarantee that each index is chosen at most once. 

If there are multiple possible answers, you many output any.

### Examples:
**Input:**
```
4
4
1 1 3 3
2 2 4 4
3
1 2 1
3 3 1
3
1 2 3
2 3 1
4
1 2 2 4
3 1 4 3
```
**Output:**
```
2
2 4
-1
0

2
3 4
```
**Note:**  
In the first test case, the operations performed are `swap(a`<sub>`2`</sub>`,b`<sub>`2`</sub>`)` and `swap(a`<sub>`4`</sub>`,b`<sub>`4`</sub>`)`, which will make `a=[1,2,3,4]` and `b=[2,1,4,3]`. Now it is possible to achieve a by rearranging the elements of b.

In the second test case, it can be shown that no matter what operations we do, we cannot make a a rearrangement of b.

