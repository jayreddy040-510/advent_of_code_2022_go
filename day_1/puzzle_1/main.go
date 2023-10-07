package main

import (
	// imports
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../elves_and_their_calories.txt")
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening puzzle input: %v", err)
	}
	scanner := bufio.NewScanner(file)
	var highCount int
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if count > highCount {
				highCount = count
			}
			count = 0
		} else {
			countAsInt, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Couldn't process puzzle input: %v", err)
			}
			count += countAsInt
		}
	}
	log.Printf("Highest calories: %d", highCount)
	return
}
