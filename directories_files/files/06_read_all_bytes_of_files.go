package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("testing2/test3.txt")
	if err != nil {
		log.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)                               // [65 66 67 66 117 102 102 101 114 101 100 32 115 116 114 105 110 103 10]
	fmt.Printf("Data as string: %s\n", data)        // Data as string: ABCBuffered string
	fmt.Println("Number of bytes read:", len(data)) // Number of bytes read: 19

}
