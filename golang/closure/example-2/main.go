package main

import "fmt"

func test() {
	fmt.Println("test")
}
func test2(v string) func(l string) {
	fmt.Println(v + ":top")
	return func(k string) {
		fmt.Println(k + ":bottom")
	}
}

func main() {
	// test():
	// A simple function.
	// When called, it prints "test" to stdout.
	// It takes no arguments, returns nothing.
	// Stored in memory as a regular function pointer.
	test()

	// test2(outer):
	// A function that returns another function.
	// When called, it prints outer:top to stdout.
	// It returns the inner function func(k string) {fmt.Println(k + ":bottom")} as a value.
	// It returns a function that, when called, prints " :bottom" to stdout.
	// The returned function is stored in memory as a closure.
	// The returned function has access to the variables in the scope of test2().
	// This is a closure because it captures the environment in which it was created.
	// in this case The returned anonymous function is discarded (not stored or called).
	// It's like calling and throwing away _ = test2()
	test2("outer")

	// I am not calling test2 here.
	// I am just assigning the function test2 itself to a variable a.
	// So now a is a function with the same type: func() func().
	// think of it as a function pointer. var a func() func() = test2
	// When I call a(), it will execute the test2 function.
	// The test2 function will print "test2- top" to stdout.
	// It will return the inner function func() { fmt.Println("test2- bottom") } as a value.
	a := test2
	a("outer")

	//  Calls test2()
	// Prints "test2- top".
	// Returns the **anonymous function** `func() { fmt.Println("test2- bottom") }`.
	//  That inner function is stored in `b`.
	b := test2("outer")
	// we are executing the inner function so it will print test2-bottom
	b("closureeee")
	fmt.Println("hiiii")
}

// OP:
// test
// outer:top
// outer:top
// outer:top
// closureeee:bottom
// hiiii

// memory:
//Go stores that anonymous function (test2-bottom part)on the heap (if it captures variables) or on the stack (if it doesn’t), depending on whether it escapes the local scope.
// In my code, the inner function:
// Doesn’t capture any variable, so it's just a regular function value.
