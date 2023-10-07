package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func updateTopThree(topThree *[3]int, num int) {
	if num > topThree[0] {
		topThree[2] = topThree[1]
		topThree[1] = topThree[0]
		topThree[0] = num
	} else if num > topThree[1] {
		topThree[2] = topThree[1]
		topThree[1] = num
	} else if num > topThree[2] {
		topThree[2] = num
	}
}

func main() {
	file, err := os.Open("../elves_and_their_calories.txt")
	if err != nil {
		log.Fatalf("Error opening puzzle input: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var topThree [3]int
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			updateTopThree(&topThree, count)
			count = 0
		} else {
			lineAsInt, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("Error processing puzzle input: %v", err)
			}
			count += lineAsInt
		}
	}
	// if file doesn't end with newline
	if count > 0 {
		updateTopThree(&topThree, count)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning puzzle input: %v", err)
	}
	log.Println(topThree[0] + topThree[1] + topThree[2])
}
