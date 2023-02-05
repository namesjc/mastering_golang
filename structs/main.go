package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@demo.com",
			zipCode: 215000,
		},
	}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "alex@demo.com"
	// alex.contact.zipCode = 215000

	// & is memory addrss
	// alexPointer := &alex
	// alexPointer.updateName("Alexandar")
	// or so called method set
	(&alex).updateName("Echo")
	alex.print()
	alex.updateName("Alexandar")
	alex.print()
}

// * means this is a type description - it means we're working with a pointer to a person
func (pointorToPerson *person) updateName(newFirstName string) {
	// this is an operator - it means we want to manipulate the value the pointer is referencing
	(*pointorToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
