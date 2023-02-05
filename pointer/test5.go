package main

import "fmt"

func main() {
	f1()
}

func f1() {
	a := [...]int{1, 2, 3}
	var ap [3]*int
	for i := 0; i < len(a); i++ {
		ap[i] = &a[i]
	}

	fmt.Printf("ap: %T\n", ap)
	fmt.Printf("ap: %v\n", ap)

	for i := 0; i < len(ap); i++ {
		fmt.Printf("%v\n", *ap[i])
	}

}
