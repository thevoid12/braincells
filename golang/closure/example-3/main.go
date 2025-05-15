package main

import (
	"fmt"
	"runtime"
)

func test2(v int) func() int {
	i := v
	return func() int { // this is a closure
		return i + 2
	}
}

func test3(v int) func() int {
	i := v
	return func() int { // this is a stateful closure
		i = i + 2
		return i
	}
}

func main() {
	// 	fmt.Println(test2(1)):
	// Calls test2(1)
	// Sets i = 1
	// Returns a closure: func() int { return i + 2 }
	// fmt.Println() prints the function itself, which results in something like a functional pointer address
	// will print the address of the function
	fmt.Println(test2(1))

	//Calls test2(1) again (separately)
	// Sets i = 1
	// Returns a new closure capturing a new i
	// Stores the returned function in variable a — a is of type func() int
	// At this point, memory looks like this:
	// a --> func() int {
	//          return i + 2  // i is captured and equals 1
	//      }

	a := test2(1)

	// ans := a()
	// Calls the closure stored in a

	// i = 1, so i + 2 = 3
	// ans = 3
	ans := a()
	fmt.Println(ans)

	// 	Calls the same closure again
	// Still returns 3 because i hasn’t changed
	ans = a()

	fmt.Println(ans)

	// 	Calls the same closure again
	// Still returns 3 because i hasn’t changed
	fmt.Println(a())

	fmt.Println("/********************************************************/")
	// 	returning a closure from test3(v) that modifies and remembers i.
	// Every time the returned function is called, it increments i by 2 and returns the new value.
	// This is a classic example of a stateful closure.

	// prints the address of the closure function
	fmt.Println(test3(1))

	// c := test3(1)
	// Calls test3(1), sets i = 1
	// Stores the returned closure in variable c
	// prints the address of the closure function as it is stored in c
	c := test3(1)
	fmt.Println(c)

	ans1 := c()
	fmt.Println(ans1) // 3
	fmt.Println(c())  // 5
	fmt.Println(c())  // 7
	// this happens because of closure lifetime. as long as the instance of the function
	// exists, The scope of i is still inside test3() but the lifetime
	// of i is extended because the closure captures it as it is stored in a heap somewhere else
	// Variables don’t "move" scope, but their storage location (heap vs. stack) may change to support longer lifetimes.

	d := test3(2)
	fmt.Println(d()) // 4
	fmt.Println(d()) // 6
	fmt.Println(d()) // 8
	// eventhough i's scope is outide the func scope of the i still lies within the overall closure
	// technically there is 2 i values has been stored in the heap
	// 	Each call to test3(v) creates a new closure with its own i
	// That i lives on the heap, because it's captured by the closure
	// The closure "remembers" its i across calls, even though test3 has returned

	fmt.Println("/************************garbage collection********************************/")
	// When nothing references the closure anymore, and it becomes unreachable, the Go garbage collector (GC) will eventually clean the heap memory.
	// 	Go automatically manages memory — you don’t need to manually free heap objects.
	// But you can leak memory if:
	// You keep closures around in global maps or slices unintentionally.
	// You use closures in long-lived goroutines that never exit.
	c = test3(1)
	fmt.Println(c()) // 3
	fmt.Println(c()) // 5

	c = nil      // remove reference
	runtime.GC() // hint GC to run
	fmt.Println("Done")

}

//even though i is local to test2, the returned function keeps it alive. (lifetime)
// This behavior occurs because the closure captures the variable i, extending its lifetime beyond the execution of test3().
// In Go, when a closure captures a local variable, that variable is moved from the stack to the heap so it can persist.
// Each time test3 is called, a new closure is created with its own independent copy of i stored on the heap.
//  Even though i is declared inside test3, it's kept alive and accessible because the closure retains a reference to it.

//Analogy: “Boxed Variables”
// Think of i as a box.

// When you return a closure, Go boxes i and keeps it on the heap.

// Every call to test3(v) makes a new box with its own i

// The closure carries a pointer to its box
