// To satisfy an interface in Go, a type must implement all of the methods defined in the interface.
package main

import "fmt"

type Calculator interface {
	Add(a, b int) (int, error)
	Subract(a, b int) (int, error)
	Multiply(a, b int) (int, error)
	Divide(a, b int) (int, error)
	Negate(a int) (int, error)
}

// if a struct need to satisfy the calculator interface, it needs to implement all the methods in the interface

type Input struct {
	CalcName string
}

func (ip *Input) Add(a, b int) (int, error) {
	fmt.Println("Adding for calc %s:", ip.CalcName)
	return a + b, nil
}

func (ip *Input) Subract(a, b int) (int, error) {
	fmt.Println("Subract for calc %s:", ip.CalcName)
	return a - b, nil
}
func (ip *Input) Multiply(a, b int) (int, error) {
	fmt.Println("Multiply for calc %s:", ip.CalcName)
	return a * b, nil
}
func (ip *Input) Divide(a, b int) (int, error) {
	fmt.Println("Divide for calc %s:", ip.CalcName)
	return a / b, nil
}
func (ip *Input) Negate(a int) (int, error) {
	fmt.Println("Negate for calc %s:", ip.CalcName)
	return -a, nil
}

func ImplementInterface(calc Calculator, a, b int) (int, error) {

	return calc.Add(a, b)
}

func main() {
	ip := &Input{}
	ans, _ := ImplementInterface(ip, 5, 5)
	fmt.Println(ans)

}
