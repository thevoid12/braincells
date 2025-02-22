## interfaces
- In Go, an interface is a type that defines a set of method signatures. Any type that implements all of the methods defined in the interface is said to satisfy the interface. This allows you to write code that is generic and can be used with any type that satisfies the interface.
- code
- interface embedding: a interface inside a interface
suppose you have interface A, interface B and interface C which implements interface A and B then, we satisfy interface C if a type satisfies both interface A and B
- code
- To satisfy an interface in Go, a type must implement all of the methods defined in the interface. 
- code
