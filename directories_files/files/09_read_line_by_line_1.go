package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testing2/test1.txt")
	if err != nil {
		log.Println(err)
	}

	//start reading from the file with a reader
	reader := bufio.NewReader(file)
	line, counter := "", 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		counter++
		fmt.Printf("line %v > Read %d characters > %v", counter, len(line), line)
	}

	if err != io.EOF { // EOF = END OF FILE
		fmt.Printf(" > Failed!: %v\n", err)
	}
	// line 1 > Read 6 characters > Hello
	// line 2 > Read 7 characters > Golang

}
