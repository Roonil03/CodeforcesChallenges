# F. Fish Fight
## Time limit per test: 4 seconds
## Memory limit per test: 512 megabytes
### Description:
In a pond, n fish are in a line. For each fish i, its size is a<sub>i</sub>, and any fish that eats it grows by b<sub>i</sub>.

Alice picks fish x, Bob picks fish y. They alternate turns, starting with Alice's fish. On each player's turn, let their fish have size p. The fish will eat an adjacent fish with size q that satisfies p ≥ q. If more than one such fish exists, it chooses uniformly at random among them. Eating increases its size by the eaten fish's b<sub>i</sub>.

If, at the start of its turn, a fish cannot eat any adjacent fish, it is instead eaten by its neighbours (A fish at an endpoint has one neighbour; a fish in the interior has two). If Alice's fish is eaten (either by Bob's fish or by her neighboring fishes), she immediately loses. Similarly, if Bob's fish is eaten (either by Alice's fish or by his neighboring fishes), he immediately loses.

Given Alice's and Bob's chosen fishes, compute the probability that Alice wins, modulo 10<sup>9</sup> + 7.

More formally, let M = 10<sup>9</sup> + 7. The answer can be written as an irreducible fraction pq with `q ≢ 0(mod M)`. Output `pq`<sup>`−1`</sup>`mod M`, the unique integer x with 0 ≤ x < M and `xq ≡ p(mod M)`.

### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 500). The description of the test cases follows.

The first line of each testcase contains a single integer n (2 ≤ n ≤ 3000) — the number of fish in pond.

The second line contains two integers x and y (1 ≤ x,y ≤ n; x ≠ y).

The third line contains n integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ 10<sup>9</sup>).

The fourth line contains n integers b<sub>1</sub>,b<sub>2</sub>,…,b<sub>n</sub> (0 ≤ b<sub>i</sub> ≤ 10<sup>9</sup>).

It is guaranteed that the sum of n does not exceed 3000 over all test cases.

### Output

For each test case output a single integer which is the probability that Alice will win modulo 10<sup>9</sup>+7.

### Example:
**Input:**
```
6
2
1 2
1 2
1 1
2
1 2
2 1
1 1
3
2 3
1 4 4
0 1 1
5
4 2
2 6 5 5 3
1 2 1 3 2
7
3 5
1 1 1 1 1 1 1
0 0 0 0 0 0 1
10
8 3
2 5 9 3 8 4 5 6 2 7
1 3 5 2 7 3 4 2 2 3
```
**Output:**
```
0
1
500000004
750000006
375000003
687500005
```
**Note:**  
In the first test case, Alice's fish can't eat any of the adjacent fish, hence Alice always loses.

In the second test case, Alice's fish has only one adjacent fish, which is Bob's fish, Since Alice's fish size is greater than or equal to Bob's fish size, Alice always wins.

In the third test case, in the first turn, Alice's fish can eat both the adjacent fish.

Case 1: It eats fish 3 (Bob's fish). Alice wins.

Case 2: It eats fish 1. The new size of Alice's fish will be a<sub>2</sub> + b<sub>1</sub> = 4 + 0 = 4. This will lead to Bob's fish always eating Alice's fish in the second turn. Alice loses.

Both cases are equally likely, hence the probability of Alice winning is 0.5

In the fourth test case, the probability of Alice winning is 0.75
