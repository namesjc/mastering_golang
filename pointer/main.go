package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the adress

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // dereferencing, * gives you the value stored at an adress when you have the adress
	fmt.Println(*&a)

	*b = 43
	fmt.Println(a)

}
