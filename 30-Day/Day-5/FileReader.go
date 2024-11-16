package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFileContent(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("file %s does not exist", filename)
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

func main() {
	content, err := readFileContent("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File Content:", content)
	}
}
