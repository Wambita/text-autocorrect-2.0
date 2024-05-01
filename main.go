package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . sample.txt result.txt")
		return
	}
	inputFilepath := os.Args[1]
	outputFilepath := os.Args[2]
}
