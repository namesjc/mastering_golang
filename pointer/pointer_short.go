package main

import "fmt"

func main() {
	str1 := "origin"
	pointer(&str1)
	fmt.Println(str1)

}

func pointer(p1 *string) {
	*p1 = "changed"
}
