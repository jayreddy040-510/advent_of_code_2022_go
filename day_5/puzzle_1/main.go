package main

import (
	"bufio"
	"fmt"
	"os"
	// "regexp"
)

func main() {
	f, err := os.Open("../puzzle_input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
        if len(line) == 0 {
            return
        }
		i++
        if i == 15 {
            return
        }
		fmt.Println(line)
	}
}
