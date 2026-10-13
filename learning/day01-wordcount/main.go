package main

import (
	"fmt"
	"os"
	"strings"
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

	// Display content of the file and its size
	fmt.Println("File content:" + string(data))
	fmt.Println("File has ", len(data), "bytes")

	// Split the content into words and count them
	words := strings.Fields(string(data))
	fmt.Println("Word count ", len(words))
	fmt.Println("First word", words[0])
	fmt.Println("Last Word ", words[len(words)-1])

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	fmt.Println("Word counts:", counts)

}

/*
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("All args", os.Args)
	fmt.Println("Number of args", len(os.Args))
}
*/
