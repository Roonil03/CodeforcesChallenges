# C. Arrange the Numbers in a Circle
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
You have cards with numbers: c<sub>1</sub>
 cards with the number 1
, c<sub>2</sub>
 cards with the number 2
, ..., c<sub>n</sub>
 cards with the number n
.

You must take at least three cards from the ones you have and arrange them in a circle so that the following condition holds:
- in every triple of consecutive cards, there are at least two equal cards.

Formally, if the numbers on the chosen cards are a<sub>0</sub>, a<sub>1</sub>, …, a<sub>k−1</sub>
 in the order they are arranged in a circle, then the following condition must hold:
- for every i
 from 0
 to k−1
, there are at least two equal numbers among a<sub>i</sub>, a<sub>(i + 1) mod k</sub>, a<sub>(i + 2) mod k</sub>
.
What is the maximum number of cards you can arrange?

### Input
The first line contains one integer t
 (1 ≤ t ≤ 10<sup>4</sup>
) — the number of test cases.

Each test case consists of two lines:
- the first line contains one integer n
 (1 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
);
- the second line contains n
 integers c<sub>1</sub>, c<sub>2</sub>, …, c<sub>n</sub>
 (1 ≤ c<sub>1</sub> ≤ c<sub>2</sub> ≤ ⋯ ≤ c<sub>n</sub> ≤ 10<sup>9</sup>
).

Additional constraint on the input: the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output one integer — the maximum number of cards you can arrange. If it is impossible to choose three or more cards and arrange them in a circle so that the condition holds, output 0
.

### Examples:
**Input:**
```
11
4
1 1 1 3
3
2 3 4
6
1 1 1 1 3 4
3
1000000000 1000000000 1000000000
3
1 1 2
1
2
2
2 2
3
1 1 1
4
1 1 2 2
3
1 1 4
9
1 1 1 1 1 1 1 1 7
```
**Output:**
```
4
9
8
3000000000
3
0
4
0
4
6
10
```
**Explaination:**  
In the first example, you can choose and arrange the cards as follows: `[4,2,4,4]`
. You cannot arrange more than 4
 cards. For example, the arrangement `[2,4,4,1,4]`
 does not work, because the cards must be arranged in a circle, and the condition also applies to the triple consisting of the second-to-last, last, and first cards.

In the second example, you can arrange all the cards: `[1,1,2,2,2,3,3,3,3]`
.

In the third example, you can choose and arrange the cards as follows: `[5,5,6,6,3,6,6,5]`
.

