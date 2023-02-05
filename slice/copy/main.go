package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s1)
	s2 := []int{7, 8, 9}
	copy(s1, s2)
	fmt.Println(s1)
}
