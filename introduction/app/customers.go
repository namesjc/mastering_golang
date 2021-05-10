package main

type (
	Customer struct {
		FirstName string
		LastName string
		FullName string
	}
)


func GetCustomers()(customers []Customer) {

	marcel := Customer{ FirstName: "Marcel", LastName: "Deo"}

	customers = append(customers, marcel,
		Customer{ FirstName: "Marco", LastName: "Smith"},
		Customer{ FirstName: "Adia", LastName: "Chan"},
		Customer{ FirstName: "Victor", LastName: "Yu"},
	)

	return customers
}