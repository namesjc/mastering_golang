package main

import (
	"fmt"
	"time"
)

func main() {
	GOTIME := "2006-01-02 15:04:05"
	ch := make(chan string, 5)
	go send(ch, GOTIME)
	for {
		select {
		case value, ok := <-ch:
			if ok {
				fmt.Println("value is", value, "ok is", ok)
			} else {
				fmt.Println("value is", value, "ok is", ok)
				return
			}
		case <-time.After(2 * time.Second):
			fmt.Println("Now is", time.Now().Format(GOTIME), "Time out")
			return

		}
	}
}

func send(ch chan string, GOTIME string) {

	defer func() {
		fmt.Println("closing channel")
		close(ch)
	}()

	for i := 0; i < 5; i++ {
		t := time.Now().Format(GOTIME)
		fmt.Println("i is", i, ", t is", t)
		ch <- t
		time.Sleep(3 * time.Second)
	}
}
