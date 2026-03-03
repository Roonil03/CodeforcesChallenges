# A. Sieve of Erato67henes
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
You are given n positive integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>.

Please determine if it is possible to select any number of elements in a, so that their product is 67.

Note that you may not select zero elements, as the product of zero elements is not defined in this problem.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains a single integer n (1≤n≤5).

The second line of each test case contains n positive integers a1,a2,…,an (1≤a<sub>i</sub>≤67).
### Output

If it is possible to select elements so that their product is 67, output "YES" on one line. Otherwise, output "NO" on one line.

You can output the answer in any case. For example, the strings "yEs", "yes", and "Yes" will also be recognized as positive responses.

### Example:
**Input:**
```
2
5
1 7 6 7 67
5
1 3 5 7 8
```
**Output:**
```
YES
NO
```
**Note**  
In the first test case, you can select a<sub>1</sub> and a<sub>5</sub> to get a<sub>1</sub> ⋅ a<sub>5</sub> = 1 ⋅ 67 = 67.  
In the second test case, it is impossible to select any number of elements so that their product is 67.
