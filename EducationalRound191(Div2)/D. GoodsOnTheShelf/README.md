# D. Goods on the Shelf
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
In a supermarket, goods of the same type are usually placed next to each other so that the shelf looks neat and it is easier for customers to find what they need.

The shelf is described by an array a
 of n
 elements, where a<sub>i</sub>
 is the type of the good at position i
.

We will say that the shelf is arranged correctly if for every two positions i
 and j
 such that 1 ≤ i < j ≤ n
 and a<sub>i</sub> = a<sub>j</sub>
, the following condition holds: for each k
 from i
 to j
, it is true that a<sub>k</sub> = a<sub>i</sub>
. In other words, goods of each type on the shelf must form one contiguous block.

You are allowed to choose two different positions at most once and swap the goods at these positions. You may also choose not to perform any swap.

Determine whether it is possible to make the shelf arranged correctly after that.

### Input
Each test contains multiple test cases. The first line contains the number of test cases t
 (1 ≤ t ≤ 10<sup>4</sup>
). The description of the test cases follows.

The first line of each test case contains an integer n
 (2 ≤ n ≤ 2 ⋅ 10<sup>5</sup>
) — the number of goods on the shelf.

The second line of each test case contains n
 integers a<sub>i</sub>
 (1 ≤ a<sub>i</sub> ≤ 10<sup>9</sup>
), where a<sub>i</sub>
 denotes the type of the good at position i
.

Additional input constraints:
- the sum of n
 over all test cases does not exceed 2 ⋅ 10<sup>5</sup>
.

### Output
For each test case, output one of the following:
- `NO`, if it is impossible to arrange the shelf correctly;
- `YES`, if it is possible to make the shelf arranged correctly with at most one swap of two goods.

You may output the answer in any case. For example, "`YeS`", "`YES`", "`NO`", "`nO`" will also be accepted.`

### Examples:
**Input:**
```
7
3
1 2 1
2
7 7
6
1 2 3 1 2 3
6
1 1 2 3 2 3
7
1 2 3 1 2 3 4
6
1 2 1 2 1 1
6
1 2 2 3 3 1
```
**Output:**
```
YES
YES
NO
YES
NO
YES
NO
```
**Explaination:**  
In the first example, you can swap the goods at positions 1
 and 2
, after which the shelf will look like this: `[2,1,1]`
.

In the second example, the shelf is already arranged correctly.

In the third example, it is impossible to arrange the shelf correctly with one swap.

In the sixth example, you can swap the goods at positions 1
 and 4
, after which the shelf will be arranged correctly.