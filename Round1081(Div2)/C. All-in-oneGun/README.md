# C. All-in-one Gun
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
You are developing a new shooter game, but since there are a lot of shooter games out there, you decide to have something unique in your game.

You have an all-in-one gun that shoots bullets in a fixed order. There are n bullets in the magazine, the i-th of which deals ai damage. The enemy starts with h health and dies when its health becomes ≤ 0.

The gun shoots one bullet per second. After firing all n bullets, it must reload, which takes k seconds. Reloading always restores the same sequence of bullets `[a`<sub>`1`</sub>`,a`<sub>`2`</sub>`,…,a`<sub>`n`</sub>`]`. You cannot reload early; you must empty the magazine first. At the start, the magazine is already full.

Before the fight begins, you may perform at most one swap: pick any indices 1 ≤ i < j ≤ n and exchange a<sub>i</sub> with a<sub>j</sub>.

Your task is to find the minimum number of seconds needed to kill the enemy, taking into account this optional single swap.
### Input

Each test contains multiple test cases. The first line contains the number of test cases t (1 ≤ t ≤ 10<sup>4</sup>). The description of the test cases follows.

The first line of each testcase contains three integer n, h and k (2 ≤ n ≤ 2⋅10<sup>5</sup>, 1 ≤ h,k ≤ 10<sup>9</sup>) — the size of magazine, health of your enemy and time required to reload the magazine.

The second line of each testcase contains n integers a<sub>1</sub>,a<sub>2</sub>,…,a<sub>n</sub> (1 ≤ a<sub>i</sub> ≤ 10<sup>9</sup>).

It is guaranteed that the sum of n does not exceed 2⋅10<sup>5</sup> over all test cases.
### Output

For each testcase, output a single integer denoting the minimum time required to kill the enemy.

### Example:
**Input:**
```
6
5 10 1
4 2 3 5 3
5 10 1
4 2 3 7 3
3 10 2
1 2 3
2 5 3
2 1
3 18 5
1 2 3
4 10 10
1 1 2 2
```
**Output:**
```
3
2
7
6
19
17
```
**Note:**  
In the first test case, you swap the bullets present at index 2 and 5. This makes array a as 4,3,3,5,2.

After 3 seconds, the health of your enemy will be 10−4−3−3=0, hence the enemy dies in 3 seconds. It can be shown that achieving time to kill less than 3 is not possible.

In the third test case, you swap bullets present at index 1 and 3. This makes array a as 3,2,1.

In 7 seconds, you shoot the entire first magazine (3 seconds) + reload a new magazine (2 seconds) + shoot the first and the second bullet from the new magazine (2 seconds).

The health of the enemy will be 10−3−2−1−3−2=−1, hence the enemy dies in 7 seconds. It can be shown that achieving time to kill less than 7 is not possible.