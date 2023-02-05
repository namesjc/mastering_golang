package main

import "fmt"

func main() {
	str1 := "test memory address"

	p := &str1

	*p = "change value"

	fmt.Println(str1)
}
