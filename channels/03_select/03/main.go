package main

import (
	"fmt"
	"time"
)

func main() {
	GOTIME := "2006-01-02 15:04:05"
	ch := make(chan string, 5)
	go send(ch, GOTIME)
	time.Sleep(10 * time.Second)
	fmt.Println("End at", time.Now().Format(GOTIME))

}

func send(ch chan string, GOTIME string) {

	for i := 0; i < 5; i++ {
		t := time.Now().Format(GOTIME)
		fmt.Println("i is", i, ", t is", t)
		ch <- t
		time.Sleep(1 * time.Second)
	}
}
