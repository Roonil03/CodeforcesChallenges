# A. Parkour Design
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
Today, Alex wants to build a parkour course for Steve to train his parkour skills on. A parkour course is a sequence p<sub>0</sub>→p<sub>1</sub>→…→p<sub>k</sub> of integer coordinates on the plane. Here, a contiguous pair of coordinates is called a move, denoted as p<sub>i−1</sub>→p<sub>i</sub>.

Alex knows that Steve can only perform the following types of moves:

- (x<sub>i</sub>,y<sub>i</sub>)→(x<sub>i+2</sub>,y<sub>i+1</sub>);
- (x<sub>i</sub>,y<sub>i</sub>)→(x<sub>i+3</sub>,y<sub>i</sub>);
- (x<sub>i</sub>,y<sub>i</sub>)→(x<sub>i+4</sub>,y<sub>i−1</sub>). 

Note that Steve will not perform any other type of moves. For example, Steve can perform `(0,0)→(2,1)` and `(2,1)→(5,1)`, but will never perform moves such as `(2,1)→(3,2)`, `(3,0)→(5,−1)`, or `(4,−1)→(6,−1)` (even though they may look very easy).

You are given an integer coordinate `(x,y)` on the plane.

Please determine if it is possible to make a parkour course q0,q1,…,qk that satisfies the following conditions:

- q<sub>0</sub>=`(0,0)`;
- q<sub>k</sub>=`(x,y)`;
- The parkour course only consists of moves that Steve can perform. 

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>3</sup>). The description of the test cases follows.

The only line of each test case contains two integers x and y (1 ≤ x ≤ 10<sup>9</sup>, −10<sup>8</sup> ≤ y ≤ 10<sup>8</sup>).
### Output

If it is possible to make a parkour course that satisfies the conditions, output "YES" on a separate line.

If it is impossible to make a parkour course that satisfies the conditions, output "NO" on a separate line.

You can output the answer in any case. For example, the strings "yEs", "yes", and "Yes" will also be recognized as positive responses. 

### Example:
**Input:**
```
11
2 1
3 0
4 -1
4 1
14 1
1 -4
3 -1
2 10
24 -1
24 -3
8 4
```
**Output:**
```
YES
YES
YES
NO
YES
NO
NO
NO
NO
YES
YES
```
**Note:**  
For the fifth test case, the parkour course must start from `(0,0)` and end on `(14,1)`.

This can be achieved by the following parkour course.

`(0,0)→(4,−1)→(7,−1)→(9,0)→(12,0)→(14,1)`

Note that the following parkour course does not satisfy the third condition stated above due to the moves highlighted in red.

(0,0)→**(4,−1)→(6,−1)**→(8,0)→**(11,0)→(14,1)**

