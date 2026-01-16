# A. Spy Detected!

**Time Limit:** 2 seconds  
**Memory Limit:** 256 megabytes  

---

## Problem Statement

You are given an array `a` consisting of `n` (**n ≥ 3**) positive integers.  
It is guaranteed that **all elements in the array are equal except one**.

For example, in the array:

```

[4, 11, 4, 4]

```

All elements are equal to `4` except `11`.

Your task is to **find and print the 1-based index** of the element that is different from the others.

---

## Input Format

- The first line contains a single integer `t`  
  *(1 ≤ t ≤ 100)* — the number of test cases.

For each test case:
- The first line contains a single integer `n`  
  *(3 ≤ n ≤ 100)* — the length of the array.
- The second line contains `n` integers  
  `a₁, a₂, …, aₙ` *(1 ≤ aᵢ ≤ 100)*.

It is guaranteed that **exactly one element is different** in each array.

---

## Output Format

For each test case, output a single integer —  
the **1-based index** of the element that does not match the others.

---

## Example

### Input
```

2
4
4 11 4 4
3
2 2 3

```

### Output
```

2
3

```

---

## Explanation

- In the first test case, `11` is the unique element → index `2`.
- In the second test case, `3` is the unique element → index `3`.

https://codeforces.com/problemset/problem/1512/A