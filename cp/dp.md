# mindset
- break the problem into sub problems
- find the smallest sub problem
- find a pattern to combine the smallest sub problem
- return the result and calculate the same for bigger sub problem
- return the final answer
  
tc: number of states(aka subproblem)* time taken to calculate each state

- pass call by reference

- in unbounded knapsack. same thing can be used again and again just do notpick and pickand stay. (dont pick and move) pickand stay+not pick
- when you are generating combinations (where order doesn't matter), always use:
    - solve(i, amount - coins[i]) to reuse a coin
    - solve(i + 1, amount) to skip the coin
    - Do not do both "same" and "pick + i+1" — they count the same thing twice.