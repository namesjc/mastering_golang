package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Fprint(os.Stdout, "Hello world!")
	io.WriteString(os.Stdout, "Hello world!")
}
