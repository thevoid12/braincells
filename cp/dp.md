# mindset
- break the problem into sub problems
- find the smallest sub problem
- find a pattern to combine the smallest sub problem
- return the result and calculate the same for bigger sub problem
- return the final answer
  
tc: number of states(aka subproblem)* time taken to calculate each state

- pass call by reference

### 0/1 knapsack
- pick or not pick

### unbounded knapsack
- in unbounded knapsack. same thing can be used again and again just do notpick and pickand stay. (dont pick and move) pickand stay+not pick
- when you are generating combinations (where order doesn't matter), always use:
    - solve(i, amount - coins[i]) to reuse a coin
    - solve(i + 1, amount) to skip the coin
    - Do not do both "same" and "pick + i+1" — they count the same thing twice.

### dp on strings
#### longest common subsequence
- compare 2 strings
- so i and j
- match or not match
- if strings match move both the index
- else min or max (based on the question) move string A keep the string B same, move string B keep the string A same
#### pallindrome
- reverse the string and do longest common subseqeuence betyween the 2 string