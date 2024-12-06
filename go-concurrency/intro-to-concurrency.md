#  introduction to concurrency

## history
- For embarrassingly parallel problems, design your application to scale horizontally. This allows running multiple instances of your program on more CPUs or machines, reducing the overall runtime.
- Scaling horizontally became much easier in the early 2000s when a new paradigm
began to take hold: cloud computing. Instead of machines that you carefully
curated, installed software on, and maintained, cloud computing implied access to vast pools of resources that were provisioned into machines for workloads on-
demand.
- Cloud computing brought new challenges, including resource provisioning, communication between instances, and result aggregation. One major hurdle was modeling code for concurrency, especially when solutions ran across multiple machines. Addressing these challenges gave rise to "web scale" software, designed to handle massive workloads by scaling horizontally—adding more instances as needed. This approach enabled features like rolling updates, elastic scalability, and geographic distribution but also added complexity in understanding and ensuring fault tolerance.
- so by default every program has to have strong parallelism to achieve this.

## Why is Concurrency is hard?

Fortunately everyone runs into the same issues when working with concurrent code.
Because of this, computer scientists have been able to label the common issues, which
allows us to discuss how they arise, why, and how to solve them.
 - **Race Condition:**
    - A race condition occurs when the behavior of a program depends on the sequence or timing of uncontrollable events, such as the order in which threads execute. It typically arises in concurrent programming when multiple threads or processes access shared resources (e.g., variables or data structures) simultaneously, and at least one thread modifies the resource. If proper synchronization is not in place, this can lead to unpredictable or incorrect results.
    - Most of the time, data races are introduced because the developers are thinking about
    the problem sequentially. They assume that because a line of code (go routine) falls before another
    that it will run first.
    - When writing concurrent code, you have to meticulously iterate through the possible
scenarios.
    - Some developers tend to rely on adding sleeps throughout their code, as it appears to resolve their concurrency issues.
    but in reality they don't. they just decrease the probability of getting into a race condition. The takeaway here is that you should always target logical correctness. Introducing
    sleeps into your code can be a handy way to debug concurrent programs, but they are
    not a solution.
    - > always the output should be deterministic ie we always need to get the same result for the same input.
  - **Atomicity:**
    -  0 or 1. either complete the task fully or don't complete the task at all. no partially doing things.
    -  If a variable/resource is inside a goroutine (context) and its not shared with other routines, then it's atomic. Changes to that resource are uninterruptible and indivisible
  - **Memory Access Synchronization:**
    -  **critical section** is the name of the section of your program that needs exclusive access to a shared resource.
      ```go
            var data int
            go func() { data++}()  -----1
            if data == 0 { -----2
            fmt.Println("the value is 0.") -----3
            } else { -----2
            fmt.Printf("the value is %v.\n", data) ---3
            } 
      ```
    - here the shared resource is data of type int and 1,2,3 are the critical sections as it has access to the shared resources.
    - There are various ways to guard your program’s critical sections, and Go has some
better ideas on how to deal with this, but one way to solve this problem is to synchronize access to the memory between your critical sections(by using lock and mutexes)
    - > Anytime developers want to access the data variable’s memory, they must first call Lock, and when
they’re finished they must call Unlock. Code between those two statements can then
assume it has exclusive access to data; we have successfully synchronized access to the
memory.
    - but the calls to Lock you see can make our program
slow. Every time we perform one of these operations, our program pauses for a period
of time.
  - **Deadlock:**
    - A deadlocked program is one in which all concurrent processes are waiting on one
another. In this state, the program will never recover without outside intervention.
 - **LiveLock:**
    - Livelocks are programs that are actively performing concurrent operations, but these
operations do nothing to move the state of the program forward.
    - Livelock occurs when two or more processes continually repeat the same interaction in response to changes in the other processes without doing any useful work. These processes are not in the waiting state, and they are running concurrently. This is different from a deadlock because in a deadlock all processes are in the waiting state.
    - A livelock is similar to a deadlock, except that the states of the processes involved in the livelock constantly change with regard to one another, none progressing. Livelock is a special case of resource starvation; the general definition states that a specific process is not progressing.
  - **Starvation:**
    - Starvation is any situation where a concurrent process cannot get all the resources it
needs to perform work.
    - More broadly, starvation usually implies that there are one or more greedy concurrent process that are unfairly preventing one or more concurrent processes from
accomplishing work as efficiently as possible, or maybe at all.
    - Starvation is a problem that is closely related to both, Livelock and Deadlock. In a dynamic system, requests for resources keep on happening. Thereby, some policy is needed to make a decision about who gets the resource and when this process, being reasonable, may lead to some processes never getting serviced even though they are not deadlocked. 
    
