package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

// *type type description as a pointer to that type
func (p *person) updateName(newFirstName string) {
	// The value of the memory address with * operator
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	// Go takes parameters as order of fields
	// alex := person {"Alex", "Anderson"}

	alex := person{firstName: "Alex", lastName: "Anderson"}

	// Go is a pass by value language
	// Access the memory address of alex
	alexPointer := &alex
	alexPointer.updateName("Alexandra")

	// Shortcut for Pointers
	alex.updateName("Alexis")
	alex.print()
}
