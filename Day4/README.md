# A. I_love_%username%

**Time limit:** 2 seconds  
**Memory limit:** 256 megabytes  

Vasya adores sport programming. He can't write programs, but he loves to watch the progress of programming contests. Vasya even has a favorite coder, and he pays special attention to him.

One day, Vasya decided to collect the results of all contests in which his favorite coder participated and track the progress of his performance. For each contest, he wrote down a single non-negative number — the number of points his favorite coder earned in that contest. The points are listed in the chronological order in which the contests were held (no two contests ran simultaneously).

Vasya considers a coder's performance in a contest **amazing** in two situations:

1. The coder earns **strictly more points** than in all previous contests (breaks the best record).
2. The coder earns **strictly fewer points** than in all previous contests (breaks the worst record).

The first contest is **not** considered amazing.

Your task is to count the total number of amazing performances throughout the coder’s contest history.

---

## Input

- The first line contains a single integer `n` (`1 ≤ n ≤ 1000`) — the number of contests.
- The second line contains `n` space-separated non-negative integers — the points earned in each contest, given in chronological order.  
  Each value does not exceed `10000`.

---

## Output

- Print a single integer — the number of amazing performances.
