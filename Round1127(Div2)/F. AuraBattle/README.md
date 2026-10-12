# F. Aura Battle
## Time limit per test: 2 seconds
## Memory limit per test: 256 megabytes
### Description:
*This is an interactive problem. Hacks are disabled.*

You're in an aura battle in Pachuca, Mexico. Each battle consists of $10$ rounds. The judges have a hidden integer $x$ ($1 \le x \le 10^4$), which remains fixed throughout the battle.

In round $i$, you choose two integers $a_i$ and $b_i$ ($-10^9 \le a_i, b_i \le 10^9$). You then lose

$$\min(\vert{}x - a_i\vert{}, \vert{}x - b_i\vert{})$$

aura and are told this amount before the next round begins. Your choices may depend on the results of all previous rounds. After the battle is over, your total aura loss is calculated as the sum of the aura lost in each of the $10$ rounds.

Let $m$ be the smallest integer such that there exists a strategy guaranteeing a total aura loss of at most $m$, regardless of the hidden value of $x$. Your task is to guarantee a total aura loss not exceeding the limit, which is $10^5$ **on the first test (the example)** and $m + 67$ on all other tests.

### Input
The first line of input contains a single integer $t$ ($1 \le t \le 10^4$) — the number of battles.

### Interaction
The $t$ battles are carried out one after another. In every battle, you repeat the following $10$ times:

* print two integers $a_i$ and $b_i$ ($-10^9 \le a_i, b_i \le 10^9$) — the values you play for the $i$-th round;
* read one integer $p$ — the aura you lost this round, that is, $p = \min(\vert{}x - a_i\vert{}, \vert{}x - b_i\vert{})$, where $x$ is the hidden value of the current battle.

You must play all $10$ rounds of a battle.

Nothing is printed or read between battles: right after the $10$-th round of a battle, the next battle begins, and after the $10$-th round of the last battle your program must terminate without printing anything else.

If you print a value outside $[-10^9, 10^9]$, the interactor stops answering and you will get the verdict **Wrong answer**. Make sure your program exits immediately in that case, otherwise you may get an arbitrary verdict. You will also get the verdict **Wrong answer** if the total aura you lose in some battle exceeds the limit, which is $10^5$ on the first test (the example) and $m + 67$ on all other tests.

After printing a value, do not forget to output the end of line and to flush the output. Otherwise, you will get the verdict Idleness limit exceeded. To do this, use:
* `fflush(stdout)` or `cout.flush()` in C++;
* `System.out.flush()` in Java;
* `flush(output)` in Pascal;
* `stdout.flush()` in Python;
* see the documentation for other languages.

### Examples:
**Input:**
```
1

13

4

2

2

4

5

9

6

9

13
```
**Output:**
```

3 53

36 49

47 38

1 38

44 32

35 35

31 67

33 46

49 50

20 27
```
**Explaintion:**  
In the example there is a single battle, and the hidden value $x$ is $40$.

* In the first round, $a_1 = 3$ and $b_1 = 53$, so you lose $13 = \min(|3 - 40|, |53 - 40|)$ aura.
* In the second round, $a_2 = 36$ and $b_2 = 49$, so you lose $4 = \min(|36 - 40|, |49 - 40|)$ aura.
* In the third round, the amount of aura you lose is $2$.
* In the following rounds, the amount of aura you lose is $2, 4, 5, 9, 6, 9, 13$ respectively.

The total amount of aura you lose is $67$. Note that this is an example of how the interaction works. To receive an accepted verdict, you must solve the problem without exceeding the aura limit outlined in the statement above.