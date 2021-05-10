package main

import (
	"fmt"

)

// global var
// var test ="blah"

func main() {
	// customers :=getData()
	// fmt.Println(customers)
	// fmt.Println(len(customers))

	customers := GetCustomers()
	for _, customer := range customers {
		fmt.Println(customer)
	}

}

// func getData()(customers []string) {
	// two ways to init the string
	// var firstName = "Marco"
	// lastName := "Doe"
	// create a empty string
	// var fullName string

	// customer := "Marco Doe"

	// customers[0] = customer
	// customers[1] = "Bob Smith"

	// customers = []string{"Marco Doe", "Bob Smith", "John Smith"}
	// customers = append(customers, "Adia Chan")

	// for x :=0; x < 4; x++ {
	// 	fmt.Println(customers[x])
	// }

	// for x , customer := range customers {

	// 	customer = customers[x]
	// 	fmt.Println(customer)
	// }

	// for _ , customer := range customers {
	// 	fmt.Println(customer)
	// }

	// for _,customer := range customers {
	// 	fmt.Println(customer)
	// }

	// return customers

// }


