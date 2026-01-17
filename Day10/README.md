# A. Remove Smallest

**Time limit per test:** 1 second  
**Memory limit per test:** 256 megabytes  

You are given an array `a` consisting of `n` positive integers.


## Problem Description

In one move, you can:

- Choose two **different indices** `i` and `j` (`i ≠ j`)
- Such that the absolute difference between the elements is **at most 1**:  
  `|a[i] - a[j]| ≤ 1`
- Remove the **smaller** of the two elements  
  - If both elements are equal, you may remove **either one** (but only one)

Your task is to determine whether it is possible to reduce the array to **exactly one element** by applying the above operation **zero or more times**.


## Input

- The first line contains a single integer `t` (`1 ≤ t ≤ 1000`) — the number of test cases.
- For each test case:
  - The first line contains an integer `n` (`1 ≤ n ≤ 50`) — the length of the array.
  - The second line contains `n` integers `a₁, a₂, ..., aₙ`  
    (`1 ≤ aᵢ ≤ 100`) — the elements of the array.


## Output

For each test case, print:

- `"YES"` — if it is possible to obtain an array of exactly one element.
- `"NO"` — otherwise.

https://codeforces.com/problemset/problem/1399/A