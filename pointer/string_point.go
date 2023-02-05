package main

import "fmt"

func main() {
	var arr [5]string
	arr = [5]string{"1", "2", "3", "4", "5"}

	var arrP *[5]string
	arrP = &arr
	fmt.Println(arrP)

	var arrpp [5]*string

	str1 := "str1"
	str2 := "str2"
	str3 := "str3"
	str4 := "str4"
	str5 := "str5"

	arrpp = [5]*string{&str1, &str2, &str3, &str4, &str5}

	*arrpp[2] = "555"

	fmt.Println(str3)

}
