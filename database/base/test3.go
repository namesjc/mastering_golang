package main

import "fmt"

func main() {
	fmt.Println("return:", test())
}

func test() (i int) { //返回值命名i
	defer func() {
		i++
		fmt.Println("defer1", i)
	}()
	defer func() {
		i++
		fmt.Println("defer2", i)
	}()
	return i // 也可直接写成 return
}
