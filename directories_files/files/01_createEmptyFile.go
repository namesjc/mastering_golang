package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("testing/test1.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.Println(newFile)
	newFile.Close()

	fileInfo, err := os.Stat("testing/test1.txt") //os.FileInfo
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	// 2023/02/09 19:21:04 &{0xc00006e180}
	// File name: test1.txt
	// Size in bytes: 0
	// Permissions: -rw-r--r--
	// Last modified: 2023-02-09 19:21:04.523123805 +0800 CST
	// Is Directory: false
}
