# H. Codeforces Heuristic Contest 001
## Time limit per test: 4 seconds
## Memory limit per test: 1024 megabytes
### Description:
There is a grid of 3n × 3n points, consisting of all integer points (x,y) such that 1 ≤ x,y ≤ 3n.

Find a largest set of triangles satisfying the following conditions:

- Each triangle has its vertices on exactly three points on the grid.
- Each triangle has an area of exactly 12. Note that they do not have to be right triangles.
- No two triangles share a common intersection point, including their vertices. 

If there exist multiple largest such sets of triangles, you may output any of them.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 30). The description of the test cases follows.

The only line of each test case contains a single integer n (1 ≤ n ≤ 166).

It is guaranteed that the sum of n<sup>2</sup> over all test cases does not exceed 166<sup>2</sup>.
### Output

Output the maximum size m of the set of triangles on one line (0 ≤ m ≤ 3n<sup>2</sup>).

Then, output m lines in the following format:
- x<sub>i,1</sub>y<sub>i,1</sub>x<sub>i,2</sub>y<sub>i,2</sub>x<sub>i,3</sub>y<sub>i,3</sub>: The vertices of the i-th triangle are (x<sub>i,1</sub>, y<sub>i,1</sub>), (x<sub>i,2</sub>, y<sub>i,2</sub>), (x<sub>i,3</sub>, y<sub>i,3</sub>). 

You may output the vertices of one triangle in any order (clockwise or counterclockwise).

Your output will be accepted if it satisfies all conditions and the maximum size given is correct.

### Examples:
**Input:**
```
2
1
2
```
**Output:**
```
2
1 1 1 2 2 1
2 3 3 2 3 3
12
1 1 1 2 2 1
2 2 3 2 3 1
1 3 1 4 2 3
2 4 3 4 3 3
1 5 1 6 2 5
2 6 3 6 3 5
4 1 4 2 5 1
5 2 6 2 6 1
4 3 4 4 5 3
5 4 6 4 6 3
4 5 4 6 5 5
5 6 6 6 6 5
```
**Note:**  
In the first test case, the example output has 2 triangles as shown in the following image:  

<img src="https://espresso.codeforces.com/91685b30b8b9b30b68d7cabc8aa5bdd91434b214.png">

In the second test case, the example output has 12 triangles as shown in the image on the left:
	
<img src="https://espresso.codeforces.com/7077bbe8222b559a4f66c9fb949479f6eff1fdcd.png"><img src="https://espresso.codeforces.com/a0141e7b204e64817c88b9b22b7b54d7bcf629ac.png">

As the triangles are not required to be right triangles, the set of triangles shown in the image on the right will also be considered valid.