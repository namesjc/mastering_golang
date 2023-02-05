package main

import "fmt"

func main() {
	x := bar()
	fmt.Println(&x)
	fmt.Printf("%T", x)
}

func bar() func() int {
	return func() int {
		return 451
	}
}
