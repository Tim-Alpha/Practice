# A. Game with Integers

**Time limit per test:** 1 second  
**Memory limit per test:** 256 megabytes  

Vanya and Vova are playing a game. Players are given an integer `n`.  
On their turn, a player can either:

- Add `1` to the current integer, or  
- Subtract `1` from the current integer.  

The players take turns, and **Vanya starts first**.

- If after **Vanya's move** the integer becomes divisible by `3`, then **Vanya wins**.
- If **10 moves** have passed and Vanya has not won, then **Vova wins**.

You need to determine the winner if both players play optimally.

---

## Input

- The first line contains an integer `t` (`1 ≤ t ≤ 100`) — the number of test cases.
- Each test case contains a single integer `n` (`1 ≤ n ≤ 1000`).

---

## Output

For each test case:

- Print `"First"` if Vanya wins.
- Print `"Second"` if Vova wins.


https://codeforces.com/problemset/problem/1899/A