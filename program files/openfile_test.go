package goreloaded

import "testing"

func TestOpenFile(t *testing.T) {
	// test for valid file path
	inputFilepath := "sample.txt"
	_, err := OpenFile(inputFilepath)
	if err != nil {
		t.Errorf("Failed to open existing file: %v", err)
	}
	outputFilepath := "result.txt"
	_, err := OpenFile(outputFilepath)
	if err != nil {
		t.Errorf("Failed to open existing file: %v", err)
	}
}

// test for non existent file
func TestNonExistentFile(t *testing.T) {
	inputFilepath := "non-existent.txt"
	_, err := OpenFile(inputFilepath)
	if err != nil {
		t.Errorf("Expected error for non existent file, but got none")
	}
	outputFilepath := "non-existent.txt"
	_, err := OpenFile(outputFilepath)
	if err != nil {
		t.Errorf("Expected error for non existent file, but got none")
	}
}

//test for wrong extension 
func TestWrongExtension(t *testing.T){
	inputFilepath = "image.png"
	if filepath.Ext(inputFilepath) == ".txt"{
		t.Errorf("Error: file has .txt expected wrong filepath")
	}
	outputFilepath = "image.png"
	if filepath.Ext(outputFilepath) == ".txt"{
		t.Errorf("Error: file has .txt expected wrong filepath")
	}
}
//test empty file
