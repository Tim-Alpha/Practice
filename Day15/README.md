# A. Police Recruits

**Time limit per test:** 1 second  
**Memory limit per test:** 256 megabytes  

The police department of your city has just started its journey. Initially, they don’t have any manpower. So, they started hiring new recruits in groups.

Meanwhile, crimes keep occurring within the city. One member of the police force can investigate only one crime during his/her lifetime.

If there is no police officer free (isn't busy with crime) during the occurrence of a crime, it will go untreated.

Given the chronological order of crime occurrences and recruit hirings, find the number of crimes which will go untreated.

---

## Input

- The first line of input contains an integer **n** (1 ≤ n ≤ 10⁵), the number of events.
- The next line contains **n** space-separated integers.
  - If the integer is **-1**, it means a crime has occurred.
  - Otherwise, the integer will be positive, representing the number of officers recruited at that time.
  - No more than **10** officers will be recruited at a time.

---

## Output

Print a single integer — the number of crimes which will go untreated.

https://codeforces.com/problemset/problem/427/A