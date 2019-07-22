package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	alex := person{"Alex", "Anderson", 32, contactInfo{"alex.a@gmail.com", 1234}}
	tony := person{firstName: "Tony", lastName: "Sly", age: 41, contact: contactInfo{email: "tony@sly.com", zipcode: 91021}}
	var dexter person

	dexter.firstName = "Dexter"
	dexter.lastName = "Morgan"

	fmt.Println(alex, tony, dexter)

	fmt.Printf("%+v\n", alex)
}
