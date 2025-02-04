To excel in competitive programming with C++, mastering the Standard Template Library (STL) is crucial.

---

### **1. Containers**
| STL Component       | Operation               | Time Complexity      | Example/Use Case                          |
|---------------------|-------------------------|----------------------|--------------------------------------------|
| **`vector`**        | `push_back()/pop_back()`| O(1)                 | `v.push_back(5);` (dynamic arrays)         |
|                     | `insert()/erase()`      | O(n)                 | `v.insert(v.begin(), 10);` (costly)        |
|                     | `operator[]`            | O(1)                 | `int x = v[2];` (random access)            |
| **`deque`**         | `push_front()/push_back()` | O(1)             | `dq.push_front(10);` (double-ended queue)  |
| **`queue`**         | `push()/pop()`          | O(1)                 | FIFO (BFS/level-order traversal)           |
| **`priority_queue`**| `push()/pop()`          | O(log n)             | `pq.push(3);` (max-heap by default)        |
|                     | `top()`                 | O(1)                 | `cout << pq.top();`                        |
| **`stack`**         | `push()/pop()`          | O(1)                 | LIFO (DFS/postfix evaluation)              |
| **`set/map`**       | `insert()/find()`       | O(log n)             | Ordered unique keys (BST-based)            |
| **`unordered_set/map`** | `insert()/find()` | O(1) avg, O(n) worst | Hash-table (fast lookups, no order)        |
| **`pair`**          | Access `first/second`   | O(1)                 | `pair<int, int> p = {1, 2};`               |
| **`bitset`**        | `set()/test()`          | O(1)                 | `bitset<8> bs(5);` (bitmasking)            |

---

### **2. Algorithms** (`#include <algorithm>`)
| Algorithm           | Function                | Time Complexity      | Example/Use Case                          |
|---------------------|-------------------------|----------------------|--------------------------------------------|
| **Sort**            | `sort(begin, end)`      | O(n log n)           | `sort(v.begin(), v.end());`                |
| **Reverse**         | `reverse(begin, end)`   | O(n)                 | `reverse(s.begin(), s.end());`             |
| **Binary Search**   | `binary_search()`       | O(log n)             | Check if `4` exists in a sorted vector.    |
| **Lower/Upper Bound**| `lower_bound()`        | O(log n)             | `auto lb = lower_bound(v.begin(), v.end(), 3);` |
| **Next Permutation**| `next_permutation()`    | O(n)                 | Generate permutations of `"123"`.          |
| **Max/Min**         | `max(a, b)`             | O(1)                 | `cout << max({1, 3, 2});` (C++11)          |

---

### **3. Utilities**
| Utility             | Operation               | Time Complexity      | Example/Use Case                          |
|---------------------|-------------------------|----------------------|--------------------------------------------|
| **`swap(a, b)`**    | Swap values             | O(1)                 | `swap(a, b);` (no copy for large objects)  |
| **`gcd(a, b)`**     | Compute GCD             | O(log min(a, b))     | `gcd(12, 18) = 6` (C++17)                 |
| **`lcm(a, b)`**     | Compute LCM             | O(log min(a, b))     | `lcm(4, 6) = 12` (C++17)                  |
| **`stringstream`**  | Parse strings           | O(n)                 | `stringstream ss("123"); int x; ss >> x;` |

---

### **4. Tips & Tricks**
1. **Fast I/O**:  
   ```cpp
   ios_base::sync_with_stdio(false); cin.tie(NULL);
   ```
2. **Lambda Comparator**:  
   ```cpp
   sort(v.begin(), v.end(), [](int a, int b) { return a > b; });
   ```
3. **Reserve Memory**:  
   ```cpp
   vector<int> v; v.reserve(1e6); // Faster for large data
   ```
4. **Erase-Remove Idiom**:  
   ```cpp
   v.erase(remove(v.begin(), v.end(), 0), v.end()); // Remove all zeros
   ```

---

### **Sample Problems**
#### **1. Find Duplicates in an Array**
- **Solution**: Use `unordered_map` (O(n) avg):  
  ```cpp
  unordered_map<int, int> freq;
  for (int num : nums) if (freq[num]++ >= 1) cout << num;
  ```

#### **2. Kth Largest Element**
- **Solution**: Use `priority_queue` (O(n log k)):  
  ```cpp
  priority_queue<int, vector<int>, greater<int>> pq; // Min-heap
  for (int num : nums) {
      pq.push(num);
      if (pq.size() > k) pq.pop();
  }
  cout << pq.top();
  ```
---

### **1. Containers**
#### **Vector** (`#include <vector>`)
- Dynamic array with O(1) random access.
- **Example**:
  ```cpp
  vector<int> v = {3, 1, 4};
  v.push_back(5);          // Add element
  sort(v.begin(), v.end()); // Sort: {1, 3, 4, 5}
  cout << v[2];            // 4 (access)
  ```

#### **Pair** (`#include <utility>`)
- Stores two elements.
- **Example**:
  ```cpp
  pair<int, string> p = {10, "apple"};
  cout << p.first << " " << p.second; // 10 apple
  ```

#### **Queue** (`#include <queue>`)
- FIFO structure (BFS).
- **Example**:
  ```cpp
  queue<int> q;
  q.push(10); q.push(20);
  cout << q.front(); // 10
  q.pop();           // Remove 10
  ```

#### **Priority Queue** (`#include <queue>`)
- Max-heap by default.
- **Example**:
  ```cpp
  priority_queue<int> pq;
  pq.push(3); pq.push(5); // Top: 5
  pq.pop();               // Remove 5
  ```

#### **Stack** (`#include <stack>`)
- LIFO structure.
- **Example**:
  ```cpp
  stack<int> s;
  s.push(10); s.push(20);
  cout << s.top(); // 20
  s.pop();         // Remove 20
  ```

#### **Set/Map** (`#include <set>`/`#include <map>`)
- Ordered unique keys (O(log n) operations).
- **Example**:
  ```cpp
  set<int> s = {3, 1, 4};
  s.insert(2);       // {1, 2, 3, 4}
  auto it = s.find(3); // Iterator to 3
  ```

#### **Unordered Set/Map** (`#include <unordered_set>`/`#include <unordered_map>`)
- Hash-based (average O(1) lookups).
- **Example**:
  ```cpp
  unordered_map<string, int> mp;
  mp["apple"] = 5; // Insert
  cout << mp["apple"]; // 5
  ```

#### **Deque** (`#include <deque>`)
- Double-ended queue.
- **Example**:
  ```cpp
  deque<int> dq;
  dq.push_front(10); dq.push_back(20); // [10, 20]
  ```

#### **Bitset** (`#include <bitset>`)
- Compact bit storage.
- **Example**:
  ```cpp
  bitset<8> bs(5); // 00000101
  bs.set(3);       // 00001101
  ```

---

### **2. Algorithms** (`#include <algorithm>`)
#### **Sort**
- **Example**:
  ```cpp
  vector<int> v = {3, 1, 4};
  sort(v.begin(), v.end()); // {1, 3, 4}
  ```

#### **Binary Search**
- **Example**:
  ```cpp
  if (binary_search(v.begin(), v.end(), 4)) {
      cout << "Found!";
  }
  ```

#### **Lower/Upper Bound**
- **Example**:
  ```cpp
  auto lb = lower_bound(v.begin(), v.end(), 3); // First element >= 3
  auto ub = upper_bound(v.begin(), v.end(), 3); // First element > 3
  ```

#### **Reverse**
- **Example**:
  ```cpp
  reverse(v.begin(), v.end()); // {4, 3, 1}
  ```

#### **Next Permutation**
- **Example**:
  ```cpp
  string s = "123";
  do {
      cout << s << endl; // 123, 132, 213, ...
  } while (next_permutation(s.begin(), s.end()));
  ```

#### **Max/Min**
- **Example**:
  ```cpp
  cout << max(3, 5); // 5
  cout << min({4, 2, 5}); // 2 (C++11)
  ```

---

### **3. Utility Functions**
#### **Swap** (`#include <utility>`)
- **Example**:
  ```cpp
  int a = 2, b = 3;
  swap(a, b); // a=3, b=2
  ```

#### **GCD/LCM** (`#include <numeric>` in C++17)
- **Example**:
  ```cpp
  cout << gcd(12, 18); // 6
  cout << lcm(4, 6);   // 12
  ```

#### **String Stream** (`#include <sstream>`)
- **Example**:
  ```cpp
  string s = "123 45";
  stringstream ss(s);
  int x; ss >> x; // x=123
  ```

---

### **4. Tips & Tricks**
1. **Use `auto` and Range-Based Loops**:
   ```cpp
   for (auto& num : v) num *= 2; // Double all elements
   ```
2. **Lambda Custom Comparators**:
   ```cpp
   sort(v.begin(), v.end(), [](int a, int b) { return a > b; }); // Descending
   ```
3. **Erase-Remove Idiom** (for vectors):
   ```cpp
   v.erase(remove(v.begin(), v.end(), 1), v.end()); // Remove all 1s
   ```
4. **Preallocate Memory** (for speed):
   ```cpp
   vector<int> v; v.reserve(1e6); // Avoid reallocations
   ```
5. **Fast I/O**:
   ```cpp
   ios_base::sync_with_stdio(false); cin.tie(NULL);
   ```

---

### **Sample Problem Solutions**
#### **Problem 1**: Find duplicates in an array.
- **Solution**:
  ```cpp
  unordered_map<int, int> freq;
  for (int num : nums) {
      if (freq[num] >= 1) cout << num;
      freq[num]++;
  }
  ```

#### **Problem 2**: Kth smallest element.
- **Solution**:
  ```cpp
  priority_queue<int> pq; // Max-heap
  for (int num : nums) {
      pq.push(num);
      if (pq.size() > k) pq.pop();
  }
  cout << pq.top();
  ```

---
