package main

import "fmt"

type Car interface {
	Type() string
	Name() string
}

type Sedan struct {
	name string
}

func (n *Sedan) Name() string {
	return "the name of sedan is:" + n.name
}

func (n *Sedan) Type() string {
	return "the type of the car is sedan"
}

type SUV struct {
	name string
}

func (n *SUV) Name() string {
	return "the name of suv is:" + n.name
}

func (n *SUV) Type() string {
	return "the type of the car is suv"
}

// Example of factory pattern
func factoryExample() {
	t, name := factory(1, "void") // client
	fmt.Println(t, name)
}

func factory(i int, s string) (string, string) { // factory
	var c Car
	if i == 1 {
		c = &Sedan{name: s}
	} else if i == 2 {
		c = &SUV{}
	}
	return c.Name(), c.Type()
}
