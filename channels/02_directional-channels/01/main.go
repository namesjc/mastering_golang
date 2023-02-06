package main

import "fmt"

func main() {
	c := make(chan int)    // General channel
	cr := make(<-chan int) // By directional: recieve only channel
	cs := make(chan<- int) // By directional: send only channel

	fmt.Println("---------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// general channel to specific channel converts
	fmt.Println("---------")
	fmt.Printf("cr\t%T\n", (<-chan int)(c))
	fmt.Printf("cs\t%T\n", (chan<- int)(c))

}
