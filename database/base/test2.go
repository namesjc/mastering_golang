package main

import "fmt"

func main() {
	fmt.Println("return:", test()) // defer 和 return之间的顺序是先返回值, i=0，后defer
}

func test() int { //这里返回值没有命名
	var i int
	defer func() {
		i++
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2", i)
	}()
	return i
}
