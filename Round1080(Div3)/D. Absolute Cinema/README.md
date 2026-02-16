# D. Absolute Cinema
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
There is a hidden sequence a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> of n integers (n ≥ 2). It is guaranteed that |a<sub>i</sub>| ≤ 1000 for all 1 ≤ i ≤ n.

Let's define a function f(x) as follows:
$$f(x) = \sum_{i=1}^{n} a_i \cdot |i - x|$$

Given n values f(1),f(2),…,f(n), please determine the values of a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub>.

It is guaranteed that the values a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> can be determined uniquely.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains a single integer n (2 ≤ n ≤ 300000).

The second line of each test case contains n integers f(1),f(2),…,f(n) (−10<sup>14</sup> ≤ f(i) ≤ 10<sup>14</sup>).

It is guaranteed that the sum of n over all test cases does not exceed 300000.
### Output

For each test case, output n integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> on a separate line (|a<sub>i</sub>| ≤ 1000).

It is guaranteed that the values a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> can be determined uniquely.

### Examples:
**Input:**
```
4
4
17 9 9 13
6
-37 -32 -15 4 27 42
5
-26 -32 -24 -4 2
2
420 -69
```
**Output:**
```
1 4 2 3
3 6 1 2 -4 -7
-6 7 6 -7 -6
-69 420
```
**Note:**
In the first test case, the hidden sequence is a=`[1,4,2,3]`.

The values f(1),f(2),…,f(n) are as follows:
- `f(1)=1⋅|1−1|+4⋅|2−1|+2⋅|3−1|+3⋅|4−1|=0+4+4+9=17`;
- `f(2)=1⋅|1−2|+4⋅|2−2|+2⋅|3−2|+3⋅|4−2|=1+0+2+6=9`;
- `f(3)=1⋅|1−3|+4⋅|2−3|+2⋅|3−3|+3⋅|4−3|=2+4+0+3=9`;
- `f(4)=1⋅|1−4|+4⋅|2−4|+2⋅|3−4|+3⋅|4−4|=3+8+2+0=13`.