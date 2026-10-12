package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument.")
		return
	}

	//data, err := os.ReadFile("sample.txt")
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:" + string(data))
	fmt.Println("File has ", len(data), "bytes")
}

/*
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("All args", os.Args)
	fmt.Println("Number of args", len(os.Args))
}
*/
