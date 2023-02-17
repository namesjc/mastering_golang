package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Truncate a file to 100 bytes
	// If file is less than 100 bytes the origianl contents will remain
	// at the beginning, and the rest of the spce is
	// filled will null bytes. If it is over 100 bytes,
	// Everything past 100 bytes will be lost.
	// Either way we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely empty file.

	err := os.Truncate("testing2/test3.txt", 3)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(err)
}
