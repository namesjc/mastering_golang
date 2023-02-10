package main

import (
	"log"
	"os"
)

func main() {
	// Difference between Mkdir and MkdirAll
	//	MkdirAll command create every parent which doesn't exist mkdir -p
	//  Mkdir is for creating one single directory

	err := os.Mkdir("dir1", 0700)
	if err != nil {
		log.Println(err)
	}

	err = os.MkdirAll("dir2/dir3", 0700) //Dir2 will be also created!
	if err != nil {
		log.Println(err)
	}

}
