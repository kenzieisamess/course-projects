package main

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Go takes parameters as order of fields
	// alex := person {"Alex", "Anderson"}

	alex := person{firstName: "Alex", lastName: "Anderson"}
}
