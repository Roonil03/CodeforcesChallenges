# E. Cake Trial
## Time limit per test: 2 seconds
## Memory limit per test: 512 megabytes
### Description:
<i><div align=right><sub>
The cake is a lie.

---

 —Portal
</sub></div></i>

GLaDOS, the artificial intelligence of the Aperture Science Enrichment Center, decided to conduct another test for Chell. This time, the test is about cakes.

GLaDOS arranged $n$ cakes in a row. Each cake can be either real ($\texttt{T}$) or fake ($\texttt{F}$). Chell must guess which cakes are real and which are fake. Chell has a unique ability: she can determine exactly whether a cake is real just by looking at it. However, in her answer, she must satisfy GLaDOS's strange condition: all fake cakes in Chell's answer must form one contiguous subsegment (possibly empty).

Initially, GLaDOS prepared some arrangement of cakes, but some cakes have not been placed yet. You are given a string $s$ of length $n$ describing the current situation:

* $s_i = \texttt{T}$, if the $i$-th cake is real;
* $s_i = \texttt{F}$, if the $i$-th cake is fake;
* $s_i = \texttt{N}$, if there is no cake at the $i$-th position yet, and GLaDOS may still place either a real or a fake cake there at her discretion.

After GLaDOS finishes the arrangement (replaces all $\texttt{N}$ with $\texttt{T}$ or $\texttt{F}$), Chell will see the final arrangement and will know exactly for each cake whether it is real or fake. Chell will choose a contiguous segment $[l, r]$ ($1 \le l \le r \le n$) that she will declare to be the set of fake cakes (or she may declare all cakes to be real). All cakes outside this segment are considered real. Chell wants her answer to be as close to the truth as possible, so she will choose the segment to minimize the number of mistakes.

GLaDOS, being cunning, wants to make Chell's life harder. She wants to place the remaining cakes (replace all $\texttt{N}$ with $\texttt{T}$ or $\texttt{F}$) so that the number of mistakes Chell is forced to make, under her optimal choice of segment, is as large as possible.

Help GLaDOS determine the maximum number of mistakes she can guarantee.

### Input

Each test contains multiple test cases. The first line contains the number of test cases $t$ ($1 \le t \le 10^4$). The description of the test cases follows.

The first line of each test case contains one integer $n$ ($1 \le n \le 500$) — the number of cakes.

The second line of each test case contains a string $s$ of length $n$ consisting of the characters $\texttt{T}$, $\texttt{F}$, or $\texttt{N}$ — the current arrangement of cakes.

It is guaranteed that the sum of $n^3$ over all test cases does not exceed $500^3$.

### Output

For each test case, output one integer — the maximum number of mistakes that GLaDOS can guarantee.

### Examples:
**Input:**
```
10
4
FTFF
5
TNFTT
6
TFTTTN
6
TNNFTF
7
TNFNTNF
6
NNFFNN
7
TNTFNTN
1
N
5
NNNNN
10
NNNTTNNNFN
```
**Output:**
```
1
0
1
2
2
2
2
0
2
3
```
**Explaination:**  
In the first test case, $s = \texttt{FTFF}$, all cakes are already placed. Chell will choose the segment $[3, 4]$, resulting in the answer $\texttt{TTFF}$, and she will have 1 error.

In the second test case, $s = \texttt{TNFTT}$, 1 cake is not placed. For any replacement of $\texttt{N}$ with $\texttt{T}$ or $\texttt{F}$, the fake cakes form a continuous segment that Chell can choose, so the answer is 0.

In the third test case, $s = \texttt{TFTTTN}$, GLaDOS will replace $\texttt{N}$ with $\texttt{F}$, obtaining $\texttt{TFTTTF}$. It can be shown that the answer is at least 1. If Chell chooses the segment $[2, 2]$, i.e., gives the answer $\texttt{TFTTTT}$, she will make exactly 1 error.

In the fourth test case, $s = \texttt{TNNFTF}$, GLaDOS will arrange the cakes as follows: $\texttt{TFTFTF}$. Chell will choose the segment $[2, 4]$ (answer: $\texttt{TFFF TT}$), making 2 errors.