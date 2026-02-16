# F. Parabola Independence
## Time limit per test: 3 seconds
## Memory limit per test: 512 megabytes
### Description:
You are given a set of n quadratic functions F = {f<sub>1</sub>,f<sub>2</sub>,…,f<sub>n</sub>}, where f<sub>i</sub>(x)=a<sub>i</sub>x<sup>2</sup>+b<sub>i</sub>x+c<sub>i</sub>.

Two functions f and g are called independent if f(x) ≠ g(x) for all x ∈ R.

Also, a set of functions G={g<sub>1</sub>,g<sub>2</sub>,…,g<sub>k</sub>} is called organized if the two functions g<sub>i</sub> and g<sub>j</sub> are independent for all 1 ≤ i < j ≤ |G|.

For each i=1,2,…,n, please find the size of the largest organized subset of F that contains fi as an element.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each test case contains a single integer n (1≤n≤3000).

Each of the n following lines contains three integers a<sub>i</sub>, b<sub>i</sub>, c<sub>i</sub> denoting the function f<sub>i</sub> (−10<sup>6</sup> ≤ a<sub>i</sub>, b<sub>i</sub>, c<sub>i</sub> ≤ 10<sup>6</sup>, a<sub>i</sub> ≠ 0).

It is guaranteed that the functions in one test case are pairwise distinct.

It is guaranteed that the sum of n2 over all test cases does not exceed 30002.
### Output

For each test case, output n integers s<sub>1</sub>,s<sub>2</sub>,…,s<sub>n</sub>, where s<sub>i</sub> is the size of the largest organized subset that contains f<sub>i</sub>.

### Examples:
**Input:**
```
3
4
1 2 -1
-3 0 -3
-1 4 -5
1 2 -4
5
3 0 0
1 0 -5
-3 0 0
-1 0 10
1 0 -10
5
884 -667 497
680 -973 213
23 -548 -412
826 359 -333
773 212 218
```
**Output:**
```
3 2 3 3
3 3 2 2 3
3 3 3 1 2
```
**Notes:**  
In the first test case, the functions are as follows:
- f<sub>1</sub>(x)=x<sup>2</sup>+2x−1;
- f<sub>2</sub>(x)=−3x<sup>2</sup>−3;
- f<sub>3</sub>(x)=−x<sup>2</sup>+4x−5;
- f<sub>4</sub>(x)=x<sup>2</sup>+2x−4. 

The functions' graphs are as shown below:

<img src="https://espresso.codeforces.com/168b5d6d822586244824257fcad8369941df3285.png"><br>

The largest organized subsets of F containing each function are as follows:

- {f<sub>1</sub>,f<sub>3</sub>,f<sub>4</sub>} is the largest organized subset that contains f<sub>1</sub>;
- {f<sub>1</sub>,f<sub>2</sub>} is the largest organized subset that contains f<sub>2</sub>;
- {f<sub>1</sub>,f<sub>3</sub>,f<sub>4</sub>} is the largest organized subset that contains f<sub>3</sub>;
- {f<sub>1</sub>,f<sub>3</sub>,f<sub>4</sub>} is the largest organized subset that contains f<sub>4</sub>. 

