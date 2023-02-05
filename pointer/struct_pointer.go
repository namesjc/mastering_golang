package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	alex := person{
		first: "Echo",
		last:  "Chan",
	}

	(&alex).updateLast("Chen")
	print(alex)
}

func (p *person) updateLast(lastName string) {
	(*p).last = lastName
}

func print(p person) {
	fmt.Printf("p: %v\n", p)
}
