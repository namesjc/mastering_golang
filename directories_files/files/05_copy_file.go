package main

import (
	"io"
	"log"
	"os"
)

func main() {
	originalFile, err := os.Open("testing2/test2.txt") //open original file
	check(err)
	defer originalFile.Close()

	newFile, err := os.Create("testing2/test3.txt") //create new file
	check(err)
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	check(err)
	log.Printf("Copied %d bytes.", bytesWritten) // 2023/02/10 22:55:12 Copied 19 bytes.

	err = newFile.Sync()
	check(err)

}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}
