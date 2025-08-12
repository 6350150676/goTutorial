package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}

	content := "red eyes dead guys"
	bytesWritten, err := io.WriteString(file, content)
	fmt.Println("Bytes written:", bytesWritten)
	if err != nil {
		fmt.Println("Error while writing to file:", err)
	}

	file.Close()

	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file:", err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}

	data, error2 := os.ReadFile("example.txt") // returns []byte
	if error2 != nil {
		fmt.Println("Error while reading file:", error2)
		return
	}
	fmt.Println("\nFile content using ReadFile():", string(data))
}
