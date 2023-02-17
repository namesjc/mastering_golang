package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("testing2/test2.txt", os.O_WRONLY, 0666)
	check(err)
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file) // create a buffered writer
	bytesWritten, err := bufferedWriter.Write([]byte{65, 66, 67})
	check(err)
	log.Printf("Bytes written: %d\n", bytesWritten) // 2023/02/10 22:45:46 Bytes written: 3

	bytesWritten, err = bufferedWriter.WriteString("Buffered string\n")
	check(err)
	log.Printf("Bytes written: %d\n", bytesWritten) // 2023/02/10 22:45:46 Bytes written: 16

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize) // 2023/02/10 22:45:46 Bytes buffered: 19

	byteAvailable := bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", byteAvailable) // 2023/02/10 22:45:46 Available buffer: 4077

	bufferedWriter.Flush() //Write memory buffer to disk
	bufferedWriter.Reset(bufferedWriter)

	byteAvailable = bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", byteAvailable) // 2023/02/10 22:45:46 Available buffer: 4096

	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)
	check(err)
	byteAvailable = bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", byteAvailable) // 2023/02/10 22:49:54 Available buffer: 8000
}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}
