# A. In Search of Convenience
## Time limit per test: 1 second
## Memory limit per test: 256 megabytes
### Description:
K1o0n got a router and placed it at point $(x_0, y_0)$; we will consider the apartment layout as a coordinate plane, and the floor is tiled, so the furniture can stand only at lattice points with integer coordinates.

The internet spreads exactly $R$ meters around the router. K1o0n wants to move his computer as far away from it as possible — but still so that the internet is available. Therefore, the desk with the computer must be placed exactly on the reception boundary, at a distance of $R$ from the router. For example, if the router is at point $(5, 5)$ and $R = 5$, then the desk can be placed at point $(2, 1)$, because $(5 - 2)^2 + (5 - 1)^2 = 5^2$.

Find any point with integer coordinates that is exactly $R$ away from $(x_0, y_0)$.

Recall that the distance from the point $(x_0, y_0)$ to the point $(x, y)$ is $\sqrt{(x_0 - x)^2 + (y_0 - y)^2}$.

### Input
The first line contains an integer $t$ ($1 \le t \le 10^4$) — the number of testcases.

The only line of each testcase contains three integers $x_0$, $y_0$, and $R$ ($-10 \le x_0, y_0 \le 10$, $1 \le R \le 25$) — the coordinates of the router and the coverage radius.

### Output
For each testcase, output two integers $x$ and $y$ — the coordinates of the desk.

If there are several suitable points, output any of them.

### Examples:
**Input:**
```
3
0 0 1
5 5 5
10 10 13
```
**Output:**
```
0 1
2 1
-2 5
```