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

	byteslice, minBytes := make([]byte, 32), 8
	numBytesRead, err := io.ReadAtLeast(file, byteslice, minBytes)
	// will return an error if it cannot find at least minBytes to read.
	// it will read as many bytes as byteSlice can hold
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Number of bytes read: %d\n", numBytesRead) // Number of bytes read: 19
	fmt.Printf("Data read: %s\n", byteslice)               // Data read: ABCBuffered string

}
