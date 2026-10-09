# G. Copper Squander
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
K1o0n accepted an alternative ending and got a job in the programming department. As is known, all programmers know how to dig trenches for cables, so his very first assignment was wiring between the university buildings — with the happy right to sell off the old stuff on his own.

There are $n$ buildings at the university, and cables are laid between some pairs of them — $m$ cables in total. Data travels in both directions along the $i$-th cable; it connects buildings $u_i$ and $v_i$, and it contains copper worth $d_i$ coins. The network was laid long before K1o0n, and nobody cared about money: somewhere an extra jumper was installed, somewhere a more expensive cable was used, and some buildings were never reached at all.

The network can be rebuilt in exactly the following order:

1. **Remove** any subset of cables. The copper from the removed $i$-th cable is sent to a recycling center and brings $d_i$ coins; there is no longer any connection through that cable.
2. **Lay** any number of new cables between any two distinct buildings. K1o0n digs the trenches himself, but the cable is installed by a construction crew, and its appetite grows with every request: the first new cable will cost $x$ coins, the second — $2x$, the third — $3x$, and so on.

After that, the network must be connected: data must be able to travel from any building to any other building. Note that initially the network may be disconnected.

Profit is the revenue from sold copper minus the installation costs.

You are given $q$ scenarios, in the $j$-th scenario the installation cost increment is $x_j$. For each scenario, find the maximum possible profit.

### Input
The first line contains one integer $t$ ($1 \le t \le 10^4$) — the number of testcases.

The first line of each testcase contains three integers $n$, $m$, and $q$ ($2 \le n \le 10^5$, $0 \le m \le 2 \cdot 10^5$, $1 \le q \le 2 \cdot 10^5$) — the number of buildings, cables, and scenarios.

The next $m$ lines contain three integers each: $u_i$, $v_i$, and $d_i$ ($1 \le u_i, v_i \le n$, $u_i \ne v_i$, $1 \le d_i \le 5 \cdot 10^8$) — the endpoints of the $i$-th cable and the value of the copper in it. A cable cannot connect a building to itself, and no pair of buildings is connected by two cables.

The last line of the testcase contains $q$ integers $x_1, x_2, \dots, x_q$ ($1 \le x_j \le 5 \cdot 10^8$) — the installation cost increments in the scenarios.

It is guaranteed that the sum of $n$ over all testcases does not exceed $10^5$, the sum of $m$ does not exceed $2 \cdot 10^5$, and the sum of $q$ does not exceed $2 \cdot 10^5$.

### Output
For each scenario, output one integer — the maximum profit. Print the answers for the scenarios of one testcase in the same order as the scenarios are given, on one line.

### Examples:
**Input:**
```
3
4 4 1
1 2 1
2 3 1
3 4 1
1 4 10
1
5 4 1
1 2 5
1 3 5
1 4 5
1 5 5
1
5 4 3
1 2 5
2 3 5
3 4 5
1 3 1
1 3 10
```
**Output:**
```
10 
10 
9 2 -5 
```