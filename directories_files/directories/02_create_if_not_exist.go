package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("dir1")

	if os.IsNotExist(err) {
		err = os.Mkdir("dir1", 0755)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Directory created!")
	} else {
		fmt.Println("Directory already exist!")
	}
}
