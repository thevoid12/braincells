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

--- 
## communication and synchronization between goroutines:
  - In Go, broadcast and signal concepts are used for communication and synchronization between goroutines, typically implemented using channels or sync.Cond. 
  ### Broadcast:
  - Broadcast is used to notify multiple goroutines that a specific condition has been met.
  - When one goroutine signals a broadcast, all waiting goroutines are notified and can proceed.
  -  The **sync.Cond** type provides a mechanism for broadcasting to multiple goroutines.

      ```go
      package main

      import (
        "fmt"
        "sync"
        "time"
      )

      func main() {
        var mu sync.Mutex
        cond := sync.NewCond(&mu)

        done := make(chan struct{})
        var wg sync.WaitGroup

        // Goroutines waiting for the signal
        for i := 1; i <= 3; i++ {
          wg.Add(1)
          go func(id int) {
            defer wg.Done()

            cond.L.Lock()
            cond.Wait() // Wait for broadcast signal
            cond.L.Unlock()

            fmt.Printf("Goroutine %d received broadcast\n", id)
          }(i)
        }

        time.Sleep(1 * time.Second)

        // Broadcasting to all waiting goroutines
        go func() {
          fmt.Println("Broadcasting to all goroutines")
          cond.Broadcast()
          close(done)
        }()

        wg.Wait()
        <-done
        fmt.Println("All goroutines finished")
      }
      ```
      **output:**
      - The Broadcast method notifies all goroutines waiting on the condition.
      - Each goroutine unblocks and proceeds after receiving the broadcast.
      Broadcasting to all goroutines
      Goroutine 1 received broadcast
      Goroutine 2 received broadcast
      Goroutine 3 received broadcast
      All goroutines finished
  ### Signal:
  - unlike broadcast it wakes up only one goroutine. The Signal method notifies one waiting goroutine.
  ```go
      package main
      import (
        "fmt"
        "sync"
        "time"
      )

      func main() {
        var mu sync.Mutex
        cond := sync.NewCond(&mu)

        var wg sync.WaitGroup

        // Goroutines waiting for the signal
        for i := 1; i <= 3; i++ {
          wg.Add(1)
          go func(id int) {
            defer wg.Done()

            cond.L.Lock()
            cond.Wait() // Wait for signal
            cond.L.Unlock()

            fmt.Printf("Goroutine %d received signal\n", id)
          }(i)
        }

        time.Sleep(1 * time.Second)

        // Signaling individual goroutines one at a time
        go func() {
          for i := 1; i <= 3; i++ {
            time.Sleep(500 * time.Millisecond)
            fmt.Println("Signaling one goroutine")
            cond.Signal()
          }
        }()

        wg.Wait()
        fmt.Println("All goroutines finished")
      }
  ```
**output:**
Signaling one goroutine
Goroutine 1 received signal
Signaling one goroutine
Goroutine 2 received signal
Signaling one goroutine
Goroutine 3 received signal
All goroutines finished
note that: The order in which goroutines are woken up depends on the runtime scheduler.

- while sync.cond gives a method do so signalling and broadcasting it is a primitive way (as of everything present in sync package) and a better way to do this is using channels 
  - ##### broadcast with channel:
    - ```go
      package main
      import (
        "fmt"
        "time"
      )

      func main() {
        done := make(chan struct{})

        // Goroutines waiting for broadcast
        for i := 1; i <= 3; i++ {
          go func(id int) {
            <-done // Wait for broadcast
            fmt.Printf("Goroutine %d received broadcast\n", id)
          }(i)
        }

        time.Sleep(1 * time.Second)

        // Broadcasting to all goroutines
        close(done) // Closing the channel broadcasts to all
        fmt.Println("Broadcast sent to all goroutines")
        time.Sleep(500 * time.Millisecond)
      }
      ```
      
   **output:**
Broadcast sent to all goroutines
Goroutine 1 received broadcast
Goroutine 2 received broadcast
Goroutine 3 received broadcast
  - ##### signal with channel:
  ```go
  package main
  import (
    "fmt"
    "time"
  )

  func main() {
    signal := make(chan struct{})

    // Goroutine waiting for signal
    go func() {
      <-signal // Wait for signal
      fmt.Println("Goroutine received signal")
    }()

    time.Sleep(1 * time.Second)

    // Signaling the goroutine
    signal <- struct{}{}
    fmt.Println("Signal sent")
    time.Sleep(500 * time.Millisecond)
  }
  ```
 **output:**
Signal sent
Goroutine received signal

## sync.Once:
 - sync.Once is a type that utilizes some sync primitives internally
to ensure that only **one** call to **Do** ever calls the function passed in—even on different
goroutines.

    ```go
          var count int
          increment := func() { count++ }
          decrement := func() { count-- }
          var once sync.Once
          once.Do(increment)
          once.Do(decrement)
          fmt.Printf("Count: %d\n", count)
          This produces:
          Count: 1
         
         
    ```
  - This is because sync.Once only
counts the number of times Do is called, not how many times unique functions passed
into Do are called. In this way, copies of sync.Once are tightly coupled to the functions they are intended to be called with
