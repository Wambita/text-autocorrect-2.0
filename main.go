package main

import (
	"fmt"
	"os"

	goreloaded "goreloaded/openfile"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . sample.txt result.txt")
		return
	}
	inputFilepath := os.Args[1]
	outputFilepath := os.Args[2]
	if err := goreloaded.OpenFile(inputFilepath, outputFilepath); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Success")
}
