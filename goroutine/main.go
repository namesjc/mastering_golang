// package main

// import (
// 	"fmt"
// 	"time"
// )

// func DelayPrint() {
// 	for i := 1; i <= 4; i ++ {
// 		time.Sleep(time.Microsecond * 250)
// 		fmt.Println(i)
// 	}
// }

// func HelloWorld() {
// 	fmt.Println("Hello goroutin!")
// }

// func main() {
// 	go DelayPrint()
// 	go HelloWorld()
// 	time.Sleep(time.Second *2)
// 	fmt.Println("main function")
// }

package main

import "fmt"

func main() {
   
   messages := make(chan string)

   go func() { messages <- "ping" }()

   msg := <-messages
   fmt.Println(msg)
}