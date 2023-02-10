package main

import (
	"fmt"
	"os"
)

func main() {
	originalPath := "testing/test1.txt"
	newPath := "testing/test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Rename(newPath, "testing2/test2.txt")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove("testing2/test2.txt")
	if err != nil {
		fmt.Println(err)
	}
}
