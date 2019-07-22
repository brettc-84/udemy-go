package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	alex := person{"Alex", "Anderson", 32}
	tony := person{firstName: "Tony", lastName: "Sly", age: 41}
	var dexter person

	dexter.firstName = "Dexter"
	dexter.lastName = "Morgan"

	fmt.Println(alex, tony, dexter)

	fmt.Printf("%+v\n", alex)
}
