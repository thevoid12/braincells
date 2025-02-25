## interfaces
- interfaces are very important implementations in go
- By generalizing, interfaces let us write functions that are more flexible and adaptable because they are not tied to the details of one paticular implementation.
- In Go, an interface is a type that defines a set of method signatures. Any type that implements all of the methods defined in the interface is said to satisfy the interface. This allows you to write code that is generic and can be used with any type that satisfies the interface.

```
package main

import (
	"fmt"
	"time"
)

// PrintInterface defines a method signature Print().
// Any type that implements the Print() method satisfies this interface.
type PrintInterface interface {
	Print()
}

// We have two different structs: Name and Date.
// - Name has a string field.
// - Date has a time.Time field.
//
// Each struct has its own Print() method with different logic:
// - Name's Print() method outputs: "The name is: <name>".
// - Date's Print() method outputs the date.
//
// Although their purposes differ, we will expose the same interface for both.
type Name struct {
	Name string
}

type Date struct {
	Date time.Time
}

// First, let's see a non-optimal approach where each struct has its own print function.
func (date Date) NonOptimalPrintDate() {
	fmt.Println(date.Date)
}

func (name Name) NonOptimalPrintName() {
	fmt.Println(name.Name)
}

// The struct Name implements the Print() method, which matches the PrintInterface signature.
// Since Print() takes no arguments and returns nothing, Name satisfies the PrintInterface.
func (name Name) Print() {
	fmt.Println("The name is:") // Custom logic for Name struct
	fmt.Println(name.Name)
}

// The struct Date also implements the Print() method in the same way,
// meaning it also satisfies the PrintInterface.
func (date Date) Print() {
	fmt.Println(date.Date)
}

// We define a generic function that accepts any type that satisfies PrintInterface.
// This function calls the Print() method on the given value, ensuring the appropriate method executes
// based on whether the value is a Name or Date struct.
func InterfacePrintValue(value PrintInterface) {
	value.Print()
	// This removes the need for separate functions like NonOptimalPrintName() and NonOptimalPrintDate().
	// We only need to add new types that implement Print(), without modifying this function.
}

func main() {
	// Demonstration of the non-optimal approach:
	// Here, we must explicitly call separate functions for each struct type.
	name := Name{
		Name: "this is void",
	}
	date := Date{
		Date: time.Now(),
	}

	fmt.Println("********* Non-optimal approach ****************")
	name.NonOptimalPrintName()
	date.NonOptimalPrintDate()
	fmt.Println("************************************************")

	// The non-optimal approach becomes inefficient when dealing with multiple similar types.
	// Since printing is a common behavior, we can generalize it using an interface.

	fmt.Println("********* Interface-based approach ****************")
	// The Name and Date structs both implement Print(), so we can use InterfacePrintValue() to call
	// their Print() methods generically.
	InterfacePrintValue(name)
	InterfacePrintValue(date)
	fmt.Println("************************************************")
}

```
- interface embedding: a interface inside a interface
suppose you have interface A, interface B and interface C which implements interface A and B then, we satisfy interface C if a type satisfies both interface A and B

```
// Interface Embedding: Go allows embedding interfaces within other interfaces.
package main

import "fmt"

// Crud is a composite interface that embeds four other interfaces.
// Any type that implements all four embedded interfaces will satisfy the Crud interface.
type Crud interface {
	Create
	Read
	Update
	Delete
}

// Four separate interfaces, each defining a single method.
type Create interface {
	PrintCreate()
}

type Read interface {
	PrintRead()
}

type Update interface {
	PrintUpdate()
}

type Delete interface {
	PrintDelete()
}

// Struct representing an application with a Name field.
type AppName struct {
	Name string
}

// The AppName struct implements all four interfaces by defining the required methods.
// Each method corresponds to its respective interface.

// AppName satisfies the Create interface because it implements PrintCreate().
func (an AppName) PrintCreate() {
	fmt.Println("Creating App Name:", an.Name)
}

// AppName satisfies the Read interface because it implements PrintRead().
func (an AppName) PrintRead() {
	fmt.Println("Reading App Name:", an.Name)
}

// AppName satisfies the Update interface because it implements PrintUpdate().
func (an AppName) PrintUpdate() {
	fmt.Println("Updating App Name:", an.Name)
}

// AppName satisfies the Delete interface because it implements PrintDelete().
func (an AppName) PrintDelete() {
	fmt.Println("Deleting App Name:", an.Name)
}

// ExposeInterface accepts any type that satisfies the Crud interface.
// Since Crud embeds Create, Read, Update, and Delete, the provided type must implement all four methods.
func ExposeInterface(an Crud) {
	an.PrintCreate()
	// Since AppName implements all four interfaces, it automatically satisfies the Crud interface.
	// This means an instance of AppName can be passed to ExposeInterface.

	// Uncomment the following lines to perform all CRUD operations.
	// an.PrintRead()
	// an.PrintUpdate()
	// an.PrintDelete()
}

func main() {
	an := AppName{
		Name: "void",
	}
	ExposeInterface(an)
}

```
- To satisfy an interface in Go, a type must implement all of the methods defined in the interface. 

```
// To satisfy an interface in Go, a type must implement all of the methods defined in the interface.
package main

import "fmt"

// Calculator interface defines five methods that a type must implement to satisfy it.
type Calculator interface {
	Add(a, b int) (int, error)
	Subtract(a, b int) (int, error)
	Multiply(a, b int) (int, error)
	Divide(a, b int) (int, error)
	Negate(a int) (int, error)
}

// If a struct needs to satisfy the Calculator interface, it must implement all the methods defined in the interface.

type Input struct {
	CalcName string
}

// The Input struct implements all methods of the Calculator interface.

func (ip *Input) Add(a, b int) (int, error) {
	fmt.Printf("Adding for calc %s:\n", ip.CalcName)
	return a + b, nil
}

func (ip *Input) Subtract(a, b int) (int, error) {
	fmt.Printf("Subtracting for calc %s:\n", ip.CalcName)
	return a - b, nil
}

func (ip *Input) Multiply(a, b int) (int, error) {
	fmt.Printf("Multiplying for calc %s:\n", ip.CalcName)
	return a * b, nil
}

func (ip *Input) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	fmt.Printf("Dividing for calc %s:\n", ip.CalcName)
	return a / b, nil
}

func (ip *Input) Negate(a int) (int, error) {
	fmt.Printf("Negating for calc %s:\n", ip.CalcName)
	return -a, nil
}

// ImplementInterface takes a Calculator interface and performs an addition operation.
func ImplementInterface(calc Calculator, a, b int) (int, error) {
	return calc.Add(a, b)
}

func main() {
	ip := &Input{CalcName: "BasicCalculator"}
	ans, _ := ImplementInterface(ip, 5, 5)
	fmt.Println("Result:", ans)
}

```

- final example:

```
package main

import "fmt"

// Define an interface that can be implemented by multiple types.
type test interface {
	Print()
}

// Concrete implementation: Country.
type Country struct {
	Name      string
	Code      int
	Rank      int
	Continent string
}

func (c *Country) Print() {
	fmt.Println("Country Details:")
	fmt.Println("Name:", c.Name)
	fmt.Println("Code:", c.Code)
	fmt.Println("Rank:", c.Rank)
	fmt.Println("Continent:", c.Continent)
}

// Another concrete implementation: State.
type State struct {
	Name       string
	Country    string
	Rank       int
	Population int
}

func (s *State) Print() {
	fmt.Println("State Details:")
	fmt.Println("Name:", s.Name)
	fmt.Println("Country:", s.Country)
	fmt.Println("Rank:", s.Rank)
	fmt.Println("Population:", s.Population)
}

// Function returning the interface.
func GetInterface(isCountry bool, name, location string, rank, extra int) test {
	if isCountry {
		return &Country{Name: name, Code: extra, Rank: rank, Continent: location}
	}
	return &State{Name: name, Country: location, Rank: rank, Population: extra}
}

func main() {
	// Using an interface to store different types.
	var entity test

	// Assigning a Country.
	entity = GetInterface(true, "India", "Asia", 1, 91)
	entity.Print()

	// Assigning a State.
	entity = GetInterface(false, "Karnataka", "India", 2, 67000000)
	entity.Print()
}

// How This Demonstrates Key OOP Principles in Go:

// 1. Abstraction & Decoupling:
//    - main() does not need to know whether it is working with a Country or State.
//    - It just calls Print() on the test interface, hiding implementation details.

// 2. Polymorphism:
//    - GetInterface() returns either a Country or a State, both implementing the test interface.
//    - Print() behaves differently depending on the stored type.

// 3. Encapsulation:
//    - The caller (main) depends only on the test interface, not on Country or State directly.

// 4. Future Extendability:
//    - You can add more types (e.g., City) that implement test without modifying GetInterface() or main().

// 5. Mocking for Testing:
//    - You can create a mock struct that implements test, enabling unit tests without real data.


```
