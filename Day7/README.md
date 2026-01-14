# A. Plus or Minus

**Time Limit per Test:** 1 second  
**Memory Limit per Test:** 256 megabytes  

## Problem Statement

You are given three integers **a**, **b**, and **c** such that **exactly one** of the following equations is true:

- `a + b = c`
- `a - b = c`

Your task is to determine which equation is correct.

- Output `+` if `a + b = c` is true  
- Output `-` otherwise

---

## Input

- The first line contains a single integer **t**  
  `(1 ≤ t ≤ 162)` — the number of test cases.

- Each test case consists of three integers **a**, **b**, and **c**  
  - `(1 ≤ a, b ≤ 9)`
  - `(-8 ≤ c ≤ 18)`

**Note:** The input is guaranteed such that exactly one of the two equations is true for each test case.

---

## Output

For each test case, output either `+` or `-` on a new line, representing the correct equation.


https://codeforces.com/problemset/problem/1807/A