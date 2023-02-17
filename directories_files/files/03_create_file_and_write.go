package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	d1 := []byte("Hello\nGolang\n")
	err := ioutil.WriteFile("testing2/test1.txt", d1, 0644)
	check(err)

	f, err := os.Create("testing2/test2.txt") //open a file for writing.
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10} // d2 is string "some", check https://www.ascii-code.com/

	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2) // wrote 5 bytes

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // Issue a Sync to flush writes to stable storage.

}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}
