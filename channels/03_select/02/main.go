package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go send(ch)

	for {
		value, ok := <-ch
		if ok {
			fmt.Println("value is", value, "ok is ", ok)
		} else {
			fmt.Println("value is", value, "ok is", ok)
			break
		}

	}
}

func send(ch chan string) {
	GOTIME := "2006-01-02 15:04:05"
	defer func() {
		fmt.Println("closing channel")
		close(ch)
	}()
	for i := 0; i < 5; i++ {
		ch <- time.Now().Format(GOTIME)
		time.Sleep(1 * time.Second)
	}
}
