// interface embedding: Go also allows you to embed interfaces within other interfaces
package main

import "fmt"

// interface embedding
type Crud interface {
	Create
	Read
	Update
	Delete
}

// 4 interfaces which has its own method as signature
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

// our type
type AppName struct {
	Name string
}

// here appname type method implements 4 different interfaces

// appname struct implements create interface as it satisfies the create interface condition of PrintCreate()
// similaly for read update delete
func (an AppName) PrintCreate() {
	fmt.Println("Creating App Name:", an.Name)
}

func (an AppName) PrintUpdate() {
	fmt.Println("Updating App Name:", an.Name)
}
func (an AppName) PrintDelete() {
	fmt.Println("Deleting App Name:", an.Name)
}
func (an AppName) PrintRead() {
	fmt.Println("Reading App Name:", an.Name)
}

// if the type wants to satisfy crud interface,
//
//	then the type must implement all 4 interfaces in the crud interface(create read update delete)
func ExposeInterface(an Crud) {
	an.PrintCreate()
	// the moment appname struct implemented the 4 interfaces it automatically can implement crud interface as well
	// using the input type to perform action is our personal preference and depends on our use case.

	// an.PrintUpdate()
	// an.PrintUpdate()
	// an.PrintDelete()
}

func main() {
	an := AppName{
		Name: "void",
	}
	ExposeInterface(an)
}
