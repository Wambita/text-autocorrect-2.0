package goreloaded

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	goreloaded "goreloaded/textmodifications"
)

// Function to open input file and create output file if one does not exist
func OpenFile(inputFilepath string, outputFilepath string) error {
	fileInfo, err := os.Stat(inputFilepath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return err
	}
	if err != nil {
		fmt.Printf("Error checking the file: %s\n", err)
		return err
	}
	// check if the file size is greater than zero
	if fileInfo.Size() < 0 {
		fmt.Println("File exists but is empty")
	}

	if filepath.Ext(inputFilepath) != ".txt" {
		fmt.Printf("Wrong file extension. Use .txt files")
	}

	inputFile, err := os.Open(inputFilepath)
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilepath)
	if filepath.Ext(outputFilepath) != ".txt" {
		fmt.Printf("Wrong file extension. Use .txt files")
	}

	if err != nil {
		fmt.Println(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		line = goreloaded.TextModification(line)
		fmt.Fprintln(outputFile, line)
		}

	return nil
}
