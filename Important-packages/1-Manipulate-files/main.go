package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	length, err := file.Write([]byte("Change hello"))
	// file.WriteString("Hello, World")

	if err != nil {
		panic(err)
	}

	fmt.Printf("The file length is %d bytes\n", length)

	// Reading
	getFile, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(getFile))

	// Reading with buffer (For big files)

	file2, err := os.Open("file.txt")

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3)
	formattedString := ""

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		formattedString += string(buffer[:n])
		fmt.Println(string(buffer[:n]))
	}

	fmt.Println(formattedString)

	os.Remove("file.txt")
}
