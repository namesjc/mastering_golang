package main

import "fmt"

func main() {
	// s1 := make([]int, 3, 6)
	// fmt.Printf("%p\n", s1)
	// s1 = append(s1, 2, 3, 4)
	// fmt.Printf("%p %v\n", s1, s1)
	// s1 = append(s1, 2)
	// fmt.Printf("%p %v", s1, s1)

	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("address of a: %p\n", a)
	s3 := a[2:5]
	s4 := a[1:3]
	fmt.Println(s3, s4)
	fmt.Printf("address of s3, s4: %p %p\n", s3, s4)
	s3[0] = 9
	fmt.Println(s3, s4)
	fmt.Printf("address of s3, s4: %p %p\n", s3, s4)
}
