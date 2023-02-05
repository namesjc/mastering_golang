package main

import "fmt"

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	fmt.Println(len(a), cap(a))
	sa := a[2:5]
	fmt.Println(len(sa), cap(sa))
	fmt.Println(string(sa))

}
