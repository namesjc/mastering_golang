package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testing2/test1.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	counter := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() { //Default delimiter = '\n'
		counter++
		err := scanner.Err()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("line %v > Read %d characters > %v\n", counter, len(scanner.Text()), scanner.Text())
		// line 1 > Read 5 characters > Hello
		// line 2 > Read 6 characters > Golang
	}
}
