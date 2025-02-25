package main

import "fmt"

// Define an interface that can be implemented by multiple types
type test interface {
	Print()
}

// Concrete implementation: Country
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

// Another concrete implementation: State
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

// Function returning the interface
func GetInterface(isCountry bool, name, location string, rank, extra int) test {
	if isCountry {
		return &Country{Name: name, Code: extra, Rank: rank, Continent: location}
	}
	return &State{Name: name, Country: location, Rank: rank, Population: extra}
}

func main() {
	// Using interface to store different types
	var entity test

	// Assigning a Country
	entity = GetInterface(true, "India", "Asia", 1, 91)
	entity.Print()

	// Assigning a State
	entity = GetInterface(false, "Karnataka", "India", 2, 67000000)
	entity.Print()
}

// How This Demonstrates Each Advantage?
// 1. Abstraction & Decoupling
// main() does not need to know whether it is working with Country or State.
// It just calls Print() on the test interface.
// 2. Polymorphism
// GetInterface returns either a Country or a State, both implementing test.
// Print() behaves differently depending on which type is stored.
// 3. Encapsulation
// The caller (main) does not directly depend on Country or State, only on test.
// 4. Future Extendability
// You can add more types (e.g., City) without changing the function signatures.
// 5. Mocking for Testing
// In a unit test, you can create a mockTest struct that implements test,
// making it easy to test behavior without depending on real data.
