# pointers in go
& is Used on the right side of the operand to return the address (actual) value.
* is Used on the left side of the operand to resolve the address value.
& makes a pointer from a variable.
* "fetches" the value stored where a pointer points to.

```go


package main

import (
    "fmt"
)

func main() {

    var a = 5
    var p = &a // copy by reference
    var x = a  // copy by value

    fmt.Println("a = ", a)   // a =  5
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  5
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  5

    fmt.Println("\n Change *p = 3")
    *p = 3
    fmt.Println("a = ", a)   // a =  3
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  3
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  5

    fmt.Println("\n Change a = 888")
    a = 888
    fmt.Println("a = ", a)   // a =  888
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  888
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  5

    fmt.Println("\n Change x = 1")
    x = 1
    fmt.Println("a = ", a)   // a =  888
    fmt.Println("p = ", p)   // p =  0x10414020
    fmt.Println("*p = ", *p) // *p =  888
    fmt.Println("&p = ", &p) // &p =  0x1040c128
    fmt.Println("x = ", x)   // x =  1
    
    &p = 3 // error: Cannot assign to &p because this is the address of variable a
}
```
But in Go:

&T is the operator that produces a pointer.

*T is the type of a pointer.
```go
Example:

v := Vertex{3, 4}
p := &v           // p has type *Vertex

1``
&v → gives you a *Vertex value (pointer to v).

*Vertex → is the type of that pointer.