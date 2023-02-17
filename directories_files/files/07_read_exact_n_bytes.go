package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testing2/test3.txt")
	if err != nil {
		log.Println(err)
	}

	byteslice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteslice)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Number of bytes read: %d\n", numBytesRead) // Number of bytes read: 2
	fmt.Printf("Data read: %s\n", byteslice)               // Data read: AB
	fmt.Println(byteslice)
	fmt.Println(string(byteslice))

}
