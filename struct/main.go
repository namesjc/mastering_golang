package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	// marco := person{firstName: "Marco", lastName: "Chan"}
	// var marco person
	// marco.firstName = "Marco"
	// marco.lastName = "Chan"
	// fmt.Println(marco)
	// fmt.Printf("%+v", marco)

	jim := person{
		firstName: "Jim",
		lastName: "Deo",
		contact: contactInfo{
			email: "marco@demo.com",
			zipCode: 215000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPersion *person) updateName(newFirstName string) {
	(*pointerToPersion).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}