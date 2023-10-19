package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"regexp"
)

func main() {
	f, err := os.Open("../puzzle_input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Could not read file: %v", err)
		os.Exit(1)
	}
	fileAsString := string(content)

	parts := strings.Split(fileAsString, "\n\n")

	stacksAsStringUncleaned := parts[0]

	_mid := strings.Split(stacksAsStringUncleaned, "\n")

	stacksAsString := _mid[:len(_mid)-1]

	reg, err := regexp.Compile("[A-Z]")
	if err != nil {
		fmt.Printf("Could not compile regex: %v", err)
	}

	for _, line := range stacksAsString {
		fmt.Println(reg.FindString(line[0:4]))
	}
}
