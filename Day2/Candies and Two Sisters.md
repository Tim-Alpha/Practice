# 🍬 A. Candies and Two Sisters

## 📌 Problem Statement

There are two sisters, **Alice** and **Betty**, and you have **n candies**.
You want to distribute the candies such that:

* Alice gets **a > 0** candies
* Betty gets **b > 0** candies
* Both receive an **integer** number of candies
* **Alice gets more candies than Betty** (`a > b`)
* All candies are distributed: `a + b = n`

Candies are **indistinguishable**.

Your task is to compute the **number of valid distributions** for each test case.

---

## 🧠 Key Insight

Since:

* `a + b = n`
* `a > b`
* `a, b > 0`

The number of valid ways is:

* If `n ≤ 2` → **0 ways**
* Otherwise → **(n − 1) // 2**

This works because:

* We count how many valid values `b` can take such that `b < n - b`
* Integer division handles both even and odd `n`

https://codeforces.com/problemset/problem/1335/A