
# Internals of WHERE vs IN Clauses in SQLite: A Simple Guide

- When working with SQLite, two common ways to filter data are the WHERE clause and the IN clause.
- In this wiki, we will look at what happens when we use **WHERE** and **IN** clauses in SQLite.

## Terminologies

Here are some key terms to get familiar with as they'll help you understand the comparisons I'll be discussing throughout this wiki. I'll be using these terms quite a bit, so knowing them might be helpful. Feel free to skip ahead.

#### Full Table Scan:

- This algorithm is called a full table scan since the entire content of the table must be read and examined in order to find the one row of interest.
- A full table scan occurs when SQLite reads every row in a table to find the rows that match the query conditions. 
- This happens when:
    - There is no index available for the columns used in the query.
    - The **query optimizer** determines that using an index would be less efficient than scanning the entire table.
    > Note: Query optimizer is a broader term. Query planning is a related concept.

#### Indexed Scan:

- SQLite uses B-trees to implement indexes. Each index is a separate B-tree structure where the keys are the indexed column values, and the values are pointers (row IDs) to the corresponding rows in the table.
- **During an indexed scan:**
  - SQLite traverses the index B-tree to find the keys that match the query conditions.
  - It retrieves the row IDs (pointers) associated with the matching keys.
  - It uses the row IDs to fetch the corresponding rows from the table. (We have control over row IDs, so it is still possible to delete row IDs and if so, then the primary key is used)
- **Composite Index:**
  - Indexing more than one column together
- **Covering Index:**
  - If an index includes all columns needed for a query, SQLite doesn't need to look up the actual table.
  - Example: Index on (EmpID, Name)  
    ```sql
    SELECT EmpID, Name FROM auditcue 
    WHERE EmpID = 'EMP001';
    ```

#### Query Planner:

- The best feature of SQLite is that it is a **declarative language**, not a **procedural language**.
- That is, we tell the system what you want to compute, not how to compute it.
- The how to compute part is taken care of by the query planner.
- The query planner is the component responsible for generating an execution plan for a SQL query.

#### Key Steps in Query Planning and Execution:

1. **Parse the Query**:
   - SQLite parses the SQL statement into an internal data structure called a **parse tree**.
   - This step ensures the query is syntactically correct and converts it into a form that the query optimizer can work with.

2. **Analyze the Query**:
   - The query optimizer analyzes the parse tree to understand the query's structure (e.g., SELECT, WHERE, JOIN, etc.).
   - It identifies which tables and columns are involved and checks for the presence of **indexes** on columns used in conditions (e.g., WHERE, JOIN, ORDER BY, GROUP BY).

3. **Generate and Evaluate Execution Plans**:
   - The **query optimizer** considers different ways to execute the query (e.g., using an index scan, full table scan, or a combination of both).
   - It estimates the **cost** of each plan based on factors like:
     - The size of the tables.
     - The selectivity of the conditions (how many rows match the WHERE clause).
     - The presence and usefulness of indexes.
   - The optimizer picks the plan with the **lowest estimated cost**.

4. **Execute the Plan**:
   - The query planner generates the **bytecode** for the chosen execution plan. This bytecode is executed by SQLite's virtual machine (VM).
   - The VM performs operations like:
     - Scanning tables or indexes.
     - Applying filters (WHERE clause).
     - Joining tables (if applicable).
     - Sorting or grouping results (if required).
   - The results are returned to the user.

# WHERE Clause:

- The WHERE clause filters rows based on specific conditions.

#### AND Connected WHERE:

- When you use **AND** to combine conditions, SQLite's query planner looks for the most efficient way to satisfy all conditions, often leveraging indexes to minimize row scanning.
- If there's a composite index on (Acue_EmpID, Acue_Name), SQLite uses it to quickly find rows where both conditions are true. This will avoid final lookup on our table to get the output.
- If no composite index exists, SQLite might use two separate indexes (if available) and combine results, but this is less efficient.
- If no index exists, a full table search is performed which is the least efficient.

#### OR Connected WHERE:

```sql
--Index B
--Index C
SELECT * FROM tableA WHERE B > 10 OR C > 20
```

- Assume that we have indexed B and C separately.
- **OR** conditions are trickier for the query planner. SQLite often can't use a single index to satisfy multiple OR terms, but it has optimizations to handle them.
- Suppose we have two indexes on the OR condition, SQLite might perform two separate index searches and combine the results (using a UNION internally).
- This approach is called **OR-by-UNION optimization**:
    - Split the query into two parts: B > 10 and C > 20
    - Use separate indexes for each part.
    - Combine the results and remove duplicates (if needed).
- If no indexes exist, it defaults to a full table scan.

### Mixed OR AND Connected WHERE

```sql
SELECT * FROM tableA WHERE (B > 10 AND D > 100) OR C > 20
```

- When mixing AND and OR, grouping using **parentheses** and index selection become critical.
- By grouping parentheses, we force the query planner to evaluate what's inside the parentheses first and then the others.
- The query inside parentheses can use any of the **OR** or **AND** index-based optimizations.
- In the example above, (B > 10 AND D > 100) is evaluated together, and if they are properly indexed, this evaluation would be much faster than without parentheses and giving all control to the query optimizer.

#### Key Takeaway from WHERE Clause:

- So what did we take away:
  1. Always index
  2. SQLite prefers composite indexes to resolve multiple AND terms efficiently.
  3. SQLite uses multiple indexes (if available) and combines results via UNION.
  4. Parentheses matter a lot.

# IN Clause

- The IN clause checks if a value matches any value in a list or a subquery result.

Example:
```sql
SELECT * FROM auditcue 
WHERE empID IN (1,2,3);
```

- Gives us all the results which match with empID 1, empID 2, empID 3
- **IN query** has two variations/forms:
  - One is similar to the example above (**list values**)
  - The other is used in a **subquery** 
    Example: 
    ```sql
        WHERE EMPID IN (SELECT EMPID FROM auditcue WHERE location = 'Chennai')
    ```
### Internals of IN Clause

- SQLite converts `IN (v1, v2, ...)` into a sequence of equality comparisons (OR conditions) with WHERE clause
```sql
SELECT * FROM t WHERE x IN (1, 2, 3);
  -- is rewritten as:
  SELECT * FROM t WHERE x = 1 OR x = 2 OR x = 3;
```
- If an index exists on x, SQLite performs index lookups for each value in the list. Without an index, a full table scan occurs, checking each row against the list.

### Subqueries
- **Subquery Materialization**:  
   - **IN** with subqueries materializes results into an **unindexed temp table**, leading to slower linear scans.  
   - Flattening the subquery into a **JOIN** avoids this overhead.  

- **Unoptimized**: Materializes results into a temp table (slow for large datasets):

  ```sql
  SELECT * FROM employees 
  WHERE department IN (SELECT department FROM departments WHERE location = 'Chennai');
  ```

- **Optimized**: Flattened into a `JOIN` for index usage:

  ```sql
  SELECT employees.* 
  FROM employees 
  JOIN departments ON employees.department = departments.department 
  WHERE departments.location = 'Chennai';
  ```

## Execution Time Comparison

| **Scenario**                     | **`WHERE` with `OR` Execution Time** | **`IN` Clause Execution Time** | **Explanation**                                                                                           |
|-----------------------------------|--------------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
| **Small Value List (Indexed)**    | Fast                                 | Fast                            | Both use **index lookups** for each value. **IN** is internally rewritten as **OR**, resulting in identical runtime. |
| **Small Value List (No Index)**   | Slow (Full scan + **OR** checks)       | Slow (Same as **OR**)             | Full table scan required. **IN** adds negligible overhead as it compiles to equivalent bytecode.            |
| **Large Value List (Indexed)**    | Moderate                             | Moderate                        | Index lookups for each value. **IN** may parse faster, but execution time is similar.                       |
| **Large Value List (No Index)**   | Very Slow (Full scan)                | Very Slow (Same as **OR**)        | Both scan the entire table. **IN** is marginally slower due to temporary list parsing.                      |
| **Subquery (Small Result Set)**   | N/A                                  | Moderate                        | **IN** materializes the subquery into an unindexed temp table, leading to linear scans.                     |
| **Subquery (Large Result Set)**   | N/A                                  | Very Slow                       | Large temp table with linear scans. `WHERE` with `JOIN` (flattened subquery) is faster.                   |
| **Subquery (Flattened to `JOIN`)** | Fast (Uses indexes)                  | Fast (Same as `JOIN`)           | Optimizer converts **IN** to a `JOIN`, leveraging indexes on both tables.                                   |

### Key Takeaway from IN Clause

- **Use IN** for value lists (cleaner syntax).  
- **Use JOIN** instead of **IN** with subqueries (avoids temp table scans).
- As usual, INDEX!!!!
- For get api's we can use **IN** clause in sqlite as it doesnt make any difference 

## References I Used to Draft This Wiki

- [Query Planning](https://www.sqlite.org/queryplanner.html)
- [WHERE Clause Filtering](https://www.sqlite.org/lang_select.html#whereclause)
- [SQLite Query Optimizer Overview](https://www.sqlite.org/optoverview.html)
- [SQLite Bytecode Engine](https://www.sqlite.org/opcode.html)  
- [IN Operator Documentation](https://www.sqlite.org/lang_expr.html#in_op)  
- [Subquery Flattening](https://www.sqlite.org/compile.html#subquery_flattening)  
- [Explain Query Plan](https://www.sqlite.org/eqp.html)

Use **EXPLAIN QUERY PLAN** to analyze specific queries and validate optimizations.
