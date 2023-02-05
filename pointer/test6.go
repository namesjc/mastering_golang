package main

import "fmt"

type Integer int

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	var b LessAdder = &a
	fmt.Printf("b: %v\n", b.Less(2))
	fmt.Printf("b: %v\n", b.Add(2))

}

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}
