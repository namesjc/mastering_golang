package main

import "fmt"

// golang value type: int, float, string, bool, structs

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip) //默认不赋值 nil

	var i int = 100
	ip = &i
	fmt.Printf("ip: %T\n", ip)
	fmt.Printf("ip: %v\n", *ip)

	var strp *string

	var str string = "test"
	strp = &str
	fmt.Printf("strp: %T\n", strp)
	fmt.Printf("strp: %v\n", *strp)

	var bp *bool

	var b bool = true
	bp = &b
	fmt.Printf("bp: %T\n", bp)
	fmt.Printf("bp: %v\n", *bp)

}
