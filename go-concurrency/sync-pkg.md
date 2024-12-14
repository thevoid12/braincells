# The sync Package
- The sync package contains the concurrency primitives that are most useful for low level memory access synchronization.
- If you’ve worked in languages that primarily handle concurrency through memory access synchronization, these types will likely
already be familiar to you.
- The difference between these languages in Go is that Go
has built a new set of concurrency primitives on top of the memory access synchronization primitives to provide you with an expanded set of things to work with.
## Wait Group:
- You can think of a WaitGroup like a concurrent-safe counter: calls to Add increment the counter by the integer passed in, and calls to Done decrement the counter by one. Calls to Wait block until the counter is zero.
- Notice that the calls to Add are done outside the goroutines they’re helping to track. else this will cause a race condition
 ```go
      var wg sync.WaitGroup
      wg.Add(1)
      go func() {
      defer wg.Done()
      fmt.Println("1st goroutine sleeping...")
      time.Sleep(1)
      }()
      wg.Add(1)
      go func() {
      defer wg.Done()
      fmt.Println("2nd goroutine sleeping...")
      time.Sleep(2)
      }()
      wg.Wait()
      fmt.Println("All goroutines complete.")
 ```
**output:** 
either 
1st goroutine sleeping...
2nd goroutine sleeping...
All goroutines complete 
**or**
2nd goroutine sleeping...
1st goroutine sleeping...
All goroutines complete 

## Mutex and RWMutex:
- Mutex stands for “mutual exclusion” and is a way to guard critical sections of your program.
- **Locking the same mutex twice without Unlocking will cause a deadlock!**
- > a critical section is an area of your program that requires exclusive access to a shared resource.
```go
type ConcurrentCounter struct {
	mu sync.Mutex
	value int
}

func (c *ConcurrentCounter) Increment() {
	c.mu.Lock() // restricts the memory pointers here. 
  // Guarantees only one write to value.
	c.value++
	c.mu.Unlock() // Being called with defer is also nice.
}
```

- The **sync.RWMutex** in Go is an enhanced version of a Mutex that provides greater flexibility by distinguishing between read and write locks. Here's how it works conceptually:
  - **Reader Lock (RLock):** Multiple readers can acquire a read lock concurrently, as long as there is no writer holding the write lock.
  - **Writer Lock (Lock):** Only one writer can acquire the write lock, and when it does, no readers are allowed access until the writer releases the lock.
   ``` go
          func main() {
          var rwMutex sync.RWMutex
          sharedData := 0

          // Writer goroutine
          go func() {
            for i := 1; i <= 5; i++ {
              rwMutex.Lock() // Acquire the write lock
              sharedData = i
              fmt.Printf("Writer updated sharedData to %d\n", sharedData)
              rwMutex.Unlock() // Release the write lock
              time.Sleep(500 * time.Millisecond) // Simulate infrequent updates
            }
          }()

          // Reader goroutines
          for j := 1; j <= 3; j++ {
            go func(readerID int) {
              for {
                rwMutex.RLock() // Acquire the read lock
                fmt.Printf("Reader %d read sharedData: %d\n", readerID, sharedData)
                rwMutex.RUnlock() // Release the read lock
                time.Sleep(100 * time.Millisecond) // Simulate frequent reads
              }
            }(j)
          }

          // Let the program run for a while to see the output
          time.Sleep(3 * time.Second)
        }

  ```
op:
Reader 1 read sharedData: 0
Reader 2 read sharedData: 0
Reader 3 read sharedData: 0
Writer updated sharedData to 1
Reader 1 read sharedData: 1
Reader 2 read sharedData: 1
Reader 3 read sharedData: 1
Writer updated sharedData to 2
Reader 1 read sharedData: 2
Reader 2 read sharedData: 2
Reader 3 read sharedData: 2
- > sync.rwmutex is very useful useful for scenarios where read operations dominate, but occasional write operations are needed.
