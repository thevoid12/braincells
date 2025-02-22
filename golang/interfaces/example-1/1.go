// https://docs.google.com/document/d/1PLJdD4POmdgmCvbs0VMAQ5UiSg4tDfq_lJ_Isa6hr3A/edit?usp=sharing

package main

import (
	"fmt"
	"time"
)

// here we have a interface which defines a method signature Print()
// Any type that has a Print() method can satisfy this interface.
type PrintInterface interface {
	Print()
}

// we have 2 struct name and date which are completely different.
// one has a string field and another has a date field and each of its corresponding function implementation
// is also different (using name as input i want to print my name is : name, using date as input i want to
// print just the date).
// eventhough our end goal is different we are goint to expose same interface for both the goals
type Name struct {
	Name string
}

type Date struct {
	Date time.Time
}

// lets first see the not optimal approach
func (date Date) NonOptimalPrintDate() {
	fmt.Println(date.Date)
}

func (name Name) NonOptimalPrintName() {
	fmt.Println(name.Name)
}

// since the stuct name is a method for print() function which takes no argument and returns anything as same as our interface implementation,
// it satisfies our PrintInterface interface
func (name Name) Print() {
	fmt.Println("the name is:") // here we can do whatever we wanted to
	fmt.Println(name.Name)
}

// since the stuct date is a method for print() function which takes no argument and returns anything as same as our interface implementation,
// it satisfies our PrintInterface interface
func (date Date) Print() {
	fmt.Println(date.Date)
}

// we are exposing 1 interface to be called and this interface will implement different methods
// this interface will take types that satisfies the interface (here Date or Name struct as input which satisfies PrintInterface)
func InterfacePrintValue(value PrintInterface) {
	value.Print()
	// based on the value we give ie type name or type date it calls the appropriate print() method.
	// its now generic, we reduced name.NonOptimalPrintName() and name.NonOptimalPrintDate() into a generic interface
	// we never have to touch this function again and all we need to do is keep adding whatever methods that satisfies print interface
}

func main() {
	// non optimal approach demonstation
	// now from main we need to call 2 func to print name and date respectively
	name := Name{
		Name: "this is void",
	}
	date := Date{
		Date: time.Now(),
	}
	// call 2  functions to print name and date
	fmt.Println("********* non optimal apporach ****************")
	name.NonOptimalPrintName()
	date.NonOptimalPrintDate()
	fmt.Println("************************************************")
	// assume we have 100s of functions which does the same printing.
	// you can notice that printing is common among all the function just the input struct
	// type is different between the 2 non optimal print functions
	// so this is a classic case of using interface

	fmt.Println("********* interface approach ****************")
	// The name and date struct satisfies the Printer interface, you can call our interface imp with our defined types
	// to evoke the appropriate methods.
	//  since it has a Print() method that takes no arguments and returns no values.
	InterfacePrintValue(name)
	InterfacePrintValue(date)
	fmt.Println("************************************************")
}
