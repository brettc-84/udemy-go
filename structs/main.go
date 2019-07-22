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
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// passed by value
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNameReference(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	alex := person{"Alex", "Anderson", 32, contactInfo{"alex.a@gmail.com", 1234}}
	tony := person{
		firstName: "Tony",
		lastName:  "Sly",
		age:       41,
		contactInfo: contactInfo{
			email:   "tony@sly.com",
			zipcode: 91021,
		},
	}
	var dexter person

	dexter.firstName = "Dexter"
	dexter.lastName = "Morgan"

	fmt.Println(alex, tony, dexter)

	alex.updateName("Alexander")
	alex.print()

	alexPtr := &alex
	alexPtr.updateNameReference("Alexander")
	alex.print()

}
