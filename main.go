package main

import (
	"fmt"
	"os"

	goreloaded "goreloaded/openfile"
)

func main() {
	args := os.Args[1:]
	inputFilepath := ""
	outputFilepath := ""
	if len(args) < 1 {
		fmt.Println("usage: go run . sample.txt result.txt")
		return
	}
	if len(args) == 1 {
		inputFilepath = args[0]
		outputFilepath = "result.txt"
	}
	if len(args) == 2 {
		inputFilepath = os.Args[1]
		outputFilepath = os.Args[2]
	}

	 err := goreloaded.OpenFile(inputFilepath, outputFilepath)
	 if  err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Success")
}
