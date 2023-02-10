package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "dir1"
	newPath := "dir4"

	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Println(err)
	}

	err = os.Remove(newPath)
	if err != nil {
		log.Println(err)
	}

	err = os.RemoveAll("dir2")
	if err != nil {
		log.Println(err)
	}
}
